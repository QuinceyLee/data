package engine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"rulex/core"
	"rulex/source"
	"rulex/statistics"
	"rulex/target"
	"rulex/typex"
	"rulex/utils"
	"runtime"
	"sync"
	"time"

	"github.com/ngaut/log"
	"github.com/shirou/gopsutil/disk"
	lua "github.com/yuin/gopher-lua"
)

//
//
//
type RuleEngine struct {
	Hooks   *sync.Map          `json:"hooks"`
	Rules   *sync.Map          `json:"rules"`
	Plugins *sync.Map          `json:"plugins"`
	InEnds  *sync.Map          `json:"inends"`
	OutEnds *sync.Map          `json:"outends"`
	Drivers *sync.Map          `json:"drivers"`
	Config  *typex.RulexConfig `json:"config"`
}

//
//
//
func NewRuleEngine(config typex.RulexConfig) typex.RuleX {
	return &RuleEngine{
		Plugins: &sync.Map{},
		Hooks:   &sync.Map{},
		Rules:   &sync.Map{},
		InEnds:  &sync.Map{},
		OutEnds: &sync.Map{},
		Drivers: &sync.Map{},
		Config:  &config,
	}
}

//
//
//
func (e *RuleEngine) Start() *typex.RulexConfig {
	typex.StartQueue(core.GlobalConfig.MaxQueueSize)
	source.LoadSt()
	target.LoadTt()
	return e.Config
}

//
//
//
func (e *RuleEngine) PushQueue(qd typex.QueueData) error {
	err := typex.DefaultDataCacheQueue.Push(qd)
	if err != nil {
		log.Error("PushQueue error:", err)
		statistics.IncInFailed()
	} else {
		statistics.IncIn()
	}
	return err
}
func (e *RuleEngine) PushInQueue(in *typex.InEnd, data string) error {
	qd := typex.QueueData{
		E:    e,
		In:   in,
		Out:  nil,
		Data: data,
	}
	err := typex.DefaultDataCacheQueue.Push(qd)
	if err != nil {
		log.Error("PushInQueue error:", err)
		statistics.IncInFailed()
	} else {
		statistics.IncIn()
	}
	return err
}
func (e *RuleEngine) PushOutQueue(out *typex.OutEnd, data string) error {
	qd := typex.QueueData{
		E:    e,
		In:   nil,
		Out:  out,
		Data: data,
	}
	err := typex.DefaultDataCacheQueue.Push(qd)
	if err != nil {
		log.Error("PushOutQueue error:", err)
		statistics.IncInFailed()
	} else {
		statistics.IncIn()
	}
	return err
}

//
//
//
func (e *RuleEngine) GetPlugins() *sync.Map {
	return e.Plugins
}
func (e *RuleEngine) AllPlugins() *sync.Map {
	return e.Plugins
}

//
//
//
func (e *RuleEngine) Version() typex.Version {
	return typex.DefaultVersion
}

//
//
func (e *RuleEngine) GetConfig() *typex.RulexConfig {
	return e.Config
}

/*
*
* TODO: 0.3.0????????????????????? SourceRegistry ??????
*
 */
func (e *RuleEngine) LoadInEnd(in *typex.InEnd) error {
	if in.Type == typex.MQTT {
		return startSources(source.NewMqttInEndSource(in.UUID, e), in, e)
	}
	if in.Type == typex.HTTP {
		return startSources(source.NewHttpInEndSource(in.UUID, e), in, e)
	}
	if in.Type == typex.COAP {
		return startSources(source.NewCoAPInEndSource(in.UUID, e), in, e)
	}
	if in.Type == typex.GRPC {
		return startSources(source.NewGrpcInEndSource(in.UUID, e), in, e)
	}
	if in.Type == typex.UART_MODULE {
		return startSources(source.NewUartModuleSource(in.UUID, e), in, e)
	}
	if in.Type == typex.MODBUS_MASTER {
		return startSources(source.NewModbusMasterSource(in.UUID, e), in, e)
	}
	if in.Type == typex.SNMP_SERVER {
		return startSources(source.NewSNMPInEndSource(in.UUID, e), in, e)
	}
	if in.Type == typex.NATS_SERVER {
		return startSources(source.NewNatsSource(e), in, e)
	}
	if in.Type == typex.SIEMENS_S7 {
		return startSources(source.NewSiemensS7Source(e), in, e)
	}
	if in.Type == typex.RULEX_UDP {
		return startSources(source.NewUdpInEndSource(e), in, e)
	}
	return fmt.Errorf("unsupported InEnd type:%s", in.Type)
}

//
// start Sources
//
/*
* Life cycle
+------------------+       +------------------+   +---------------+        +---------------+
|     Register     |------>|   Start          |-->|     Test      |--+ --->|  Stop         |
+------------------+  |    +------------------+   +---------------+  |     +---------------+
                      |                                              |
                      |                                              |
                      +-------------------Error ---------------------+
*/
func startSources(source typex.XSource, in *typex.InEnd, e *RuleEngine) error {
	//
	// ?????????, ????????????????????????????????????
	//
	// ???????????????ID???????????????, ????????????????????????????????????????????????????????????
	e.SaveInEnd(in)
	// Load config
	config := e.GetInEnd(in.UUID).Config
	if config == nil {
		e.RemoveInEnd(in.UUID)
		err := fmt.Errorf("source [%v] config is nil", in.Name)
		return err
	}

	if err := source.Init(in.UUID, config); err != nil {
		log.Error(err)
		e.RemoveInEnd(in.UUID)
		return err
	}
	// Set sources to inend
	in.Source = source
	// ??????????????????
	if err := startSource(source, e, in.UUID); err != nil {
		log.Error(err)
		e.RemoveInEnd(in.UUID)
		return err
	}
	ticker := time.NewTicker(time.Duration(time.Second * 5))
	go func(ctx context.Context) {
		// 5 seconds
	TICKER:
		<-ticker.C
		select {
		case <-ctx.Done():
			{
				return
			}
		default:
			{
				goto CHECK
			}
		}
	CHECK:
		{
			//
			// ??????HTTP?????????????????????, ??????????????????, ???????????????????????????, ???????????????
			//
			if source.Details() == nil {
				return
			}
			//------------------------------------
			// ???????????????????????????, ??????????????????????????????
			//------------------------------------
			tryIfRestartSource(source, e, in.UUID)
			// checkDriverState(source, e, in.UUID)
			//------------------------------------
			goto TICKER
		}

	}(typex.GCTX)
	log.Infof("InEnd [%v, %v] load successfully", in.Name, in.UUID)
	return nil
}

/*
*
* ????????????????????????????????????
* ???????????????????????????????????????????????????????????????????????????
*
 */
func checkDriverState(source typex.XSource, e *RuleEngine, id string) {
	if source.Driver() != nil {
		// ???????????????????????????????????????
		if source.Status() == typex.UP {
			// ??????????????????, ????????????????????????
			if source.Driver().State() == typex.STOP {
				log.Warn("Driver stopped:", source.Driver().DriverDetail().Name)
				// ???????????????????????????, ??????????????????
				source.Stop()
			}
		}
	}

}

//
// test SourceState
//
func tryIfRestartSource(source typex.XSource, e *RuleEngine, id string) {
	checkDriverState(source, e, id)
	if source.Status() == typex.DOWN {
		source.Details().SetState(typex.DOWN)
		//----------------------------------
		// ?????????????????????????????????, ????????????
		//----------------------------------
		log.Warnf("Source %v %v down. try to restart it", source.Details().UUID, source.Details().Name)
		source.Stop()
		//----------------------------------
		// ????????????????????????
		//----------------------------------
		runtime.Gosched()
		runtime.GC() // GC ?????????, ?????????????????????, ????????????
		startSource(source, e, id)
	} else {
		source.Details().SetState(typex.UP)
	}
}

//
//
//
func startSource(source typex.XSource, e *RuleEngine, id string) error {
	//----------------------------------
	// ???????????? ?????????????????????????????????
	//----------------------------------
	ctx, cancelCTX := context.WithCancel(typex.GCTX)

	if err := source.Start(typex.CCTX{Ctx: ctx, CancelCTX: cancelCTX}); err != nil {
		log.Error("Source start error:", err)
		if source.Status() == typex.UP {
			source.Stop()
		}
		if source.Driver() != nil {
			if source.Driver().State() == typex.RUNNING {
				source.Driver().Stop()
			}
		}
		return err
	} else {
		//----------------------------------
		// ??????????????????, ????????????
		//----------------------------------
		if source.Driver() != nil {
			if source.Driver().State() == typex.RUNNING {
				source.Driver().Stop()
			}
			// Start driver
			if err := source.Driver().Init(); err != nil {
				log.Error("Driver initial error:", err)
				return errors.New("Driver initial error:" + err.Error())
			}
			log.Infof("Try to start driver: [%v]", source.Driver().DriverDetail().Name)
			if err := source.Driver().Work(); err != nil {
				log.Error("Driver work error:", err)
				return errors.New("Driver work error:" + err.Error())
			}
			log.Infof("Driver start successfully: [%v]", source.Driver().DriverDetail().Name)
		}
		return nil
	}

}

//
// LoadOutEnd
//
func (e *RuleEngine) LoadOutEnd(out *typex.OutEnd) error {
	return tryCreateOutEnd(out, e)
}

//
// CreateOutEnd
//
func tryCreateOutEnd(out *typex.OutEnd, e typex.RuleX) error {
	if out.Type == typex.MONGO_SINGLE {
		return startTarget(target.NewMongoTarget(e), out, e)
	}
	if out.Type == typex.MQTT_TARGET {
		return startTarget(target.NewMqttTarget(e), out, e)
	}
	if out.Type == typex.NATS_TARGET {
		return startTarget(target.NewNatsTarget(e), out, e)
	}
	if out.Type == typex.HTTP_TARGET {
		return startTarget(target.NewHTTPTarget(e), out, e)
	}
	if out.Type == typex.TDENGINE_TARGET {
		return startTarget(target.NewTdEngineTarget(e), out, e)
	}
	if out.Type == typex.GRPC_CODEC_TARGET {
		return startTarget(target.NewCodecTarget(e), out, e)
	}
	return errors.New("unsupported target type:" + out.Type.String())

}

//
// Start output target
//
// Target life cycle:
//     Register -> Start -> running/restart cycle
//
func startTarget(target typex.XTarget, out *typex.OutEnd, e typex.RuleX) error {
	//
	// ?????????, ????????????????????????????????????
	//
	e.SaveOutEnd(out)

	// Load config
	config := e.GetOutEnd(out.UUID).Config
	if config == nil {
		e.RemoveOutEnd(out.UUID)
		err := fmt.Errorf("target [%v] config is nil", out.Name)
		return err
	}
	if err := target.Init(out.UUID, config); err != nil {
		log.Error(err)
		e.RemoveInEnd(out.UUID)
		return err
	}
	// ??????????????????
	ctx, cancelCTX := context.WithCancel(typex.GCTX)
	if err := target.Start(typex.CCTX{Ctx: ctx, CancelCTX: cancelCTX}); err != nil {
		log.Error(err)
		e.RemoveOutEnd(out.UUID)
		return err
	}
	// Set sources to inend
	out.Target = target
	//
	ticker := time.NewTicker(time.Duration(time.Second * 5))
	tryIfRestartTarget(target, e, out.UUID)
	go func(ctx context.Context) {

		// 5 seconds
		//
	TICKER:
		<-ticker.C
		select {
		case <-ctx.Done():
			{
				return
			}
		default:
			{
				goto CHECK
			}
		}
	CHECK:
		{
			if target.Details() == nil {
				return
			}
			tryIfRestartTarget(target, e, out.UUID)
			goto TICKER
		}

	}(typex.GCTX)
	log.Infof("Target [%v, %v] load successfully", out.Name, out.UUID)
	return nil
}

//
// ????????????, ??????????????????
//
func tryIfRestartTarget(target typex.XTarget, e typex.RuleX, id string) {
	if target.Status() == typex.DOWN {
		target.Details().SetState(typex.DOWN)
		log.Warnf("Target [%v, %v] down. try to restart it", target.Details().Name, target.Details().UUID)
		target.Stop()
		runtime.Gosched()
		runtime.GC()
		ctx, cancelCTX := context.WithCancel(typex.GCTX)
		target.Start(typex.CCTX{Ctx: ctx, CancelCTX: cancelCTX})
	} else {
		target.Details().SetState(typex.UP)
	}
}

// LoadRule
func (e *RuleEngine) LoadRule(r *typex.Rule) error {
	if err := core.VerifyCallback(r); err != nil {
		return err
	} else {
		// ??????????????????ID
		if len(r.From) > 0 {
			for _, inUUId := range r.From {
				// ???????????????????????????????????????
				if in := e.GetInEnd(inUUId); in != nil {
					// Bind to rule, Key:RuleId, Value: Rule
					// RULE_0f8619ef-3cf2-452f-8dd7-aa1db4ecfdde {
					// ...
					// ...
					// }
					// ??????????????????????????????????????????
					(in.Binds)[r.UUID] = *r
					//--------------------------------------------------------------
					// Load LoadBuildInLuaLib
					//--------------------------------------------------------------
					LoadBuildInLuaLib(e, r)
					//--------------------------------------------------------------
					// Save to rules map
					//--------------------------------------------------------------
					e.SaveRule(r)
					log.Infof("Rule [%v, %v] load successfully", r.Name, r.UUID)
					return nil
				} else {
					return errors.New("'InEnd':" + inUUId + " is not exists")
				}
			}
		}
	}
	return errors.New("'From' can not be empty")

}

//
// GetRule a rule
//
func (e *RuleEngine) GetRule(id string) *typex.Rule {
	v, ok := (e.Rules).Load(id)
	if ok {
		return v.(*typex.Rule)
	} else {
		return nil
	}
}

//
//
//
func (e *RuleEngine) SaveRule(r *typex.Rule) {
	e.Rules.Store(r.UUID, r)
}

//
// RemoveRule and inend--rule bindings
//
func (e *RuleEngine) RemoveRule(ruleId string) {
	if rule := e.GetRule(ruleId); rule != nil {
		// ?????? InEnd ??? bind ??????
		inEnds := e.AllInEnd()
		inEnds.Range(func(key, value interface{}) bool {
			inEnd := value.(*typex.InEnd)
			for _, r := range inEnd.Binds {
				if rule.UUID == r.UUID {
					delete(inEnd.Binds, ruleId)
				}
			}
			return true
		})
		e.Rules.Delete(ruleId)
		rule = nil
		log.Infof("Rule [%v] has been deleted", ruleId)
	}
}

//
//
//
func (e *RuleEngine) AllRule() *sync.Map {
	return e.Rules
}

//
// Stop
//
func (e *RuleEngine) Stop() {
	log.Info("Ready to stop rulex")
	e.InEnds.Range(func(key, value interface{}) bool {
		inEnd := value.(*typex.InEnd)
		if inEnd.Source != nil {
			log.Info("Stop InEnd:", inEnd.Name, inEnd.UUID)
			e.GetInEnd(inEnd.UUID).SetState(typex.DOWN)
			inEnd.Source.Stop()
			if inEnd.Source.Driver() != nil {
				inEnd.Source.Driver().SetState(typex.STOP)
				inEnd.Source.Driver().Stop()
			}
		}
		return true
	})
	e.OutEnds.Range(func(key, value interface{}) bool {
		outEnd := value.(*typex.OutEnd)
		if outEnd.Target != nil {
			log.Info("Stop Target:", outEnd.Name, outEnd.UUID)
			outEnd.Target.Stop()
		}
		return true
	})

	e.Plugins.Range(func(key, value interface{}) bool {
		plugin := value.(typex.XPlugin)
		log.Info("Stop plugin:", plugin.PluginMetaInfo().Name)
		plugin.Stop()
		return true
	})
	// ????????????
	runtime.Gosched()
	runtime.GC()
	log.Info("Stop Rulex successfully")
}

//
// ????????????: Work
//
func (e *RuleEngine) Work(in *typex.InEnd, data string) (bool, error) {
	if err := e.PushInQueue(in, data); err != nil {
		return false, err
	}
	return true, nil
}

//
// ??????lua??????
//
func (e *RuleEngine) RunLuaCallbacks(in *typex.InEnd, callbackArgs string) {
	for _, rule := range in.Binds {
		if rule.Status == typex.RULE_RUNNING {
			_, err := core.ExecuteActions(&rule, lua.LString(callbackArgs))
			if err != nil {
				log.Error("RunLuaCallbacks error:", err)
				_, err := core.ExecuteFailed(rule.VM, lua.LString(err.Error()))
				if err != nil {
					log.Error(err)
				}
			} else {
				_, err := core.ExecuteSuccess(rule.VM)
				if err != nil {
					log.Error(err)
					return
				}
			}
		}
	}
}

//
// ????????????????????????    ????????????????????????    ????????????????????????
// ??? Init ?????????????????? Load ?????????????????? Stop ???
// ????????????????????????    ????????????????????????    ????????????????????????
//
func (e *RuleEngine) LoadPlugin(sectionK string, p typex.XPlugin) error {
	section := utils.GetINISection(core.INIPath, sectionK)
	key, err1 := section.GetKey("enable")
	if err1 != nil {
		return err1
	}
	enable, err2 := key.Bool()
	if err2 != nil {
		return err2
	}
	if !enable {
		log.Infof("Plugin is not enable:%s", p.PluginMetaInfo().Name)
		return nil
	}

	if err := p.Init(section); err != nil {
		return err
	}
	_, ok := e.Plugins.Load(p.PluginMetaInfo().Name)
	if ok {
		return errors.New("plugin already installed:" + p.PluginMetaInfo().Name)
	}

	if err := p.Start(); err != nil {
		return err
	}

	e.Plugins.Store(p.PluginMetaInfo().Name, p)
	log.Infof("Plugin start successfully:[%v]", p.PluginMetaInfo().Name)
	return nil

}

//
// LoadHook
//
func (e *RuleEngine) LoadHook(h typex.XHook) error {
	value, _ := e.Hooks.Load(h.Name())
	if value != nil {
		return errors.New("hook have been loaded:" + h.Name())
	} else {
		e.Hooks.Store(h.Name(), h)
		return nil
	}
}

//
// RunHooks
//
func (e *RuleEngine) RunHooks(data string) {
	e.Hooks.Range(func(key, value interface{}) bool {
		if err := runHook(value.(typex.XHook), data); err != nil {
			value.(typex.XHook).Error(err)
		}
		return true
	})
}
func runHook(h typex.XHook, data string) error {
	return h.Work(data)
}

//
//
//
func (e *RuleEngine) GetInEnd(uuid string) *typex.InEnd {
	v, ok := (e.InEnds).Load(uuid)
	if ok {
		return v.(*typex.InEnd)
	} else {
		return nil
	}
}

//
//
//
func (e *RuleEngine) SaveInEnd(in *typex.InEnd) {
	e.InEnds.Store(in.UUID, in)
}

//
//
//
func (e *RuleEngine) RemoveInEnd(id string) {
	if inEnd := e.GetInEnd(id); inEnd != nil {
		e.InEnds.Delete(id)
		inEnd = nil
		log.Infof("InEnd [%v] has been deleted", id)
	}
}

//
//
//
func (e *RuleEngine) AllInEnd() *sync.Map {
	return e.InEnds
}

//
//
//
func (e *RuleEngine) GetOutEnd(id string) *typex.OutEnd {
	v, ok := e.OutEnds.Load(id)
	if ok {
		return v.(*typex.OutEnd)
	} else {
		return nil
	}

}

//
//
//
func (e *RuleEngine) SaveOutEnd(out *typex.OutEnd) {
	e.OutEnds.Store(out.UUID, out)

}

//
//
//
func (e *RuleEngine) RemoveOutEnd(uuid string) {
	if inEnd := e.GetOutEnd(uuid); inEnd != nil {
		e.OutEnds.Delete(uuid)
		inEnd = nil
		log.Infof("InEnd [%v] has been deleted", uuid)
	}
}

//
//
//
func (e *RuleEngine) AllOutEnd() *sync.Map {
	return e.OutEnds
}

//-----------------------------------------------------------------
// ?????????????????????
//-----------------------------------------------------------------
func (e *RuleEngine) SnapshotDump() string {
	inends := []interface{}{}
	rules := []interface{}{}
	plugins := []interface{}{}
	outends := []interface{}{}
	e.AllInEnd().Range(func(key, value interface{}) bool {
		inends = append(inends, value)
		return true
	})
	e.AllRule().Range(func(key, value interface{}) bool {
		rules = append(rules, value)
		return true
	})
	e.AllPlugins().Range(func(key, value interface{}) bool {
		plugins = append(plugins, (value.(typex.XPlugin)).PluginMetaInfo())
		return true
	})
	e.AllOutEnd().Range(func(key, value interface{}) bool {
		outends = append(outends, value)
		return true
	})
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	system := map[string]interface{}{
		"version":  e.Version().Version,
		"diskInfo": int(diskInfo.UsedPercent),
		"system":   bToMb(m.Sys),
		"alloc":    bToMb(m.Alloc),
		"total":    bToMb(m.TotalAlloc),
		"osArch":   runtime.GOOS + "-" + runtime.GOARCH,
	}
	data := map[string]interface{}{
		"rules":      rules,
		"plugins":    plugins,
		"inends":     inends,
		"outends":    outends,
		"statistics": statistics.AllStatistics(),
		"system":     system,
		"config":     core.GlobalConfig,
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
	}
	return string(b)
}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

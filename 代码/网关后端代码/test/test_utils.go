package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rulex/core"
	"rulex/engine"
	"rulex/typex"
	"time"

	"github.com/ngaut/log"
)

//
func HttpPost(data map[string]interface{}, url string) string {
	p, errs1 := json.Marshal(data)
	if errs1 != nil {
		log.Fatal(errs1)
	}
	r, errs2 := http.Post(url, "application/json", bytes.NewBuffer(p))
	if errs2 != nil {
		log.Fatal(errs2)
	}
	defer r.Body.Close()

	body, errs5 := ioutil.ReadAll(r.Body)
	if errs5 != nil {
		log.Fatal(errs5)
	}
	return string(body)
}

func HttpGet(api string) string {
	var err error
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Error(err)
		return ""
	}

	response, err := (&http.Client{}).Do(request)
	if err != nil {
		log.Error(err)
		return ""
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return ""
	}
	return string(body)
}

/*
*
* 起一个测试服务
*
 */
func TestEngine() typex.RuleX {
	mainConfig := core.InitGlobalConfig("conf/rulex.ini")
	core.StartStore(core.GlobalConfig.MaxQueueSize)
	engine := engine.NewRuleEngine(mainConfig)
	return engine
}

/*
*
* 生成测试数据库的文件名
*
 */
func GenDate() string {
	return "rulex-test_" + time.Now().Format("2006-01-02-15_04_05")
}

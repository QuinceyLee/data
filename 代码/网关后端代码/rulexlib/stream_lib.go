package rulexlib

import (
	"rulex/typex"

	lua "github.com/yuin/gopher-lua"
)

/*
*
* 数据写入资源外挂的驱动, 比如数据可能是串口输入进来的
*
 */

func WriteInStream(rx typex.RuleX) func(*lua.LState) int {
	return func(l *lua.LState) int {
		id := l.ToString(2)
		data := l.ToString(3)
		ok := writeInStream(rx, id, data)
		if ok {
			l.Push(lua.LTrue)
		} else {
			l.Push(lua.LFalse)
		}
		return 1
	}
}

/*
*
* 数据输出到出口的外挂驱动, 比如控制某个硬件
*
 */
func WriteOutStream(rx typex.RuleX) func(*lua.LState) int {
	return func(l *lua.LState) int {
		id := l.ToString(2)
		data := l.ToString(3)
		ok := writeOutStream(rx, id, data)
		if ok {
			l.Push(lua.LTrue)
		} else {
			l.Push(lua.LFalse)
		}
		return 1
	}
}

/*
*
* 数据写到资源的驱动里面
*
 */
func writeInStream(e typex.RuleX, id string, data string) bool {
	err := e.GetInEnd(id).Source.OnStreamApproached(data)
	if err != nil {
		return true
	} else {
		return false
	}
}

/*
*
* 数据写到出口的驱动里面
*
 */
func writeOutStream(e typex.RuleX, id string, data string) bool {
	err := e.GetOutEnd(id).Target.OnStreamApproached(data)
	if err != nil {
		return true
	} else {
		return false
	}
}

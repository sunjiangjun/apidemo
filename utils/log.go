package utils

import (
	"github.com/astaxie/beego/logs"
)


var Log  *logs.BeeLogger

func init()  {
	Log= logs.NewLogger()
	logs.SetLogger(logs.AdapterConsole)
	logs.EnableFuncCallDepth(true)
}

func Debug(msg interface{})  {
	Log.Debug("",msg)
}

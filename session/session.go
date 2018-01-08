package session

import "github.com/astaxie/beego/session"

var GlobalSession *session.Manager

func init()  {

	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}

	GlobalSession,_=session.NewManager("memory",sessionConfig)

	go GlobalSession.GC()
}

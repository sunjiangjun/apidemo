// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apidemo/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"net/http"
	"apidemo/session"
	"apidemo/utils"
	"apidemo/cache"
)

type ImageHandler struct {

}

func (image *ImageHandler) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	sess,_:=session.GlobalSession.SessionStart(res,req)
	defer sess.SessionRelease(res)
	res.Write([]byte("welcome ...."))

	//session
	sess.Set("username","sunhongtao")
	utils.Debug(sess.Get("username"))
	utils.Debug(sess.SessionID())

	//cache

	cache.CacheClient.Put("username","sun_hongtao",0)

	utils.Debug(cache.CacheClient.Get("username"))
}



func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSGet("/date", func(i *context.Context) {
			i.ResponseWriter.Write([]byte("hello.world"))
		}),


		beego.NSRouter("/log",&controllers.MessageController{}),
	)
	beego.AddNamespace(ns)

	beego.Router("/v1/api/message", &controllers.MessageController{},"*:Get")
	beego.Get("/v1/api/data", func(context *context.Context) {
		context.ResponseWriter.Write([]byte("hello.world"))
	})

	beego.Handler("/v1/api/image",&ImageHandler{})

	beego.AutoRouter(&controllers.MessageController{})

	//beego.Include(&controllers.ApiController{})
}

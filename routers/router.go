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
	"github.com/astaxie/beego/orm"
	"apidemo/models"
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

	//orm

	//insert
	//l:=models.Nblog{Id:1,Data:"hello"}
	//o:=orm.NewOrm()
	//index,e:=o.Insert(&l)
	//if e!=nil {
	//	panic(e)
	//}
	//utils.Debug(index)

   //read
	o:=orm.NewOrm()/**/
	//s:=models.Nblog{}
	//s.Id=2
	//o.Read(&s)
	//utils.Debug(s)


	//c,_,_:=o.ReadOrCreate(&s,"Data")
	//utils.Debug(c)


	//s.Data="sunhongtao"
	//o.Update(&s)


	//o.Delete(&s)

	//qs:=o.QueryTable(&s)
	//qs.Filter("id",1)
	//utils.Debug(s)


	//o.Raw("select * from n_blog where id=?",1).QueryRow(&s)
	//utils.Debug(s)


	//RowsToStruct
	s:=new(models.Nblog)
	i,_:=o.Raw("select id,data from n_blog ").RowsToStruct(s,"id","data")
	utils.Debug(s.Data)
	utils.Debug(i)


	//Prepare
	//p,_:=o.Raw("update n_blog set data=? where id = ?").Prepare()
	//p.Exec("sunhongtao",2)

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

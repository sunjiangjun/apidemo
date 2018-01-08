package cache

import "github.com/astaxie/beego/cache"

var CacheClient cache.Cache

func init()  {
	CacheClient, _ = cache.NewCache("memory", `{"interval":60}`)
	}

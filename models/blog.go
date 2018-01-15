package models

import "github.com/astaxie/beego/orm"

func init()  {
	orm.RegisterModel(new(Nblog))
}

type Nblog struct {
	Id int  `pk:"auto"`
	Data string
}

func (u *Nblog) TableName() string {
	return "n_blog"
}

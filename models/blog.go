package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

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

func (n *Nblog) String() string  {
	return "n.Id = "+strconv.Itoa(n.Id)+", n.Data = "+n.Data
}
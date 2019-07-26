package repository

import (
	"github.com/Deansquirrel/goToolMSSql2000"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
)

type repWl struct {
	dbConfig *goToolMSSql2000.MSSqlConfig
}

func NewRepWl() *repWl {
	comm := NewCommon()
	return &repWl{
		dbConfig: comm.ConvertDbConfigTo2000(comm.GetLocalDbConfig()),
	}
}

//获取配送出库货品明细
func (r *repWl) GetZ3PsCkHpDt() ([]*object.Z3PsCkHpDt, error) {
	//TODO
	return nil, nil
}

//删除配送出库货品明细索引
func (r *repWl) DelZ3PsCkHpDtSy(id string) error {
	//TODO
	return nil
}

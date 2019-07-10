package worker

import (
	"fmt"
	"github.com/Deansquirrel/goToolCron"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/global"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
	"strconv"
)

import log "github.com/Deansquirrel/goToolLog"

type common struct {
}

func NewCommon() *common {
	return &common{}
}

//系统配置检查
func (c *common) checkSysConfig() bool {
	if global.SysConfig.OnLineDb.Db == "" {
		log.Warn("线上库地址不能为空")
		global.Cancel()
		return false
	}
	if global.SysConfig.LocalDb.Server == "" ||
		global.SysConfig.LocalDb.DbName == "" ||
		global.SysConfig.LocalDb.User == "" {
		log.Warn("本地库地址、名称、用户名不能为空")
		global.Cancel()
		return false
	}
	return true
}

func (c *common) StartService(sType object.RunMode) {
	log.Debug("RunMode " + strconv.Itoa(int(sType)))
	if !c.checkSysConfig() {
		return
	}
	switch sType {
	case object.RunModeMd:
		c.addMdWorker()
	case object.RunModeGc:
		c.addWlWorker()
	default:
		log.Warn(fmt.Sprintf("unknown runmode %v", sType))
	}
}

func (c *common) panicHandle(v interface{}) {
	log.Error(fmt.Sprintf("panicHandle: %s", v))
}

func (c *common) addWorker(key string, cmd func()) {
	err := goToolCron.AddFunc(key, global.SysConfig.Task.Cron, cmd, c.panicHandle)
	if err != nil {
		errMsg := fmt.Sprintf("add job [%s] error: %s", key, err.Error())
		log.Error(errMsg)
	}
}

func (c *common) addMdWorker() {
	log.Debug("add md worker")
	worker := NewMdWorker()
	c.addWorker("Z3XsCkt", worker.Z3XsCkt)
	c.addWorker("Z3XsTht", worker.Z3XsTht)
	c.addWorker("Z3MdDbCkt", worker.Z3MdDbCkt)
	c.addWorker("Z3DbTzt", worker.Z3DbTzt)
	c.addWorker("Z3HpCkDjt", worker.Z3HpCkDjt)
	c.addWorker("Z3HpRkDjt", worker.Z3HpRkDjt)
	c.addWorker("Z3PkDjt", worker.Z3PkDjt)
	c.addWorker("Z3PsTzDt", worker.Z3PsTzDt)
	c.addWorker("Z3XsDdThDt", worker.Z3XsDdThDt)
}

func (c *common) addWlWorker() {
	log.Debug("add wl worker")
	worker := NewWlWorker()
	c.addWorker("Z3MdPsCkDt", worker.Z3MdPsCkDt)
	c.addWorker("Z3PsXzt", worker.Z3PsXzt)
	c.addWorker("Z3MdPsCht", worker.Z3MdPsCht)
	c.addWorker("Z3SheZhXsCkt", worker.Z3SheZhXsCkt)
	c.addWorker("ShengChWgRkt", worker.ShengChWgRkt)
	c.addWorker("XtTz", worker.XtTz)
	c.addWorker("Z3MdDhDt", worker.Z3MdDhDt)
	c.addWorker("Z3MdDhDtMd", worker.Z3MdDhDtMd)
	c.addWorker("OodXv1DdShCkt", worker.OodXv1DdShCkt)
}

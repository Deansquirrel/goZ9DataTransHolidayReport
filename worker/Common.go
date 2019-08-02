package worker

import (
	"fmt"
	"github.com/Deansquirrel/goToolCron"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"github.com/Deansquirrel/goToolSVRV3"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/global"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
	"github.com/kataras/iris/core/errors"
	"strconv"
	"time"
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
		log.Error("线上库地址不能为空")
		global.Cancel()
		return false
	}
	if object.RunMode(global.SysConfig.Task.RunMode) == object.RunModeDzq {
		if global.SysConfig.LocalDb.Server == "" ||
			global.SysConfig.LocalDb.DbName == "" ||
			global.SysConfig.LocalDb.User == "" {
			log.Error("本地库地址、名称、用户名不能为空")
			global.Cancel()
			return false
		}
	} else {
		err := c.refreshLocalDbConfig()
		if err != nil {
			return false
		}
	}
	return true
}

func (c *common) refreshLocalDbConfig() error {
	port := -1
	appType := ""
	clientType := ""

	switch object.RunMode(global.SysConfig.Task.RunMode) {
	case object.RunModeMd:
		port = 7083
		appType = "83"
		clientType = "8301"
	case object.RunModeGs:
		port = 7081
		appType = "81"
		clientType = "8101"
	case object.RunModeGc:
		port = 7082
		appType = "82"
		clientType = "8201"
	default:
		errMsg := fmt.Sprintf("unexpected runmode %d", global.SysConfig.Task.RunMode)
		log.Error(errMsg)
		global.Cancel()
		return errors.New(errMsg)
	}

	dbConfig, err := goToolSVRV3.GetSQLConfig(global.SysConfig.SvrV3Info.Address, port, appType, clientType)
	if err != nil {
		errMsg := fmt.Sprintf("get dbConfig from svr v3 err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}

	if dbConfig == nil {
		errMsg := fmt.Sprintf("get dbConfig from svr v3 return nil")
		log.Error(errMsg)
		return errors.New(errMsg)
	}

	accList, err := goToolSVRV3.GetAccountList(goToolMSSqlHelper.ConvertDbConfigTo2000(dbConfig), appType)
	if err != nil {
		errMsg := fmt.Sprintf("get acc list err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}

	if accList == nil || len(accList) <= 0 {
		errMsg := "acc list is empty"
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	global.SysConfig.LocalDb.Server = dbConfig.Server
	global.SysConfig.LocalDb.Port = dbConfig.Port
	global.SysConfig.LocalDb.User = dbConfig.User
	global.SysConfig.LocalDb.Pwd = dbConfig.Pwd

	if global.SysConfig.LocalDb.DbName != "" {
		flag := false
		for _, acc := range accList {
			if acc == global.SysConfig.LocalDb.DbName {
				log.Debug(acc)
				flag = true
				break
			}
		}
		if !flag {
			log.Warn(fmt.Sprintf("db [%s] is not a effective acc", global.SysConfig.LocalDb.DbName))
			global.SysConfig.LocalDb.DbName = ""
		}
	}
	if global.SysConfig.LocalDb.DbName == "" {
		global.SysConfig.LocalDb.DbName = accList[0]
	}
	return nil
}

func (c *common) StartService(sType object.RunMode) {
	log.Debug("RunMode " + strconv.Itoa(int(sType)))
	for {
		r := c.checkSysConfig()
		if r {
			break
		} else {
			time.Sleep(time.Minute * 30)
		}
	}
	log.Debug(fmt.Sprintf("dbName: %s", global.SysConfig.LocalDb.DbName))
	switch sType {
	case object.RunModeMd:
		c.addMdWorker()
	case object.RunModeGc:
		c.addWlWorker()
	case object.RunModeGs:
		c.addGsWorker()
	default:
		log.Warn(fmt.Sprintf("unknown runmode %v", sType))
		global.Cancel()
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

func (c *common) addGsWorker() {
	log.Debug("add gs worker")
	worker := NewGsWorker()
	c.addWorker("Z3Hpa", worker.Z3Hpa)
	c.addWorker("Z3HpDwFja", worker.Z3HpDwFja)
	c.addWorker("Z3JlDwa", worker.Z3JlDwa)
	c.addWorker("Z3HpEjFlt", worker.Z3HpEjFlt)
	c.addWorker("Z3KhDja", worker.Z3KhDja)
	c.addWorker("Z3JrDzqSzFja", worker.Z3JrDzqSzFja)
	c.addWorker("Z3FzJga", worker.Z3FzJga)
	c.addWorker("Z3DzqSza", worker.Z3DzqSza)
	c.addWorker("Z3Ppa", worker.Z3Ppa)
}

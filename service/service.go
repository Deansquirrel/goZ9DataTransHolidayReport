package service

import (
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/global"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/worker"
)

//启动服务内容
func StartService() error {
	log.Debug("Start Service")
	defer log.Debug("Start Service Complete")

	comm := worker.NewCommon()
	comm.StartService(object.RunMode(global.SysConfig.Task.RunMode))

	return nil
}

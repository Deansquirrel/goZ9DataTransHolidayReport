package worker

import (
	"fmt"
	"github.com/Deansquirrel/goServiceSupportHelper"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/repository"
)

type mdWorker struct {
}

func NewMdWorker() *mdWorker {
	return &mdWorker{}
}

//门店销售出库z3xsckt主表和z3xsckdt明细表
func (w *mdWorker) Z3XsCkt(jobId string) {
	log.Debug(fmt.Sprintf("Z3XsCkt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3XsCkt %s complete", jobId))
	repMd := repository.NewRepMd()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3XsCkt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	counter := 0
	for {
		rList, err := repMd.GetZ3XsCkDt()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3XsCkt get data error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, d := range rList {
			err = repOnLine.UpdateMdZ3XsCkDt(d)
			if err != nil {
				errMsg := fmt.Sprintf("update online Z3XsCkt ckmxhh[%s] err: %s", d.CkdMxHh, err.Error())
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
		}
		err = repMd.DelZ3XsCkDtSy(rList[0].CkdLsh)
		if err != nil {
			errMsg := fmt.Sprintf("del md Z3XsCkt CkdLsh[%s] err: %s", rList[0].CkdLsh, err.Error())
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		counter = counter + 1
		log.Debug(fmt.Sprintf("md Z3XsCkt[%s] Updated ", rList[0].CkdLsh))
	}
	if counter > 0 {
		log.Info(fmt.Sprintf("md Z3XsCkt Update %d", counter))
	}
}

//门店：门店销售退货z3xstht表和z3xsthhpdt明细表
func (w *mdWorker) Z3XsTht(jobId string) {
	log.Debug(fmt.Sprintf("Z3XsTht %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3XsTht %s complete", jobId))
	repMd := repository.NewRepMd()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3XsTht get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	counter := 0
	for {
		rList, err := repMd.GetZ3XsTht()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3XsTht get data error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, d := range rList {
			err = repOnLine.UpdateMdZ3XsTht(d)
			if err != nil {
				errMsg := fmt.Sprintf("update online Z3XsTht ThHpMxHh[%s] err: %s", d.ThHpMxHh, err.Error())
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
		}
		err = repMd.DelZ3XsThtSy(rList[0].ThLsh)
		if err != nil {
			errMsg := fmt.Sprintf("del md Z3XsTht ThLsh[%s] err: %s", rList[0].ThLsh, err.Error())
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		counter = counter + 1
		log.Debug(fmt.Sprintf("md Z3XsTht[%s] Updated ", rList[0].ThLsh))
	}
	if counter > 0 {
		log.Info(fmt.Sprintf("md Z3XsTht Update %d", counter))
	}
}

//门店：门店调拨出库z3mddbckt和z3mddbckdt明细表
func (w *mdWorker) Z3MdDbCkt(jobId string) {
	log.Debug(fmt.Sprintf("Z3MdDbCkt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3MdDbCkt %s complete", jobId))
	repMd := repository.NewRepMd()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3MdDbCkt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	counter := 0
	for {
		rList, err := repMd.GetZ3MdDbCkDt()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3MdDbCkt get data error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, d := range rList {
			err = repOnLine.UpdateMdZ3MdDbCkDt(d)
			if err != nil {
				errMsg := fmt.Sprintf("update online Z3MdDbCkt DbdMxHh[%s] err: %s", d.DbdMxHh, err.Error())
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
		}
		err = repMd.DelZ3MdDbCkDtSy(rList[0].DbdLsh)
		if err != nil {
			errMsg := fmt.Sprintf("del md Z3MdDbCkt ThLsh[%s] err: %s", rList[0].DbdLsh, err.Error())
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		counter = counter + 1
		log.Debug(fmt.Sprintf("md Z3MdDbCkt[%s] Updated ", rList[0].DbdLsh))
	}
	if counter > 0 {
		log.Info(fmt.Sprintf("md Z3MdDbCkt Update %d", counter))
	}
}

//门店：门店调拨调整z3dbtzt和z3dbtzdt明细表
func (w *mdWorker) Z3DbTzt(jobId string) {
	log.Debug(fmt.Sprintf("Z3DbTzt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3DbTzt %s complete", jobId))
	repMd := repository.NewRepMd()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3DbTzt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	counter := 0
	for {
		rList, err := repMd.GetZ3MdDbTzDt()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3DbTzt get data error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, d := range rList {
			err = repOnLine.UpdateMdZ3MdDbTzDt(d)
			if err != nil {
				errMsg := fmt.Sprintf("update online Z3DbTzt DbdMxHh[%s] err: %s", d.TzdMxHh, err.Error())
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
		}
		err = repMd.DelZ3MdDbTzDtSy(rList[0].TzdLsh)
		if err != nil {
			errMsg := fmt.Sprintf("del md Z3DbTzt TzdLsh[%s] err: %s", rList[0].TzdLsh, err.Error())
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		counter = counter + 1
		log.Debug(fmt.Sprintf("md Z3DbTzt[%s] Updated ", rList[0].TzdLsh))
	}
	if counter > 0 {
		log.Info(fmt.Sprintf("md Z3DbTzt Update %d", counter))
	}
}

//门店：门店其他出库z3hpckdjt和z3hpckdjdt明细表
func (w *mdWorker) Z3HpCkDjt(jobId string) {
	log.Debug(fmt.Sprintf("Z3HpCkDjt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3HpCkDjt %s complete", jobId))
	repMd := repository.NewRepMd()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3HpCkDjt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	counter := 0
	for {
		rList, err := repMd.GetZ3HpCkDjDt()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3HpCkDjt get data error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, d := range rList {
			err = repOnLine.UpdateMdZ3HpCkDjDt(d)
			if err != nil {
				errMsg := fmt.Sprintf("update online Z3HpCkDjt CkdMxHh[%s] err: %s", d.CkdMxHh, err.Error())
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
		}
		err = repMd.DelZ3HpCkDjDt(rList[0].CkdLsh)
		if err != nil {
			errMsg := fmt.Sprintf("del md Z3HpCkDjt CkdLsh[%s] err: %s", rList[0].CkdLsh, err.Error())
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		counter = counter + 1
		log.Debug(fmt.Sprintf("md Z3HpCkDjt[%s] Updated ", rList[0].CkdLsh))
	}
	if counter > 0 {
		log.Info(fmt.Sprintf("md Z3HpCkDjt Update %d", counter))
	}
}

//门店：门店其它入库z3hprkdjt和z3hprkdjdt明细表
func (w *mdWorker) Z3HpRkDjt(jobId string) {
	log.Debug(fmt.Sprintf("Z3HpRkDjt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3HpRkDjt %s complete", jobId))
	repMd := repository.NewRepMd()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3HpRkDjt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	counter := 0
	for {
		rList, err := repMd.GetZ3HpRkDjDt()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3HpRkDjt get data error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, d := range rList {
			err = repOnLine.UpdateMdZ3HpRkDjDt(d)
			if err != nil {
				errMsg := fmt.Sprintf("update online Z3HpRkDjt RkdMxHh[%s] err: %s", d.RkdMxHh, err.Error())
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
		}
		err = repMd.DelZ3HpRkDjDtSy(rList[0].RkdLsh)
		if err != nil {
			errMsg := fmt.Sprintf("del md Z3HpRkDjt RkdLsh[%s] err: %s", rList[0].RkdLsh, err.Error())
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		counter = counter + 1
		log.Debug(fmt.Sprintf("md Z3HpRkDjt[%s] Updated ", rList[0].RkdLsh))
	}
	if counter > 0 {
		log.Info(fmt.Sprintf("md Z3HpRkDjt Update %d", counter))
	}
}

//门店：门店盘亏登记z3pkdjt和z3pkdjdt明细表
func (w *mdWorker) Z3PkDjt(jobId string) {
	log.Debug(fmt.Sprintf("Z3PkDjt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3PkDjt %s complete", jobId))
	repMd := repository.NewRepMd()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3PkDjt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	counter := 0
	for {
		rList, err := repMd.GetZ3PkDjDt()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3PkDjt get data error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, d := range rList {
			err = repOnLine.UpdateMdZ3PkDjDt(d)
			if err != nil {
				errMsg := fmt.Sprintf("update online Z3PkDjt PkdMxHh[%s] err: %s", d.PkdMxHh, err.Error())
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
		}
		err = repMd.DelZ3PkDjDtSy(rList[0].PkdLsh)
		if err != nil {
			errMsg := fmt.Sprintf("del md Z3PkDjt PkdLsh[%s] err: %s", rList[0].PkdLsh, err.Error())
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		counter = counter + 1
		log.Debug(fmt.Sprintf("md Z3PkDjt[%s] Updated ", rList[0].PkdLsh))
	}
	if counter > 0 {
		log.Info(fmt.Sprintf("md Z3PkDjt Update %d", counter))
	}
}

//门店：门店配送收货调整z3pstzdt和z3pstzt
func (w *mdWorker) Z3PsTzDt(jobId string) {
	log.Debug(fmt.Sprintf("Z3PsTzDt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3PsTzDt %s complete", jobId))
	repMd := repository.NewRepMd()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3PsTzDt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	counter := 0
	for {
		rList, err := repMd.GetZ3MdPsTzDt()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3PsTzDt get data error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, d := range rList {
			err = repOnLine.UpdateMdZ3MdPsTzDt(d)
			if err != nil {
				errMsg := fmt.Sprintf("update online Z3PsTzDt TzdMxHh[%s] err: %s", d.TzdMxHh, err.Error())
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
		}
		err = repMd.DelZ3MdPsTzDtSy(rList[0].TzdLsh)
		if err != nil {
			errMsg := fmt.Sprintf("del md Z3PsTzDt PkdLsh[%s] err: %s", rList[0].TzdLsh, err.Error())
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		counter = counter + 1
		log.Debug(fmt.Sprintf("md Z3PsTzDt[%s] Updated ", rList[0].TzdLsh))
	}
	if counter > 0 {
		log.Info(fmt.Sprintf("md Z3PsTzDt Update %d", counter))
	}
}

//门店：门店提货单z3xsddthdt
func (w *mdWorker) Z3XsDdThDt(jobId string) {
	log.Debug(fmt.Sprintf("Z3XsDdThDt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3XsDdThDt %s complete", jobId))
	repMd := repository.NewRepMd()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3XsDdThDt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	counter := 0
	for {
		rList, err := repMd.GetZ3XsDdThDt()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3XsDdThDt get data error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, d := range rList {
			err = repOnLine.UpdateMdZ3XsDdThDt(d)
			if err != nil {
				errMsg := fmt.Sprintf("update online Z3XsDdThDt ThdMxHh[%s] err: %s", d.ThdMxHh, err.Error())
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
		}
		err = repMd.DelZ3XsDdThDtSy(rList[0].ThdLsh)
		if err != nil {
			errMsg := fmt.Sprintf("del md Z3XsDdThDt ThdLsh[%s] err: %s", rList[0].ThdLsh, err.Error())
			log.Error(errMsg)
			return
		}
		counter = counter + 1
		log.Debug(fmt.Sprintf("md Z3XsDdThDt[%s] Updated ", rList[0].ThdLsh))
	}
	if counter > 0 {
		log.Info(fmt.Sprintf("md Z3XsDdThDt Update %d", counter))
	}
}

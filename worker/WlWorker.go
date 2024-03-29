package worker

import (
	"fmt"
	"github.com/Deansquirrel/goServiceSupportHelper"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/repository"
	"time"
)

var xtTzLastUpdate time.Time

func init() {
	xtTzLastUpdate = goToolMSSqlHelper.GetDefaultOprTime()
}

type wlWorker struct {
}

func NewWlWorker() *wlWorker {
	return &wlWorker{}
}

//配送出库&配送出库二级明细表
func (w *wlWorker) Z3PsCkHpDt(jobId string) {
	log.Debug(fmt.Sprintf("Z3PsCkHpDt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3PsCkHpDt %s complete", jobId))
	repWl := repository.NewRepWl()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3PsCkHpDt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	udCounter := 0
	for {
		rList, err := repWl.GetZ3PsCkHpDtSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3PsCkHpDt get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repWl.GetZ3PsCkHpDt(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3PsCkHpDt get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3PsCkHpDt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			}
			ddList, err := repWl.GetZ3PsCkHpFjDt(id)
			if ddList == nil {
				errMsg := fmt.Sprintf("Z3PsCkHpFjDt get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(ddList) > 0 {
				for _, d := range ddList {
					err := repOnLine.UpdateZ3PsCkHpFjDt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				udCounter = udCounter + 1
			}
			err = repWl.DelZ3PsCkHpDtSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("wl Z3PsCkHpDt[%s] Update", id))
		}
	}
	if uCounter > 0 || udCounter > 0 {
		log.Info(fmt.Sprintf("wl Z3PsCkHpDt Update %d,Z3PsCkHpFjDt Update %d", uCounter, udCounter))
	}
}

//工厂配送修正 z3psxzt 和z3psxzdt
func (w *wlWorker) Z3PsXzDt(jobId string) {
	log.Debug(fmt.Sprintf("Z3PsXzDt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3PsXzDt %s complete", jobId))
	repWl := repository.NewRepWl()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3PsXzDt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	for {
		rList, err := repWl.GetZ3PsXzDtSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3PsXzDt get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repWl.GetZ3PsXzDt(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3PsXzDt get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3PsXzDt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			}
			err = repWl.DelZ3PsXzDtSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("wl Z3PsXzDt[%s] Update", id))
		}
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("wl Z3PsXzDt Update %d", uCounter))
	}
}

//工厂配送冲红z3mdpscht 和z3mdpschdt
func (w *wlWorker) Z3MdPsChDt(jobId string) {
	log.Debug(fmt.Sprintf("Z3MdPsChDt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3MdPsChDt %s complete", jobId))
	repWl := repository.NewRepWl()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3MdPsChDt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	for {
		rList, err := repWl.GetZ3MdPsChDtSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3MdPsChDt get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repWl.GetZ3MdPsChDt(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3MdPsChDt get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3MdPsChDt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			}
			err = repWl.DelZ3MdPsChDtSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("wl Z3MdPsChDt[%s] Update", id))
		}
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("wl Z3MdPsChDt Update %d", uCounter))
	}
}

//工厂赊账销售出库z3shezhxsckt和z3shezhxshpmxt
func (w *wlWorker) Z3SheZhXsHpMxt(jobId string) {
	log.Debug(fmt.Sprintf("Z3SheZhXsHpMxt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3SheZhXsHpMxt %s complete", jobId))
	repWl := repository.NewRepWl()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3SheZhXsHpMxt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	for {
		rList, err := repWl.GetZ3SheZhXsHpMxtSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3SheZhXsHpMxt get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repWl.GetZ3SheZhXsHpMxt(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3SheZhXsHpMxt get data error: return list can not be nil")
				log.Error(errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3SheZhXsHpMxt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			}
			err = repWl.DelZ3SheZhXsHpMxtSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("wl Z3SheZhXsHpMxt[%s] Update", id))
		}
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("wl Z3SheZhXsHpMxt Update %d", uCounter))
	}
}

//工厂完工入库shengchwgrkdt 和shengchwgrkt
func (w *wlWorker) ShengChWgRkDt(jobId string) {
	log.Debug(fmt.Sprintf("ShengChWgRkDt %s start", jobId))
	defer log.Debug(fmt.Sprintf("ShengChWgRkDt %s complete", jobId))
	repWl := repository.NewRepWl()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("ShengChWgRkDt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	for {
		rList, err := repWl.GetShengChWgRkDtSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("ShengChWgRkDt get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repWl.GetShengChWgRkDt(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("ShengChWgRkDt get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateShengChWgRkDt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			}
			err = repWl.DelShengChWgRkDtSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("wl ShengChWgRkDt[%s] Update", id))
		}
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("wl ShengChWgRkDt Update %d", uCounter))
	}
}

//工厂即时台账表xttz
func (w *wlWorker) XtTz(jobId string) {
	log.Debug(fmt.Sprintf("XtTz %s start", jobId))
	defer log.Debug(fmt.Sprintf("XtTz %s complete", jobId))
	repWl := repository.NewRepWl()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("XtTz get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	currTime := xtTzLastUpdate.Add(-time.Second)
	lastTime := goToolMSSqlHelper.GetDefaultOprTime()
	rList, err := repWl.GetXtTz(currTime)
	if err != nil {
		_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
		return
	}
	if rList == nil {
		errMsg := fmt.Sprintf("XtTz get data error: return list can not be nil")
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	for _, tx := range rList {
		err := repOnLine.UpdateXtTz(tx)
		if err != nil {
			log.Error(err.Error())
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if goToolCommon.GetDateTimeStrWithMillisecond(tx.UpdateDate) > goToolCommon.GetDateTimeStrWithMillisecond(lastTime) {
			lastTime = tx.UpdateDate
		}
		uCounter = uCounter + 1
	}
	if goToolCommon.GetDateTimeStrWithMillisecond(lastTime) > goToolCommon.GetDateTimeStrWithMillisecond(xtTzLastUpdate) {
		xtTzLastUpdate = lastTime
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("wl XtTz Update %d, lastUpdateTime: %s",
			uCounter, goToolCommon.GetDateTimeStrWithMillisecond(xtTzLastUpdate)))
	}
}

//门店订货单
func (w *wlWorker) Z3MdDhDdt(jobId string) {
	log.Debug(fmt.Sprintf("Z3MdDhDdt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3MdDhDdt %s complete", jobId))
	repWl := repository.NewRepWl()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3MdDhDdt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	for {
		rList, err := repWl.GetZ3MdDhDdtSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3MdDhDdt get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repWl.GetZ3MdDhDdt(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3MdDhDdt get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3MdDhDdt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			}
			err = repWl.DelZ3MdDhDdtSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("wl Z3MdDhDdt[%s] Update", id))
		}
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("wl Z3MdDhDdt Update %d", uCounter))
	}
}

//工厂订单提货单oodxv1ddshckt
func (w *wlWorker) OodXv1DdShCkt(jobId string) {
	log.Debug(fmt.Sprintf("OodXv1DdShCkt %s start", jobId))
	defer log.Debug(fmt.Sprintf("OodXv1DdShCkt %s complete", jobId))
	repWl := repository.NewRepWl()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("OodXv1DdShCkt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	for {
		rList, err := repWl.GetOoDxV1DdShCktSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("OodXv1DdShCkt get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repWl.GetOoDxV1DdShCkt(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("OodXv1DdShCkt get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateOoDxV1DdShCkt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			}
			err = repWl.DelOoDxV1DdShCktSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("wl OodXv1DdShCkt[%s] Update", id))
		}
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("wl OodXv1DdShCkt Update %d", uCounter))
	}
}

//门店订货代单
func (w *wlWorker) Z3MdDhDdDt(jobId string) {
	log.Debug(fmt.Sprintf("Z3MdDhDdDt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3MdDhDdDt %s complete", jobId))
	repWl := repository.NewRepWl()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3MdDhDdDt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	for {
		rList, err := repWl.GetZ3MdDhDdDtSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3MdDhDdDt get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repWl.GetZ3MdDhDdDt(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3MdDhDdDt get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3MdDhDdDt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			}
			err = repWl.DelZ3MdDhDdDtSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("wl Z3MdDhDdDt[%s] Update", id))
		}
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("wl Z3MdDhDdDt Update %d", uCounter))
	}
}

package worker

import (
	"fmt"
	"github.com/Deansquirrel/goServiceSupportHelper"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/repository"
)

type gsWorker struct {
}

func NewGsWorker() *gsWorker {
	return &gsWorker{}
}

//集团通用货品设置A
func (w *gsWorker) Z3Hpa(id string) {
	log.Debug(fmt.Sprintf("Z3Hpa %s start", id))
	defer log.Debug(fmt.Sprintf("Z3Hpa %s complete", id))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3Hpa get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(id, errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3HpaSy()
		if err != nil {
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3Hpa get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(id, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, sy := range rList {
			dList, err := repGs.GetZ3Hpa(sy)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3Hpa get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3Hpa(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3Hpa(sy)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3HpaSy(sy)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("gs Z3Hpa[%d] Update", sy))
		}
	}
	if uCounter > 0 || dCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3Hpa Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//货品计量单位附加表
func (w *gsWorker) Z3HpDwFja(jobId string) {
	log.Debug(fmt.Sprintf("Z3HpDwFja %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3HpDwFja %s complete", jobId))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3HpDwFja get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3HpDwFjaSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3HpDwFja get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3HpDwFja(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3HpDwFja get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3HpDwFja(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3HpDwFja(id.DwFjHpId, id.DwFjDwId)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3HpDwFjaSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("gs Z3HpDwFja[%d_%d] Update", id.DwFjHpId, id.DwFjDwId))
		}
	}
	if uCounter > 0 || dCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3HpDwFja Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//计量单位设置
func (w *gsWorker) Z3JlDwa(jobId string) {
	log.Debug(fmt.Sprintf("Z3JlDwa %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3JlDwa %s complete", jobId))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3JlDwa get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3JlDwaSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3JlDwa get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3JlDwa(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3JlDwa get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3JlDwa(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3JlDwa(id)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3JlDwaSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("gs Z3JlDwa[%d] Update", id))
		}
	}
	if uCounter > 0 || dCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3JlDwa Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//货品二级分类
func (w *gsWorker) Z3HpEjFlt(jobId string) {
	log.Debug(fmt.Sprintf("Z3HpEjFlt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3HpEjFlt %s complete", jobId))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3HpEjFlt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3HpEjFltSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3HpEjFlt get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3HpEjFlt(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3HpEjFlt get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3HpEjFlt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3HpEjFlt(id)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3HpEjFltSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("gs Z3HpEjFlt[%d] Update", id))
		}
	}
	if uCounter > 0 || dCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3HpEjFlt Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//客户登记通用信息表
func (w *gsWorker) Z3KhDja(jobId string) {
	log.Debug(fmt.Sprintf("Z3KhDja %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3KhDja %s complete", jobId))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3KhDja get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3KhDjaSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3KhDja get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3KhDja(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3KhDja get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3KhDja(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3KhDja(id)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3KhDjaSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("gs Z3KhDja[%d] Update", id))
		}
	}
	if uCounter > 0 || dCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3KhDja Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//节日电子券设置附加表
func (w *gsWorker) Z3JrDzqSzFja(jobId string) {
	log.Debug(fmt.Sprintf("Z3JrDzqSzFja %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3JrDzqSzFja %s complete", jobId))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3JrDzqSzFja get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3JrDzqSzFjaSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3JrDzqSzFja get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3JrDzqSzFja(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3JrDzqSzFja get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3JrDzqSzFja(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3JrDzqSzFja(id)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3JrDzqSzFjaSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("gs Z3JrDzqSzFja[%d] Update", id))
		}
	}
	if uCounter > 0 || dCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3JrDzqSzFja Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//机构表V/A
func (w *gsWorker) Z3FzJga(jobId string) {
	log.Debug(fmt.Sprintf("Z3FzJga %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3FzJga %s complete", jobId))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3FzJga get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3FzJgaSy()
		if err != nil {
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3FzJga get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3FzJga(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3FzJga get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3FzJga(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3FzJga(id)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3FzJgaSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("gs Z3FzJga[%d] Update", id))
		}
	}
	if uCounter > 0 || dCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3FzJga Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//核心组织零售价
func (w *gsWorker) Z3HxZzLsjt(jobId string) {
	log.Debug(fmt.Sprintf("Z3HxZzLsjt %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3HxZzLsjt %s complete", jobId))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3HxZzLsjt get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3HxZzLsjtSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3HxZzLsjt get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3HxZzLsjt(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3HxZzLsjt get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3HxZzLsjt(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3HxZzLsjt(id)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3HxZzLsjtSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("gs Z3HxZzLsjt[%d,%d] Update", id.LsjZzId, id.LsjHpId))
		}
	}
	if uCounter > 0 || dCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3HxZzLsjt Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//电子券设置
func (w *gsWorker) Z3DzqSza(jobId string) {
	log.Debug(fmt.Sprintf("Z3DzqSza %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3DzqSza %s complete", jobId))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3DzqSza get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3DzqSzaSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3DzqSza get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3DzqSza(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3DzqSza get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3DzqSza(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3DzqSza(id)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3DzqSzaSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("gs Z3DzqSza[%d] Update", id))
		}
	}
	if uCounter > 0 || dCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3DzqSza Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//货品品牌设置
func (w *gsWorker) Z3Ppa(jobId string) {
	log.Debug(fmt.Sprintf("Z3Ppa %s start", jobId))
	defer log.Debug(fmt.Sprintf("Z3Ppa %s complete", jobId))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3Ppa get rep online err: %s", err.Error())
		log.Error(errMsg)
		_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3PpaSy()
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3Ppa get sy error: return list can not be nil")
			log.Error(errMsg)
			_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3Ppa(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3Ppa get data error: return list can not be nil")
				log.Error(errMsg)
				_ = goServiceSupportHelper.JobErrRecord(jobId, errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3Ppa(d)
					if err != nil {
						_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3Ppa(id)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3PpaSy(id)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(jobId, err.Error())
				return
			}
			log.Debug(fmt.Sprintf("gs Z3Ppa[%d] Update", id))
		}
	}
	if uCounter > 0 || dCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3Ppa Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

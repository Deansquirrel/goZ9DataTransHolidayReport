package worker

import (
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/repository"
)

type gsWorker struct {
}

func NewGsWorker() *gsWorker {
	return &gsWorker{}
}

//集团通用货品设置A
func (w *gsWorker) Z3Hpa() {
	id := goToolCommon.Guid()
	log.Debug(fmt.Sprintf("Z3Hpa %s start", id))
	defer log.Debug(fmt.Sprintf("Z3Hpa %s complete", id))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3Hpa get rep online err: %s", err.Error())
		log.Error(errMsg)
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
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3Hpa(id)
			if err != nil {
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3Hpa get data error: return list can not be nil")
				log.Error(errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3Hpa(d)
					if err != nil {
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3Hpa(id)
				if err != nil {
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3HpaSy(id)
			if err != nil {
				return
			}
			log.Debug(fmt.Sprintf("gs Z3Hpa[%d] Update", id))
		}
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3Hpa Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//货品计量单位附加表
func (w *gsWorker) Z3HpDwFja() {
	id := goToolCommon.Guid()
	log.Debug(fmt.Sprintf("Z3HpDwFja %s start", id))
	defer log.Debug(fmt.Sprintf("Z3HpDwFja %s complete", id))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3HpDwFja get rep online err: %s", err.Error())
		log.Error(errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3HpDwFjaSy()
		if err != nil {
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3HpDwFja get sy error: return list can not be nil")
			log.Error(errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3HpDwFja(id)
			if err != nil {
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3HpDwFja get data error: return list can not be nil")
				log.Error(errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3HpDwFja(d)
					if err != nil {
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3HpDwFja(id.DwFjHpId, id.DwFjDwId)
				if err != nil {
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3HpDwFjaSy(id)
			if err != nil {
				return
			}
			log.Debug(fmt.Sprintf("gs Z3HpDwFja[%d_%d] Update", id.DwFjHpId, id.DwFjDwId))
		}
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3HpDwFja Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//计量单位设置
func (w *gsWorker) Z3JlDwa() {
	id := goToolCommon.Guid()
	log.Debug(fmt.Sprintf("Z3JlDwa %s start", id))
	defer log.Debug(fmt.Sprintf("Z3JlDwa %s complete", id))
	repGs := repository.NewRepGs()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3JlDwa get rep online err: %s", err.Error())
		log.Error(errMsg)
		return
	}
	uCounter := 0
	dCounter := 0
	for {
		rList, err := repGs.GetZ3JlDwaSy()
		if err != nil {
			return
		}
		if rList == nil {
			errMsg := fmt.Sprintf("Z3JlDwa get sy error: return list can not be nil")
			log.Error(errMsg)
			return
		}
		if len(rList) == 0 {
			break
		}
		for _, id := range rList {
			dList, err := repGs.GetZ3JlDwa(id)
			if err != nil {
				return
			}
			if dList == nil {
				errMsg := fmt.Sprintf("Z3JlDwa get data error: return list can not be nil")
				log.Error(errMsg)
				return
			}
			if len(dList) > 0 {
				for _, d := range dList {
					err := repOnLine.UpdateZ3JlDwa(d)
					if err != nil {
						return
					}
				}
				uCounter = uCounter + 1
			} else {
				err := repOnLine.DelZ3JlDwa(id)
				if err != nil {
					return
				}
				dCounter = dCounter + 1
			}
			err = repGs.DelZ3JlDwaSy(id)
			if err != nil {
				return
			}
			log.Debug(fmt.Sprintf("gs Z3JlDwa[%d] Update", id))
		}
	}
	if uCounter > 0 {
		log.Info(fmt.Sprintf("gs Z3JlDwa Update %d,Del %d,Total %d", uCounter, dCounter, uCounter+dCounter))
	}
}

//货品计量单位附加表
//计量单位设置（A）
//货品二级分类t
//客户登记通用信息表
//节日电子券设置附加表
//机构表V/A
//核心组织零售价

package worker

import (
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/repository"
)

type mdWorker struct {
}

func NewMdWorker() *mdWorker {
	return &mdWorker{}
}

//门店销售出库z3xsckt主表和z3xsckdt明细表
func (w *mdWorker) Z3XsCkt() {
	repMd := repository.NewRepMd()
	repOnLine, err := repository.NewRepZxZc()
	if err != nil {
		errMsg := fmt.Sprintf("Z3XsCkt get rep online err: %s", err.Error())
		log.Error(errMsg)
		return
	}
	counter := 0
	for {
		rList, err := repMd.GetZ3XsCkDt()
		if err != nil {
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
				return
			}
		}
		err = repMd.DeleteZ3XsCkDtSy(rList[0].CkdLsh)
		if err != nil {
			errMsg := fmt.Sprintf("del md Z3XsCkt ckmxhh[%s] err: %s", rList[0].CkdMxHh, err.Error())
			log.Error(errMsg)
			return
		}
		counter = counter + 1
	}
	if counter > 0 {
		log.Info(fmt.Sprintf("md Z3XsCkt Update %d", counter))
	}
}

//门店：门店销售退货z3xstht表和z3xsthhpdt明细表
func (w *mdWorker) Z3XsTht() {
	log.Debug("Z3XsTht")
}

//门店：门店调拨出库z3mddbckt和z3mddbckdt明细表
func (w *mdWorker) Z3MdDbCkt() {
	log.Debug("Z3MdDbCkt")
}

//门店：门店调拨调整z3dbtzt和z3dbtzdt明细表
func (w *mdWorker) Z3DbTzt() {
	log.Debug("Z3DbTzt")
}

//门店：门店其他出库z3hpckdjt和z3hpckdjdt明细表
func (w *mdWorker) Z3HpCkDjt() {
	log.Debug("Z3HpCkDjt")
}

//门店：门店其它入库z3hprkdjt和z3hprkdjdt明细表
func (w *mdWorker) Z3HpRkDjt() {
	log.Debug("Z3HpRkDjt")
}

//门店：门店盘亏登记z3pkdjt和z3pkdjdt明细表
func (w *mdWorker) Z3PkDjt() {
	log.Debug("Z3PkDjt")
}

//门店：门店配送收货调整z3pstzdt和z3pstzt
func (w *mdWorker) Z3PsTzDt() {
	log.Debug("Z3PsTzDt")
}

//门店：门店提货单z3xsddthdt
func (w *mdWorker) Z3XsDdThDt() {
	log.Debug("Z3XsDdThDt")
}

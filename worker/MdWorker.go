package worker

import log "github.com/Deansquirrel/goToolLog"

type mdWorker struct {
}

func NewMdWorker() *mdWorker {
	return &mdWorker{}
}

//门店销售出库z3xsckt主表和z3xsckdt明细表
func (w *mdWorker) Z3XsCkt() {
	log.Debug("Z3XsCkt")
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

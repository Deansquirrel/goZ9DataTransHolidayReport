package worker

import log "github.com/Deansquirrel/goToolLog"

type wlWorker struct {
}

func NewWlWorker() *wlWorker {
	return &wlWorker{}
}

//工厂配送单z3mdpsckdt和z3mdpsckfjt，z3mdpsckt
func (w *wlWorker) Z3MdPsCkDt() {
	log.Debug("Z3MdPsCkDt")
}

//工厂配送修正 z3psxzt 和z3psxzdt
func (w *wlWorker) Z3PsXzt() {
	log.Debug("Z3PsXzt")
}

//工厂配送冲红z3mdpscht 和z3mdpschdt
func (w *wlWorker) Z3MdPsCht() {
	log.Debug("Z3MdPsCht")
}

//工厂赊账销售出库z3shezhxsckt和z3shezhxshpmxt
func (w *wlWorker) Z3SheZhXsCkt() {
	log.Debug("Z3SheZhXsCkt")
}

//工厂完工入库shengchwgrkdt 和shengchwgrkt
func (w *wlWorker) ShengChWgRkt() {
	log.Debug("ShengChWgRkt")
}

//工厂即时台账表xttz
func (w *wlWorker) XtTz() {
	log.Debug("XtTz")
}

//工厂代门店订货单z3mddhdt
func (w *wlWorker) Z3MdDhDt() {
	log.Debug("Z3MdDhDt")
}

//工厂订货单修改z3mddhdt_md和z3mddht_md
func (w *wlWorker) Z3MdDhDtMd() {
	log.Debug("Z3MdDhDtMd")
}

//工厂订单提货单oodxv1ddshckt
func (w *wlWorker) OodXv1DdShCkt() {
	log.Debug("OodXv1DdShCkt")
}

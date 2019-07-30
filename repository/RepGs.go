package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goToolMSSql2000"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
)

import log "github.com/Deansquirrel/goToolLog"

const (
	sqlGetZ3HpaSy = "" +
		"select top 1 hpid " +
		"from z3hpsyt"
	sqlGetZ3Hpa = "" +
		"select [hpid],[hpbm],[hpszmc],[hpppid],[hpsym], " +
		"	[hppym],[hpzxdwid],[hpzxdwmc],[hpejflid],[hpmdszbj], " +
		"	[hpdkhdhhpbj],[hpwddhhpbj],[hpms],[hpzybj],[hpisforbidden], " +
		"	[hpszid],[hpnsflm],[hpjrhp],[hpdghp] " +
		"from [z3hpa] " +
		"where [hpid] = ?"
	sqlDelZ3HpaSy = "" +
		"delete from z3hpsyt " +
		"where hpid = ?"

	sqlGetZ3HpDwFjaSy = "" +
		"select top 1 [dwfjhpid],[dwfjdwid] " +
		"from [z3hpdwfjsyt]"
	sqlGetZ3HpDwFja = "" +
		"select [dwfjhpid],[dwfjdwid],[dwfjdwlx],[dwfjhsl],[dwfjtxm],[dwfjisforbidden] " +
		"from [z3hpdwfja] " +
		"where [dwfjhpid] = ? and [dwfjdwid] = ?"
	sqlDelZ3HpDwFjaSy = "" +
		"delete from [z3hpdwfjsyt] " +
		"where [dwfjhpid] = ? and [dwfjdwid] = ?"

	sqlGetZ3JlDwaSy = "" +
		"select top 1 [dwid] " +
		"from [z3jldwsyt]"
	sqlGetZ3JlDwa = "" +
		"select [dwid],[dwszmc],[dwglmc],[dwsym],[dwpym],[dwisforbidden] " +
		"from [z3jldwa] " +
		"where [dwid] = ?"
	sqlDelZ3JlDwaSy = "" +
		"delete from [z3jldwsyt] " +
		"where [dwid] = ?"

	sqlGetZ3HpEjFltSy = "" +
		"select top 1 [ejflid] " +
		"from [z3hpejflsyt]"
	sqlGetZ3HpEjFlt = "" +
		"select [ejflid],[ejflmc],[ejflbm],[ejflppid],[ejflsym], " +
		"	[ejflpym],[ejflhpyjflid] " +
		"from [z3hpejflt] " +
		"where [ejflid] = ?"
	sqlDelZ3HpEjFltSy = "" +
		"delete from [z3hpejflsyt] " +
		"where [ejflid] = ?"
)

type repGs struct {
	dbConfig *goToolMSSql2000.MSSqlConfig
}

func NewRepGs() *repGs {
	comm := NewCommon()
	return &repGs{
		dbConfig: comm.ConvertDbConfigTo2000(comm.GetLocalDbConfig()),
	}
}

//获取集团通用货品设置A索引
func (r *repGs) GetZ3HpaSy() ([]int64, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpaSy)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]int64, 0)
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpaSy", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpaSy", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//获取集团通用货品设置A
func (r *repGs) GetZ3Hpa(id int64) ([]*object.Z3Hpa, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3Hpa, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var hpId, hpPpId, hpZxDwId, hpEjFlId, hpMdSzBj, hpdDkhDhHpBj, hpWdDhHpBj, hpZyNj, hpIsForbidden, hpSzId, hpJrHp, hpDgHp int
	var hpBm, hpSzMc, hpSym, hpPym, hpZxDwMc, hpMs, hpNsFlm string
	rList := make([]*object.Z3Hpa, 0)
	for rows.Next() {
		err := rows.Scan(&hpId, &hpBm, &hpSzMc, &hpPpId, &hpSym,
			&hpPym, &hpZxDwId, &hpZxDwMc, &hpEjFlId, &hpMdSzBj,
			&hpdDkhDhHpBj, &hpWdDhHpBj, &hpMs, &hpZyNj, &hpIsForbidden,
			&hpSzId, &hpNsFlm, &hpJrHp, &hpDgHp)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3Hpa", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3Hpa{
			HpId:          hpId,
			HpBm:          hpBm,
			HpSzMc:        hpSzMc,
			HpPpId:        hpPpId,
			HpSym:         hpSym,
			HpPym:         hpPym,
			HpZxDwId:      hpZxDwId,
			HpZxDwMc:      hpZxDwMc,
			HpEjFlId:      hpEjFlId,
			HpMdSzBj:      hpMdSzBj,
			HpdDkhDhHpBj:  hpdDkhDhHpBj,
			HpWdDhHpBj:    hpWdDhHpBj,
			HpMs:          hpMs,
			HpZyNj:        hpZyNj,
			HpIsForbidden: hpIsForbidden,
			HpSzId:        hpSzId,
			HpNsFlm:       hpNsFlm,
			HpJrHp:        hpJrHp,
			HpDgHp:        hpDgHp,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3Hpa", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除集团通用货品设置A索引
func (r *repGs) DelZ3HpaSy(id int64) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDelZ3HpaSy, id)
}

//获取货品计量单位附加表索引
func (r *repGs) GetZ3HpDwFjaSy() ([]*object.Z3HpDwFjSyt, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpDwFjaSy)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]*object.Z3HpDwFjSyt, 0)
	var hpId, dwId int
	for rows.Next() {
		err := rows.Scan(&hpId, &dwId)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpDwFjaSy", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3HpDwFjSyt{
			DwFjHpId: hpId,
			DwFjDwId: dwId,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpaSy", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//获取货品计量单位附加表
func (r *repGs) GetZ3HpDwFja(sy *object.Z3HpDwFjSyt) ([]*object.Z3HpDwFja, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpDwFja, sy.DwFjHpId, sy.DwFjDwId)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var dwFjHpId, dwFjDwId, dwFjDwLx, dwFjIsForbidden int
	var dwFjHsl float64
	var dwFjTxm string
	rList := make([]*object.Z3HpDwFja, 0)
	for rows.Next() {
		err := rows.Scan(&dwFjHpId, &dwFjDwId, &dwFjDwLx, &dwFjHsl, &dwFjTxm, &dwFjIsForbidden)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpDwFja", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3HpDwFja{
			DwFjHpId:        dwFjHpId,
			DwFjDwId:        dwFjDwId,
			DwFjDwLx:        dwFjDwLx,
			DwFjHsl:         dwFjHsl,
			DwFjTxm:         dwFjTxm,
			DwFjIsForbidden: dwFjIsForbidden,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpDwFja", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除货品计量单位附加表索引
func (r *repGs) DelZ3HpDwFjaSy(sy *object.Z3HpDwFjSyt) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDelZ3HpDwFjaSy, sy.DwFjHpId, sy.DwFjDwId)
}

//获取计量单位设置索引
func (r *repGs) GetZ3JlDwaSy() ([]int64, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3JlDwaSy)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]int64, 0)
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JlDwaSy", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JlDwaSy", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//获取计量单位设置
func (r *repGs) GetZ3JlDwa(id int64) ([]*object.Z3JlDwa, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3JlDwa, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var dwId, dwIsForbidden int
	var dwSzMc, dwGlMc, dwSym, dwPym string

	rList := make([]*object.Z3JlDwa, 0)
	for rows.Next() {
		err := rows.Scan(&dwId, &dwSzMc, &dwGlMc, &dwSym, &dwPym, &dwIsForbidden)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JlDwa", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3JlDwa{
			DwId:          dwId,
			DwSzMc:        dwSzMc,
			DwGlMc:        dwGlMc,
			DwSym:         dwSym,
			DwPym:         dwPym,
			DwIsForbidden: dwIsForbidden,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JlDwa", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除计量单位设置索引
func (r *repGs) DelZ3JlDwaSy(id int64) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDelZ3JlDwaSy, id)
}

//获取货品二级分类索引
func (r *repGs) GetZ3HpEjFltSy() ([]int64, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpEjFltSy)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]int64, 0)
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpEjFltSy", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpEjFltSy", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//获取货品二级分类
func (r *repGs) GetZ3HpEjFlt(id int64) ([]*object.Z3HpEjFlt, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpEjFlt, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var ejFlId, ejFlPpId, ejFlHpYjFlId int
	var ejFlMc, ejFlBm, ejFlSym, ejFlPym string

	rList := make([]*object.Z3HpEjFlt, 0)
	for rows.Next() {
		err := rows.Scan(&ejFlId, &ejFlMc, &ejFlBm, &ejFlPpId, &ejFlSym,
			&ejFlPym, &ejFlHpYjFlId)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpEjFlt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3HpEjFlt{
			EjFlId:       ejFlId,
			EjFlMc:       ejFlMc,
			EjFlBm:       ejFlBm,
			EjFlPpId:     ejFlPpId,
			EjFlSym:      ejFlSym,
			EjFlPym:      ejFlPym,
			EjFlHpYjFlId: ejFlHpYjFlId,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpEjFlt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除货品二级分类索引
func (r *repGs) DelZ3HpEjFltSy(id int64) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDelZ3HpEjFltSy, id)
}

package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goToolMSSql2000"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
	"time"
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

	sqlGetZ3KhDjaSy = "" +
		"select top 1 [khid] " +
		"from [z3khdjsyt]"
	sqlGetZ3KhDja = "" +
		"select [khid],[khszmc],[khqc],[khbm],[khnbygbj], " +
		"	[khxb],[khpym],[khjb],[khsjhm],[khqtlsfs], " +
		"	[khkhyt],[khrcxyj],[khrcxyjsydd],[khrcyh],[khflid], " +
		"	[khlszz],[khisforbidden] " +
		"from [z3khdja] " +
		"where [khid] = ?"
	sqlDelZ3KhDjaSy = "" +
		"delete from [z3khdjsyt] " +
		"where [khid] = ?"

	sqlGetZ3JrDzqSzFjaSy = "" +
		"select top 1 [dzqid] " +
		"from [z3jrdzqszfjsyt]"
	sqlGetZ3JrDzqSzFja = "" +
		"select [dzqid],[dzqdpqhpid],[dzqjybz] " +
		"from [z3jrdzqszfja] " +
		"where [dzqid]=?"
	sqlDelZ3JrDzqSzFjaSy = "" +
		"delete from [z3jrdzqszfjsyt] " +
		"where [dzqid] = ?"

	sqlGetZ3FzJgaSy = "" +
		"select top 1 [fzjgid] " +
		"from [z3fzjgsyt]"
	sqlGetZ3FzJga = "" +
		"select [fzjgid],[fzjgszmc],[fzjgqc],[fzjglx],[fzjgsxrq], " +
		"	[fzjgpym],[fzjgsym],[fzjgtdm],[fzjgxsm],[fzjglszz], " +
		"	[fzjgisforbidden] " +
		"from [z3fzjga] " +
		"where [fzjgid] = ?"
	sqlDelZ3FzJgaSy = "" +
		"delete from [z3fzjgsyt] " +
		"where [fzjgid] = ?"

	sqlGetZ3DzqSzaSy = "" +
		"select top 1 [dzqid] " +
		"from [z3dzqszsyt]"
	sqlGetZ3DzqSza = "" +
		"select [dzqid],[dzqszmc],[dzqdwxsmc],[dzqbm],[dzqpym], " +
		"	[dzqjrqbj],[dzqzzqmz],[dzqyt],[dzqhxhly],[dzqdbbs], " +
		"	[dzqhsxz],[dzqctqid],[dzqzdzkl],[dzqqxzkyh],[dzqqhsjzr], " +
		"	[dzqjybz] " +
		"from [z3dzqsza] " +
		"where [dzqid] = ?"
	sqlDelZ3DzqSzaSy = "" +
		"delete from [z3dzqszsyt] " +
		"where [dzqid] = ?"
)

type repGs struct {
	dbConfig *goToolMSSql2000.MSSqlConfig
}

func NewRepGs() *repGs {
	comm := NewCommon()
	return &repGs{
		dbConfig: goToolMSSqlHelper.ConvertDbConfigTo2000(comm.GetLocalDbConfig()),
	}
}

//获取集团通用货品设置A索引
func (r *repGs) GetZ3HpaSy() ([]int64, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpaSy)
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
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3Hpa, id)
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
	return goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3HpaSy, id)
}

//获取货品计量单位附加表索引
func (r *repGs) GetZ3HpDwFjaSy() ([]*object.Z3HpDwFjSyt, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpDwFjaSy)
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
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpDwFja, sy.DwFjHpId, sy.DwFjDwId)
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
	return goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3HpDwFjaSy, sy.DwFjHpId, sy.DwFjDwId)
}

//获取计量单位设置索引
func (r *repGs) GetZ3JlDwaSy() ([]int64, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3JlDwaSy)
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
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3JlDwa, id)
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
	return goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3JlDwaSy, id)
}

//获取货品二级分类索引
func (r *repGs) GetZ3HpEjFltSy() ([]int64, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpEjFltSy)
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
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpEjFlt, id)
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
	return goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3HpEjFltSy, id)
}

//获取客户登记通用信息表索引
func (r *repGs) GetZ3KhDjaSy() ([]int64, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3KhDjaSy)
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
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3KhDjaSy", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3KhDjaSy", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//获取客户登记通用信息表
func (r *repGs) GetZ3KhDja(id int64) ([]*object.Z3KhDja, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3KhDja, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var khId, khNbYgBj, khXb, khJb, khQtLsFs, khKHyt, khRcXyj, khRcXyjSyDd, khFlId, khIsForbidden int
	var khSzMc, khQc, khBm, khPym, khSjHm, khLsZz string
	var khrRcYh float64

	rList := make([]*object.Z3KhDja, 0)
	for rows.Next() {
		err := rows.Scan(&khId, &khSzMc, &khQc, &khBm, &khNbYgBj,
			&khXb, &khPym, &khJb, &khSjHm, &khQtLsFs,
			&khKHyt, &khRcXyj, &khRcXyjSyDd, &khrRcYh, &khFlId,
			&khLsZz, &khIsForbidden)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3KhDja", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3KhDja{
			KhId:          khId,
			KhSzMc:        khSzMc,
			KhQc:          khQc,
			KhBm:          khBm,
			KhNbYgBj:      khNbYgBj,
			KhXb:          khXb,
			KhPym:         khPym,
			KhJb:          khJb,
			KhSjHm:        khSjHm,
			KhQtLsFs:      khQtLsFs,
			KhKHyt:        khKHyt,
			KhRcXyj:       khRcXyj,
			KhRcXyjSyDd:   khRcXyjSyDd,
			KhrRcYh:       khrRcYh,
			KhFlId:        khFlId,
			KhLsZz:        khLsZz,
			KhIsForbidden: khIsForbidden,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3KhDja", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除客户登记通用信息表索引
func (r *repGs) DelZ3KhDjaSy(id int64) error {
	return goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3KhDjaSy, id)
}

//获取节日电子券设置附加表索引
func (r *repGs) GetZ3JrDzqSzFjaSy() ([]int64, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3JrDzqSzFjaSy)
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
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JrDzqSzFjaSy", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JrDzqSzFjaSy", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//获取节日电子券设置附加表
func (r *repGs) GetZ3JrDzqSzFja(id int64) ([]*object.Z3JrDzqSzFja, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3JrDzqSzFja, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var dzqId, dzqDpqHpId, dzqJyBz int

	rList := make([]*object.Z3JrDzqSzFja, 0)
	for rows.Next() {
		err := rows.Scan(&dzqId, &dzqDpqHpId, &dzqJyBz)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JrDzqSzFja", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3JrDzqSzFja{
			DzqId:      dzqId,
			DzqDpqHpId: dzqDpqHpId,
			DzqJyBz:    dzqJyBz,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JrDzqSzFja", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除节日电子券设置附加表索引
func (r *repGs) DelZ3JrDzqSzFjaSy(id int64) error {
	return goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3JrDzqSzFjaSy, id)
}

//获取机构表A索引
func (r *repGs) GetZ3FzJgaSy() ([]int64, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3FzJgaSy)
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
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JrDzqSzFjaSy", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JrDzqSzFjaSy", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//获取机构表A
func (r *repGs) GetZ3FzJga(id int64) ([]*object.Z3FzJga, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3FzJga, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var fzJgId, fzJgLx, fzJgIsForbidden int
	var fzJgsSxRq time.Time
	var fzJgSzMc, fzJgQc, fzJgPym, fzJgSym, fzJgTdm, fzJgXsm, fzJgLsZz string
	rList := make([]*object.Z3FzJga, 0)
	for rows.Next() {
		err := rows.Scan(&fzJgId, &fzJgSzMc, &fzJgQc, &fzJgLx, &fzJgsSxRq,
			&fzJgPym, &fzJgSym, &fzJgTdm, &fzJgXsm, &fzJgLsZz,
			&fzJgIsForbidden)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JrDzqSzFja", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3FzJga{
			FzJgId:          fzJgId,
			FzJgSzMc:        fzJgSzMc,
			FzJgQc:          fzJgQc,
			FzJgLx:          fzJgLx,
			FzJgsSxRq:       fzJgsSxRq,
			FzJgPym:         fzJgPym,
			FzJgSym:         fzJgSym,
			FzJgTdm:         fzJgTdm,
			FzJgXsm:         fzJgXsm,
			FzJgLsZz:        fzJgLsZz,
			FzJgIsForbidden: fzJgIsForbidden,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JrDzqSzFja", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除机构表A索引
func (r *repGs) DelZ3FzJgaSy(id int64) error {
	return goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3FzJgaSy, id)
}

//获取电子券设置索引
func (r *repGs) GetZ3DzqSzaSy() ([]int64, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3DzqSzaSy)
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
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3DzqSzaSy", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3DzqSzaSy", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//获取电子券设置
func (r *repGs) GetZ3DzqSza(id int64) ([]*object.Z3DzqSza, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3DzqSza, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var dzqId, dzqJrqBj, dzqYt, dzqHxhLy, dzqDbBs, dzqHsXz, dzqCtqId, dzqQxZkYh, dzqJyBz int
	var dzqSzMc, dzqDwXsMc, dzqBm, dzqPym string
	var dzqZzqMz, dzqZdZkl float64
	var dzqQhsJzr time.Time

	rList := make([]*object.Z3DzqSza, 0)
	for rows.Next() {
		err := rows.Scan(&dzqId, &dzqSzMc, &dzqDwXsMc, &dzqBm, &dzqPym,
			&dzqJrqBj, &dzqZzqMz, &dzqYt, &dzqHxhLy, &dzqDbBs,
			&dzqHsXz, &dzqCtqId, &dzqZdZkl, &dzqQxZkYh, &dzqQhsJzr,
			&dzqJyBz)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JrDzqSzFja", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3DzqSza{
			DzqId:     dzqId,
			DzqSzMc:   dzqSzMc,
			DzqDwXsMc: dzqDwXsMc,
			DzqBm:     dzqBm,
			DzqPym:    dzqPym,
			DzqJrqBj:  dzqJrqBj,
			DzqZzqMz:  dzqZzqMz,
			DzqYt:     dzqYt,
			DzqHxhLy:  dzqHxhLy,
			DzqDbBs:   dzqDbBs,
			DzqHsXz:   dzqHsXz,
			DzqCtqId:  dzqCtqId,
			DzqZdZkl:  dzqZdZkl,
			DzqQxZkYh: dzqQxZkYh,
			DzqQhsJzr: dzqQhsJzr,
			DzqJyBz:   dzqJyBz,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3JrDzqSzFja", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除电子券设置索引
func (r *repGs) DelZ3DzqSzaSy(id int64) error {
	return goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3DzqSzaSy, id)
}

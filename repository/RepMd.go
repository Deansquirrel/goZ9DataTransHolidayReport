package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goToolMSSql2000"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

const (
	sqlGetZ3XsCkDt = "" +
		"declare @lsh varchar(255) " +
		"select top 1 @lsh=cklsh from z3xscksyt order by cklsh " +
		"select ckdmxhh,ckdjh,ckdlsh,ckdbj,ckdcxbj, " +
		"	ckdyyr,ckdmdid,ckdckid,ckdhpid,ckddwid, " +
		"	ckdhsl,ckdjmsl,ckdbqjxj,ckdcjjxj,ckdfjxx, " +
		"	ckzdrid,ckkhid,ckbz,ckzdzsj " +
		"from ( " +
		"	select a.ckdmxhh as ckdmxhh,b.ckdjh as ckdjh,b.cklsh as ckdlsh,a.ckdbj as ckdbj,a.ckdcxbj as ckdcxbj, " +
		"		a.ckdyyr as ckdyyr,a.ckdmdid as ckdmdid,a.ckdckid as ckdckid,a.ckdhpid as ckdhpid,a.ckddwid as ckddwid, " +
		"		a.ckdhsl as ckdhsl,a.ckdjmsl as ckdjmsl,a.ckdbqjxj as ckdbqjxj,a.ckdcjjxj as ckdcjjxj,a.ckdfjxx as ckdfjxx, " +
		"		b.ckzdrid as ckzdrid,b.ckkhid as ckkhid,b.ckbz as ckbz,b.ckzdzsj as ckzdzsj " +
		"	from z3xsckdlst a " +
		"	inner join z3xscklst b on left(a.ckdmxhh,12) = b.cklsh " +
		"	where exists (select * from z3xscksyt b where b.cklsh = left(a.ckdmxhh,12)  and b.cklsh = @lsh) " +
		"	union all " +
		"	select a.ckdmxhh as ckdmxhh,b.ckdjh as ckdjh,b.cklsh as ckdlsh,a.ckdbj as ckdbj,a.ckdcxbj as ckdcxbj, " +
		"		a.ckdyyr as ckdyyr,a.ckdmdid as ckdmdid,a.ckdckid as ckdckid,a.ckdhpid as ckdhpid,a.ckddwid as ckddwid, " +
		"		a.ckdhsl as ckdhsl,a.ckdjmsl as ckdjmsl,a.ckdbqjxj as ckdbqjxj,a.ckdcjjxj as ckdcjjxj,a.ckdfjxx as ckdfjxx, " +
		"		b.ckzdrid as ckzdrid,b.ckkhid as ckkhid,b.ckbz as ckbz,b.ckzdzsj as ckzdzsj " +
		"	from z3xsckdt a " +
		"	inner join z3xsckt b on left(a.ckdmxhh,12) = b.cklsh " +
		"	where exists (select * from z3xscksyt b where b.cklsh = left(a.ckdmxhh,12) and b.cklsh = @lsh) " +
		") a " +
		"order by a.ckzdzsj asc"
	sqlDeleteZ3XsCkDtSy = "" +
		"delete from z3xscksyt " +
		"where cklsh = ?"
)

type repMd struct {
	dbConfig *goToolMSSql2000.MSSqlConfig
}

func NewRepMd() *repMd {
	comm := NewCommon()
	return &repMd{
		dbConfig: comm.ConvertDbConfigTo2000(comm.GetLocalDbConfig()),
	}
}

//销售出库货品明细
func (r *repMd) GetZ3XsCkDt() ([]*object.Z3XsCkDt, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3XsCkDt)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var ckdMxHh, ckDjh, ckdLsh, ckdFjXx, ckBz string
	var ckdBj, ckdCxBj, ckdMdId, ckdCkId, ckdHpId, ckdDwId, ckZdRid, ckKhId int
	var ckdYyr, ckZdzSj time.Time
	var ckdHsl, ckdJmSl, ckdBqjXj, ckdCjjXj float64
	rList := make([]*object.Z3XsCkDt, 0)
	for rows.Next() {
		err := rows.Scan(&ckdMxHh, &ckDjh, &ckdLsh, &ckdBj, &ckdCxBj,
			&ckdYyr, &ckdMdId, &ckdCkId, &ckdHpId, &ckdDwId,
			&ckdHsl, &ckdJmSl, &ckdBqjXj, &ckdCjjXj, &ckdFjXx,
			&ckZdRid, &ckKhId, &ckBz, &ckZdzSj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3XsCkDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3XsCkDt{
			CkdMxHh:  ckdMxHh,
			CkDjh:    ckDjh,
			CkdLsh:   ckdLsh,
			CkdBj:    ckdBj,
			CkdCxBj:  ckdCxBj,
			CkdYyr:   ckdYyr,
			CkdMdId:  ckdMdId,
			CkdCkId:  ckdCkId,
			CkdHpId:  ckdHpId,
			CkdDwId:  ckdDwId,
			CkdHsl:   ckdHsl,
			CkdJmSl:  ckdJmSl,
			CkdBqjXj: ckdBqjXj,
			CkdCjjXj: ckdCjjXj,
			CkdFjXx:  ckdFjXx,
			CkZdRid:  ckZdRid,
			CkKhId:   ckKhId,
			CkBz:     ckBz,
			CkZdzSj:  ckZdzSj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3XsCkDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//销售出库货品明细
func (r *repMd) DeleteZ3XsCkDtSy(id string) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDeleteZ3XsCkDtSy, id)
}

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
		"select top 1 @lsh=cklsh from z3xscksyt order by cklsh asc " +
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

	sqlGetZ3XsTht = "" +
		"declare @lsh varchar(255) " +
		"select top 1 @lsh=thlsh from z3xsthsyt order by thlsh asc " +
		"select a.thhpmxhh,b.thdjh,b.thlsh,a.thyyrq,a.thmdid, " +
		"	a.thqtckid,a.thhpid,a.thjldwid,a.thzhl,a.thjmsl, " +
		"	a.thbqjjexj,a.thzzcjjexj,a.thhpfjxx,b.thzdrid,b.thkhid," +
		"	b.thbz,b.thzdjzsj " +
		"from z3xsthhpdt a " +
		"inner join z3xstht b on left(a.thhpmxhh,12) = b.thlsh and b.thlsh = @lsh"
	sqlDelZ3XsThtSy = "" +
		"delete from z3xsthsyt " +
		"where thlsh = ?"

	sqlGetZ3MdDbCkDt = "" +
		"declare @lsh varchar(255) " +
		"select top 1 @lsh=dbdlsh from z3mddbcksyt order by dbdlsh asc " +
		"select a.dbdmxhh,b.dbdjh,b.dblsh,a.dbdckr,a.dbddcjgid, " +
		"	a.dbddcckid,b.dbrkshjgid,b.dbrkppid,a.dbdhpid,a.dbddwid, " +
		"	a.dbdhsl,a.dbdjmsl,a.dbdlsj,a.dbddhj,b.dbshrid, " +
		"	b.dbbz,b.dbshsj " +
		"from z3mddbckdt a " +
		"inner join z3mddbckt b on left(a.dbdmxhh,12) = b.dblsh and b.dblsh = @lsh"
	sqlDelZ3MdDbCkDtSy = "" +
		"delete from z3mddbcksyt " +
		"where dbdlsh = ?"
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

//获取销售出库货品明细
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

//删除销售出库货品明细索引
func (r *repMd) DelZ3XsCkDtSy(id string) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDeleteZ3XsCkDtSy, id)
}

//获取门店销售退货明细
func (r *repMd) GetZ3XsTht() ([]*object.Z3XsTht, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3XsTht)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var thhpmxhh, thdjh, thlsh, thhpfjxx, thbz string
	var thyyrq, thzdzsj time.Time
	var thmdid, thqtckid, thhpid, thjldwid, thzdrid, thkhid int
	var thzhl, thjmsl, thbqjjexj, thzzcjjexj float64
	rList := make([]*object.Z3XsTht, 0)
	for rows.Next() {
		err := rows.Scan(&thhpmxhh, &thdjh, &thlsh, &thyyrq, &thmdid,
			&thqtckid, &thhpid, &thjldwid, &thzhl, &thjmsl,
			&thbqjjexj, &thzzcjjexj, &thhpfjxx, &thzdrid, &thkhid,
			&thbz, &thzdzsj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3XsTht", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3XsTht{
			ThHpMxHh:   thhpmxhh,
			ThDjh:      thdjh,
			ThLsh:      thlsh,
			ThYyRq:     thyyrq,
			ThMdid:     thmdid,
			ThQtCkId:   thqtckid,
			ThHpId:     thhpid,
			ThJlDwId:   thjldwid,
			ThZhl:      thzhl,
			ThJmSl:     thjmsl,
			ThBqjJeXj:  thbqjjexj,
			ThZzCjJeXj: thzzcjjexj,
			ThHpFjXx:   thhpfjxx,
			ThZdrId:    thzdrid,
			ThKhId:     thkhid,
			ThBz:       thbz,
			ThZdzSj:    thzdzsj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3XsTht", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除门店销售退货明细索引
func (r *repMd) DelZ3XsThtSy(id string) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDelZ3XsThtSy, id)
}

//获取调拨出库明细
func (r *repMd) GetZ3MdDbCkDt() ([]*object.Z3MdDbCkDt, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3MdDbCkDt)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var dbdMxHh, dbdDjh, dbdLsh, dbdBz string
	var dbdCkr, dbdShSj time.Time
	var dbdDcJgId, dbdDcCkId, dbrKShJgId, dbrKPpId, dbdHpId, dbdDwId, dbdZdrId int
	var dbdHsl, dbdJmSl, dbdLsj, dbdDhj float64
	rList := make([]*object.Z3MdDbCkDt, 0)
	for rows.Next() {
		err := rows.Scan(&dbdMxHh, &dbdDjh, &dbdLsh, &dbdCkr, &dbdDcJgId,
			&dbdDcCkId, &dbrKShJgId, &dbrKPpId, &dbdHpId, &dbdDwId,
			&dbdHsl, &dbdJmSl, &dbdLsj, &dbdDhj, &dbdZdrId,
			&dbdBz, &dbdShSj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3MdDbCkDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3MdDbCkDt{
			DbdMxHh:    dbdMxHh,
			DbdDjh:     dbdDjh,
			DbdLsh:     dbdLsh,
			DbdCkr:     dbdCkr,
			DbdDcJgId:  dbdDcJgId,
			DbdDcCkId:  dbdDcCkId,
			DbrKShJgId: dbrKShJgId,
			DbrKPpId:   dbrKPpId,
			DbdHpId:    dbdHpId,
			DbdDwId:    dbdDwId,
			DbdHsl:     dbdHsl,
			DbdJmSl:    dbdJmSl,
			DbdLsj:     dbdLsj,
			DbdDhj:     dbdDhj,
			DbdZdrId:   dbdZdrId,
			DbdBz:      dbdBz,
			DbdShSj:    dbdShSj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3MdDbCkDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除调拨出库索引
func (r *repMd) DelZ3MdDbCkDtSy(id string) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDelZ3MdDbCkDtSy, id)
}

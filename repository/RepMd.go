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

	sqlGetZ3MdDbTzDt = "" +
		"declare @lsh varchar(255) " +
		"select top 1 @lsh=tzdlsh from z3dbtzsyt order by tzdlsh asc " +
		"select a.tzdmxhh,b.tzdjh as tzddjh,b.tzlsh as tzdlsh,a.tzdckr,a.tzdjgid," +
		"	a.tzdckid,b.tzrkjgid,b.tzrkckid,a.tzdhpid,a.tzddwid," +
		"	a.tzdhsl,a.tzdjmtzsl,a.tzdlsj,a.tzddhj,b.tzqrrid as tzdzdrid," +
		"	b.tzbz as tzdbz,b.tzqrsj as tzdshsj " +
		"from z3dbtzdt a " +
		"inner join z3dbtzt b on left(a.tzdmxhh,12) = b.tzlsh and b.tzlsh = @lsh"
	sqlDelZ3MdDbTzDtSy = "" +
		"delete from z3dbtzsyt " +
		"where tzdlsh = ?"

	sqlGetZ3HpCkDjDt = "" +
		"declare @lsh varchar(255) " +
		"select top 1 @lsh=cklsh from z3hpckdjsyt order by cklsh asc " +
		"select a.ckdmxhh,b.ckdjh as ckddjh,b.cklsh as ckdlsh,a.ckdyyr,a.ckdckfzjgid, " +
		"	a.ckdckid,a.ckdhpid,a.ckddwid,a.ckdhsl,a.ckdjmsl, " +
		"	a.ckdlsj,a.ckddhj,b.ckbz,b.ckshrid,b.ckshsj " +
		"from z3hpckdjdt a " +
		"inner join z3hpckdjt b on left(a.ckdmxhh,12) = b.cklsh and b.cklsh = @lsh"
	sqlDelZ3HpCkDjDtSy = "" +
		"delete from z3hpckdjsyt " +
		"where cklsh = ?"

	sqlGetZ3HpRkDjDt = "" +
		"declare @lsh varchar(255) " +
		"select top 1 @lsh=rklsh from z3hprkdjsyt order by rklsh asc " +
		"select a.rkdmxhh,b.rkdjh as rkddjh,b.rklsh as rkdlsh,a.rkdyyr,a.rkdrkfzjgid, " +
		"	a.rkdckid,a.rkdhpid,a.rkddwid,a.rkdhsl,a.rkdjmsl, " +
		"	a.rkddhj,a.rkdlsj,b.rkbz,b.rkshrid,b.rkshsj " +
		"from z3hprkdjdt a inner join z3hprkdjt b on left(a.rkdmxhh,12) = b.rklsh and b.rklsh = @lsh"
	sqlDelZ3HpRkDjDtSy = "" +
		"delete from z3hprkdjsyt " +
		"where rklsh = ?"

	sqlGetZ3PkDjDt = "" +
		"declare @lsh varchar(255) " +
		"select top 1 @lsh=pkdlsh from z3pkdjsyt order by pkdlsh asc " +
		"select a.pkdmxhh,b.pkdjh as pkddjh,b.pklsh as pkdlsh,a.pkdyyr,a.pkdpdjgid, " +
		"	a.pkdpdckid,a.pkdhpid,a.pkddwid,a.pkdhsl,a.pkdjmsl, " +
		"	a.pkdlsj,a.pkddhj,b.pkbz as pkdbz,b.pkshrid as pkdshrid,b.pkshsj as pkdshsj " +
		"from z3pkdjdt a inner join z3pkdjt b on left(a.pkdmxhh,12) = b.pklsh and b.pklsh = @lsh"
	sqlDelZ3PkDjDt = "" +
		"delete from z3pkdjsyt " +
		"where pkdlsh = ?"
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

//获取调拨调整明细
func (r *repMd) GetZ3MdDbTzDt() ([]*object.Z3MdDbTzDt, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3MdDbTzDt)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var tzdMxHh, tzdDjh, tzdLsh, tzdBz string
	var tzdCkr, tzdShSj time.Time
	var tzdJgId, tzdCkId, tzRkJgId, tzRkCkId, tzdHpId, tzdDwId, tzdZdrId int
	var tzdHsl, tzdJmTzSl, tzdLsj, tzdDhj float64
	rList := make([]*object.Z3MdDbTzDt, 0)
	for rows.Next() {
		err := rows.Scan(&tzdMxHh, &tzdDjh, &tzdLsh, &tzdCkr, &tzdJgId,
			&tzdCkId, &tzRkJgId, &tzRkCkId, &tzdHpId, &tzdDwId,
			&tzdHsl, &tzdJmTzSl, &tzdLsj, &tzdDhj, &tzdZdrId,
			&tzdBz, &tzdShSj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3MdDbTzDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3MdDbTzDt{
			TzdMxHh:   tzdMxHh,
			TzdDjh:    tzdDjh,
			TzdLsh:    tzdLsh,
			TzdCkr:    tzdCkr,
			TzdJgId:   tzdJgId,
			TzdCkId:   tzdCkId,
			TzRkJgId:  tzRkJgId,
			TzRkCkId:  tzRkCkId,
			TzdHpId:   tzdHpId,
			TzdDwId:   tzdDwId,
			TzdHsl:    tzdHsl,
			TzdJmTzSl: tzdJmTzSl,
			TzdLsj:    tzdLsj,
			TzdDhj:    tzdDhj,
			TzdZdrId:  tzdZdrId,
			TzdBz:     tzdBz,
			TzdShSj:   tzdShSj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3MdDbTzDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除调拨调整明细
func (r *repMd) DelZ3MdDbTzDtSy(id string) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDelZ3MdDbTzDtSy, id)
}

//获取货品出库登记
func (r *repMd) GetZ3HpCkDjDt() ([]*object.Z3HpCkDjDt, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpCkDjDt)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var ckdMxHh, ckdDjh, ckdLsh, ckBz string
	var ckdYyr, ckShSj time.Time
	var ckdCkFzJgId, ckdCkId, ckdHpId, ckdDwId, ckShrId int
	var ckdHsl, ckdJmSl, ckdLsj, ckdDhj float64
	rList := make([]*object.Z3HpCkDjDt, 0)
	for rows.Next() {
		err := rows.Scan(&ckdMxHh, &ckdDjh, &ckdLsh, &ckdYyr, &ckdCkFzJgId,
			&ckdCkId, &ckdHpId, &ckdDwId, &ckdHsl, &ckdJmSl,
			&ckdLsj, &ckdDhj, &ckBz, &ckShrId, &ckShSj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpCkDjDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3HpCkDjDt{
			CkdMxHh:     ckdMxHh,
			CkdDjh:      ckdDjh,
			CkdLsh:      ckdLsh,
			CkdYyr:      ckdYyr,
			CkdCkFzJgId: ckdCkFzJgId,
			CkdCkId:     ckdCkId,
			CkdHpId:     ckdHpId,
			CkdDwId:     ckdDwId,
			CkdHsl:      ckdHsl,
			CkdJmSl:     ckdJmSl,
			CkdLsj:      ckdLsj,
			CkdDhj:      ckdDhj,
			CkBz:        ckBz,
			CkShrId:     ckShrId,
			CkShSj:      ckShSj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpCkDjDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除货品出库登记
func (r *repMd) DelZ3HpCkDjDt(id string) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDelZ3HpCkDjDtSy, id)
}

//获取货品入库登记
func (r *repMd) GetZ3HpRkDjDt() ([]*object.Z3HpRkDjDt, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3HpRkDjDt)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var rkdMxHh, rkdDjh, rkdLsh, rkBz string
	var rkdYyr, rksShSj time.Time
	var rkdRkFzJgId, rkdCkId, rkdHpId, rkdDwId, rkShrId int
	var rkdHsl, rkdJmSl, rkdDhj, rkdLsj float64
	rList := make([]*object.Z3HpRkDjDt, 0)
	for rows.Next() {
		err := rows.Scan(&rkdMxHh, &rkdDjh, &rkdLsh, &rkdYyr, &rkdRkFzJgId,
			&rkdCkId, &rkdHpId, &rkdDwId, &rkdHsl, &rkdJmSl,
			&rkdDhj, &rkdLsj, &rkBz, &rkShrId, &rksShSj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpRkDjDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3HpRkDjDt{
			RkdMxHh:     rkdMxHh,
			RkdDjh:      rkdDjh,
			RkdLsh:      rkdLsh,
			RkdYyr:      rkdYyr,
			RkdRkFzJgId: rkdRkFzJgId,
			RkdCkId:     rkdCkId,
			RkdHpId:     rkdHpId,
			RkdDwId:     rkdDwId,
			RkdHsl:      rkdHsl,
			RkdJmSl:     rkdJmSl,
			RkdDhj:      rkdDhj,
			RkdLsj:      rkdLsj,
			RkBz:        rkBz,
			RkShrId:     rkShrId,
			RksShSj:     rksShSj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3HpRkDjDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除货品入库登记
func (r *repMd) DelZ3HpRkDjDtSy(id string) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDelZ3HpRkDjDtSy, id)
}

//获取盘亏登记
func (r *repMd) GetZ3PkDjDt() ([]*object.Z3PkDjDt, error) {
	comm := NewCommon()
	rows, err := comm.GetRowsBySQL2000(r.dbConfig, sqlGetZ3PkDjDt)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var pkdMxHh, pkdDjh, pkdLsh, pkdBz string
	var pkdYyr, pkdShSj time.Time
	var pkdPdJgId, pkdPdCkId, pkdHpId, pkdDwId, pkdShrId int
	var pkdHsl, pkdJmSl, pkdLsj, pkdDhj float64
	rList := make([]*object.Z3PkDjDt, 0)
	for rows.Next() {
		err := rows.Scan(&pkdMxHh, &pkdDjh, &pkdLsh, &pkdYyr, &pkdPdJgId,
			&pkdPdCkId, &pkdHpId, &pkdDwId, &pkdHsl, &pkdJmSl,
			&pkdLsj, &pkdDhj, &pkdBz, &pkdShrId, &pkdShSj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PkDjDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3PkDjDt{
			PkdMxHh:   pkdMxHh,
			PkdDjh:    pkdDjh,
			PkdLsh:    pkdLsh,
			PkdYyr:    pkdYyr,
			PkdPdJgId: pkdPdJgId,
			PkdPdCkId: pkdPdCkId,
			PkdHpId:   pkdHpId,
			PkdDwId:   pkdDwId,
			PkdHsl:    pkdHsl,
			PkdJmSl:   pkdJmSl,
			PkdLsj:    pkdLsj,
			PkdDhj:    pkdDhj,
			PkdBz:     pkdBz,
			PkdShrId:  pkdShrId,
			PkdShSj:   pkdShSj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PkDjDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//删除盘亏登记
func (r *repMd) DelZ3PkDjDtSy(id string) error {
	comm := NewCommon()
	return comm.SetRowsBySQL2000(r.dbConfig, sqlDelZ3PkDjDt, id)
}

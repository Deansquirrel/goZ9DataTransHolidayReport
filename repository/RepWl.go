package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolMSSql2000"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

const (
	sqlGetZ3PsCkHpDtSy = "" +
		"select top 1 ckdlsh " +
		"from z3mdpscksyt"
	sqlGetZ3PsCkHpDt = "" +
		"select a.[ckdlsh],a.[ckdhpid],c.[ckdjh] as [ckddjh],a.[ckddwid],a.[ckdhsl], " +
		"	a.[ckdcckfzjgid],a.[ckdcckid],c.[ckmdid],c.[ckppid],a.[ckdeddjlx], " +
		"	a.[ckddhzxsl],a.[ckdzxdhj],a.[ckddhje],a.[ckdzspsj],c.ckhrq as [ckdhrq], " +
		"	a.[ckdshsj],a.[ckdpsjmsl],d.[ckpsbz],d.[ckfcczrid] " +
		"from [z3mdpsckdt] a " +
		"inner join [z3mdpsckt] c on c.[cklsh] = a.[ckdlsh] " +
		"inner join [z3mdpsckfjt] d on a.[ckdlsh] = d.[cklsh] " +
		"where a.[ckdlsh]=?"
	sqlGetZ3PsCkHpFjDt = "" +
		"select [ckdlsh],[ckdhpid],[ckghdmxhh],[ckddwid],[ckdhsl], " +
		"	[ckzxddsl],[ckzxpssl],[ckphbj] " +
		"from [z3psckhpfjdt] " +
		"where [ckdlsh]=?"
	sqlDelZ3PsCkHpDtSy = "" +
		"DELETE FROM [z3mdpscksyt] " +
		"WHERE [ckdlsh]=?"

	sqlGetZ3PsXzDtSy = "" +
		"select top 1 xzdlsh " +
		"from z3psxzsyt"
	sqlGetZ3PsXzDt = "" +
		"select [xzdmxhh],[xzdjh],[xzlsh],[xzdxzr],[xzdshjgid], " +
		"	[xzdshckid],[xzrkjgid],[xzrkckid],[xzdhpid],[xzddwid], " +
		"	[xzdhsl],[xzdjmtzsl],[xzdzxtzsl],[xzdjmxzsl],[xzddhj], " +
		"	[xzdpsj],[xzbz],[xzrid],[xzsj] " +
		"from z3psxzdt a " +
		"inner join z3psxzt b on a.xzdmxhh between b.xzlsh+'000' and  b.xzlsh+'zzz' " +
		"where b.xzlsh = ?"
	sqlDelZ3PsXzDtSy = "" +
		"DELETE FROM [z3psxzsyt] " +
		"WHERE [xzdlsh]=?"

	sqlGetZ3MdPsChDtSy = "" +
		"select top 1 chdlsh " +
		"from z3pschsyt"
	sqlGetZ3MdPsChDt = "" +
		"select [chdlsh],[chdjh],[chlsh],[chdhpid],[chdcckfzjgid], " +
		"	[chdcckid],[chdhsr],[chdhsxzid],[chdhsbmid],[chddwid], " +
		"	[chdhsl],[chdpsjmsl],[chdpsj],[chbz],[chshrid], " +
		"	[chshsj] " +
		"from z3mdpschdt b " +
		"inner join z3mdpscht c on b.chdlsh = c.chlsh " +
		"where b.chdlsh = ?"
	sqlDelZ3MdPsChDtSy = "" +
		"DELETE FROM [z3pschsyt] " +
		"WHERE [chdlsh]=?"

	sqlGetShengChWgRkDtSy = "" +
		"SELECT TOP 1 [wgrklsh] " +
		"FROM [shengchwgrksyt]"
	sqlGetShengChWgRkDt = "" +
		"select [wgrkmxhh],[wgrkdjh],[wgrklsh],b.[wgrkhsrq],[wgrkhpid], " +
		"	[wgrkdwid],[wgrkhsl],[wgrkrfzjgid],[wgrkrckid],[wgrkjmsl], " +
		"	[wgrkcpczdj],[wgrkjg],c.[wgrkjdrid] as [wgrkzdrid],[wgrkshsj],[wgrkbz] " +
		"from shengchwgrkdt b " +
		"inner join shengchwgrkt c on  b.wgrkmxhh between c.wgrklsh+'000' and c.wgrklsh+'zzz' " +
		"where c.wgrklsh = ?"
	sqlDelShengChWgRkDtSy = "" +
		"DELETE FROM [shengchwgrksyt] " +
		"WHERE [wgrklsh]=?"

	sqlGetZ3SheZhXsHpMxtSy = "" +
		"SELECT TOP 1 [cklsh] " +
		"FROM [z3shezhxssyt]"
	sqlGetZ3SheZhXsHpMxt = "" +
		"select [ckmxhh],[ckdjh],[cklsh],b.[ckhsrq],b.[ckcckfzjgid], " +
		"	b.[ckcckid],[ckhpid],[ckdwid],[ckhsl],[ckjmdj], " +
		"	[ckjmsl],[ckcjjexj],[ckjmlsj],[cklsjexj],c.[ckjdrid] as [ckzdrid], " +
		"	[ckshsj],[ckbz] " +
		"from z3shezhxshpmxt b " +
		"inner join z3shezhxsckt c on b.ckmxhh  between c.cklsh+'000' and c.cklsh+'zzz' " +
		"where c.cklsh = ?"
	sqlDelZ3SheZhXsHpMxtSy = "" +
		"DELETE FROM [z3shezhxssyt] " +
		"WHERE [cklsh]=?"

	sqlGetXtTz = "" +
		"SELECT [dptid],[brid],[gsid],[gsqty],[updatedate] " +
		"FROM [xttz] " +
		"WHERE [updatedate] > ?"

	sqlGetOoDxV1DdShCktSy = "" +
		"SELECT TOP 1 [shmxhh] " +
		"FROM [oodxv1ddshcksyt]"
	sqlGetOoDxV1DdShCkt = "" +
		"select shlsh,shmxhh,shhsr,shfzjgid,shckid, " +
		"	shbmid,shxzid,shsendrid,shkhid,shdddjh, " +
		"	shddlsh,shddmxhh,shphbj,shhpid,shjldwid, " +
		"	shhsl,shddzxsl,shjhjmsl,ShShJmSl,shshrid, " +
		"	shshsj " +
		"from oodxv1ddshckt	" +
		"where shmxhh = ?"
	sqlDelOoDxV1DdShCktSy = "" +
		"DELETE FROM [oodxv1ddshcksyt] " +
		"WHERE [shmxhh]=?"

	sqlGetZ3MdDhDdDtSy = "" +
		"SELECT TOP 1 [dhdlsh] " +
		"FROM [z3mddhsyt]"
	sqlGetZ3MdDhDdDt = "" +
		"select dhdlsh,dhdhpid,dhdjh,dhxsrq,dhdhwid," +
		"	dhdhsl,dhdjljmsjdhs,dhddhj,dhdxdrid,dhdxdsj " +
		"from z3mddhdt b " +
		"inner join z3mddht c on b.dhdlsh = c.dhlsh " +
		"where b.dhdlsh = ?"
	sqlDelZ3MdDhDdDtSy = "" +
		"DELETE FROM [z3mddhsyt] " +
		"WHERE [dhdlsh]=?"
)

type repWl struct {
	dbConfig *goToolMSSql2000.MSSqlConfig
}

func NewRepWl() *repWl {
	comm := NewCommon()
	return &repWl{
		dbConfig: goToolMSSqlHelper.ConvertDbConfigTo2000(comm.GetLocalDbConfig()),
	}
}

//配送出库&配送出库二级明细表
func (r *repWl) GetZ3PsCkHpDtSy() ([]string, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3PsCkHpDtSy)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]string, 0)
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PsCkHpDtSy", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PsCkHpDtSy", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}
func (r *repWl) GetZ3PsCkHpDt(id string) ([]*object.Z3PsCkHpDt, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3PsCkHpDt, id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var ckdLsh, ckdDjh, ckPsBz string
	var ckdHpId, ckdDwId, ckdCckFzJgId, ckdCckId, ckMdId, ckPpId, ckdEdDjLx, ckFcCzrId int
	var ckdHsl, ckdDhZxSl, ckdZxDhj, ckdDhJe, ckdZsPsj, ckdPsJmSl float64
	var ckdHrq, ckdShSj time.Time
	rList := make([]*object.Z3PsCkHpDt, 0)
	for rows.Next() {
		err := rows.Scan(&ckdLsh, &ckdHpId, &ckdDjh, &ckdDwId, &ckdHsl,
			&ckdCckFzJgId, &ckdCckId, &ckMdId, &ckPpId, &ckdEdDjLx,
			&ckdDhZxSl, &ckdZxDhj, &ckdDhJe, &ckdZsPsj, &ckdHrq,
			&ckdShSj, &ckdPsJmSl, &ckPsBz, &ckFcCzrId)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PsCkHpDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3PsCkHpDt{
			CkdLsh:       ckdLsh,
			CkdHpId:      ckdHpId,
			CkdDjh:       ckdDjh,
			CkdDwId:      ckdDwId,
			CkdHsl:       ckdHsl,
			CkdCckFzJgId: ckdCckFzJgId,
			CkdCckId:     ckdCckId,
			CkMdId:       ckMdId,
			CkPpId:       ckPpId,
			CkdEdDjLx:    ckdEdDjLx,
			CkdDhZxSl:    ckdDhZxSl,
			CkdZxDhj:     ckdZxDhj,
			CkdDhJe:      ckdDhJe,
			CkdZsPsj:     ckdZsPsj,
			CkdHrq:       ckdHrq,
			CkdShSj:      ckdShSj,
			CkdPsJmSl:    ckdPsJmSl,
			CkPsBz:       ckPsBz,
			CkFcCzrId:    ckFcCzrId,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PsCkHpDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}
func (r *repWl) GetZ3PsCkHpFjDt(id string) ([]*object.Z3PsCkHpFjDt, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3PsCkHpFjDt, id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var ckdLsh, ckGhdMxHh string
	var ckdHpId, ckdDwId, ckPhBj int
	var ckdHsl, ckZxDdSl, ckZxPsSl float64
	rList := make([]*object.Z3PsCkHpFjDt, 0)
	for rows.Next() {
		err := rows.Scan(&ckdLsh, &ckdHpId, &ckGhdMxHh, &ckdDwId, &ckdHsl,
			&ckZxDdSl, &ckZxPsSl, &ckPhBj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PsCkHpFjDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3PsCkHpFjDt{
			CkdLsh:    ckdLsh,
			CkdHpId:   ckdHpId,
			CkGhdMxHh: ckGhdMxHh,
			CkdDwId:   ckdDwId,
			CkdHsl:    ckdHsl,
			CkZxDdSl:  ckZxDdSl,
			CkZxPsSl:  ckZxPsSl,
			CkPhBj:    ckPhBj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PsCkHpFjDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}
func (r *repWl) DelZ3PsCkHpDtSy(id string) error {
	err := goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3PsCkHpDtSy, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//配送修正
func (r *repWl) GetZ3PsXzDtSy() ([]string, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3PsXzDtSy)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]string, 0)
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PsXzDtSy", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PsXzDtSy", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}
func (r *repWl) GetZ3PsXzDt(id string) ([]*object.Z3PsXzDt, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3PsXzDt, id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var xzdMxHh, xzDjh, xzLsh, xzBz string
	var xzdXzr, xzSj time.Time
	var xzdShJgId, xzdShCkId, xzRkJgId, xzRkCkId, xzdHpId, xzdDwId, xzrId int
	var xzdHsl, xzdJmTzSl, xzdZxTzSl, xzdJmXzSl, xzdDhj, xzdPsj float64
	rList := make([]*object.Z3PsXzDt, 0)
	for rows.Next() {
		err := rows.Scan(&xzdMxHh, &xzDjh, &xzLsh, &xzdXzr, &xzdShJgId,
			&xzdShCkId, &xzRkJgId, &xzRkCkId, &xzdHpId, &xzdDwId,
			&xzdHsl, &xzdJmTzSl, &xzdZxTzSl, &xzdJmXzSl, &xzdDhj,
			&xzdPsj, &xzBz, &xzrId, &xzSj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PsXzDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3PsXzDt{
			XzdMxHh:   xzdMxHh,
			XzDjh:     xzDjh,
			XzLsh:     xzLsh,
			XzdXzr:    xzdXzr,
			XzdShJgId: xzdShJgId,
			XzdShCkId: xzdShCkId,
			XzRkJgId:  xzRkJgId,
			XzRkCkId:  xzRkCkId,
			XzdHpId:   xzdHpId,
			XzdDwId:   xzdDwId,
			XzdHsl:    xzdHsl,
			XzdJmTzSl: xzdJmTzSl,
			XzdZxTzSl: xzdZxTzSl,
			XzdJmXzSl: xzdJmXzSl,
			XzdDhj:    xzdDhj,
			XzdPsj:    xzdPsj,
			XzBz:      xzBz,
			XzrId:     xzrId,
			XzSj:      xzSj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3PsXzDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}
func (r *repWl) DelZ3PsXzDtSy(id string) error {
	err := goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3PsXzDtSy, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//配送冲红单
func (r *repWl) GetZ3MdPsChDtSy() ([]string, error) {
	return r.getSyStr("GetZ3MdPsChDtSy", sqlGetZ3MdPsChDtSy)
}
func (r *repWl) GetZ3MdPsChDt(id string) ([]*object.Z3MdPsChDt, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3MdPsChDt, id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var chdLsh, chDjh, chLsh, chBz string
	var chdHpId, chdCckFzJgId, chdCckId, chdHsXzId, chdHsBmId, chdDwId, chShrId int
	var chdHsr, chShSj time.Time
	var chdHsl, chdPsJmSl, chdPsj float64
	rList := make([]*object.Z3MdPsChDt, 0)
	for rows.Next() {
		err := rows.Scan(&chdLsh, &chDjh, &chLsh, &chdHpId, &chdCckFzJgId,
			&chdCckId, &chdHsr, &chdHsXzId, &chdHsBmId, &chdDwId,
			&chdHsl, &chdPsJmSl, &chdPsj, &chBz, &chShrId,
			&chShSj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3MdPsChDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3MdPsChDt{
			ChdLsh:       chdLsh,
			ChDjh:        chDjh,
			ChLsh:        chLsh,
			ChdHpId:      chdHpId,
			ChdCckFzJgId: chdCckFzJgId,
			ChdCckId:     chdCckId,
			ChdHsr:       chdHsr,
			ChdHsXzId:    chdHsXzId,
			ChdHsBmId:    chdHsBmId,
			ChdDwId:      chdDwId,
			ChdHsl:       chdHsl,
			ChdPsJmSl:    chdPsJmSl,
			ChdPsj:       chdPsj,
			ChBz:         chBz,
			ChShrId:      chShrId,
			ChShSj:       chShSj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3MdPsChDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}
func (r *repWl) DelZ3MdPsChDtSy(id string) error {
	err := goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3MdPsChDtSy, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//完工入库
func (r *repWl) GetShengChWgRkDtSy() ([]string, error) {
	return r.getSyStr("GetShengChWgRkDtSy", sqlGetShengChWgRkDtSy)
}
func (r *repWl) GetShengChWgRkDt(id string) ([]*object.ShengChWgRkDt, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetShengChWgRkDt, id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var wgRkMxHh, wgRkDjh, wgRkLsh, wgRkBz string
	var wgRkHsRq, wgRkShSj time.Time
	var wgRkHpId, wgRkDwId, wgRkrFzJgId, wgRkrCkId, wgRkZdrId int
	var wgRkHsl, wgRkJmSl, wgRkCpCzDj, wgRkJg float64
	rList := make([]*object.ShengChWgRkDt, 0)
	for rows.Next() {
		err := rows.Scan(&wgRkMxHh, &wgRkDjh, &wgRkLsh, &wgRkHsRq, &wgRkHpId,
			&wgRkDwId, &wgRkHsl, &wgRkrFzJgId, &wgRkrCkId, &wgRkJmSl,
			&wgRkCpCzDj, &wgRkJg, &wgRkZdrId, &wgRkShSj, &wgRkBz)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetShengChWgRkDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.ShengChWgRkDt{
			WgRkMxHh:    wgRkMxHh,
			WgRkDjh:     wgRkDjh,
			WgRkLsh:     wgRkLsh,
			WgRkHsRq:    wgRkHsRq,
			WgRkHpId:    wgRkHpId,
			WgRkDwId:    wgRkDwId,
			WgRkHsl:     wgRkHsl,
			WgRkrFzJgId: wgRkrFzJgId,
			WgRkrCkId:   wgRkrCkId,
			WgRkJmSl:    wgRkJmSl,
			WgRkCpCzDj:  wgRkCpCzDj,
			WgRkJg:      wgRkJg,
			WgRkZdrId:   wgRkZdrId,
			WgRkShSj:    wgRkShSj,
			WgRkBz:      wgRkBz,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetShengChWgRkDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}
func (r *repWl) DelShengChWgRkDtSy(id string) error {
	err := goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelShengChWgRkDtSy, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//赊账销售出库
func (r *repWl) GetZ3SheZhXsHpMxtSy() ([]string, error) {
	return r.getSyStr("GetZ3SheZhXsHpMxtSy", sqlGetZ3SheZhXsHpMxtSy)
}
func (r *repWl) GetZ3SheZhXsHpMxt(id string) ([]*object.Z3SheZhXsHpMxt, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3SheZhXsHpMxt, id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var ckMxHh, ckDjh, ckLsh, ckBz string
	var ckHsRq, ckShSj time.Time
	var ckCckFzJgId, ckCckId, ckHpId, ckDwId, ckZdrId int
	var ckHsl, ckJmDj, ckJmSl, ckCjJeXj, ckJmLsj, ckLsJeXj float64
	rList := make([]*object.Z3SheZhXsHpMxt, 0)
	for rows.Next() {
		err := rows.Scan(&ckMxHh, &ckDjh, &ckLsh, &ckHsRq, &ckCckFzJgId,
			&ckCckId, &ckHpId, &ckDwId, &ckHsl, &ckJmDj,
			&ckJmSl, &ckCjJeXj, &ckJmLsj, &ckLsJeXj, &ckZdrId,
			&ckShSj, &ckBz)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3SheZhXsHpMxt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3SheZhXsHpMxt{
			CkMxHh:      ckMxHh,
			CkDjh:       ckDjh,
			CkLsh:       ckLsh,
			CkHsRq:      ckHsRq,
			CkCckFzJgId: ckCckFzJgId,
			CkCckId:     ckCckId,
			CkHpId:      ckHpId,
			CkDwId:      ckDwId,
			CkHsl:       ckHsl,
			CkJmDj:      ckJmDj,
			CkJmSl:      ckJmSl,
			CkCjJeXj:    ckCjJeXj,
			CkJmLsj:     ckJmLsj,
			CkLsJeXj:    ckLsJeXj,
			CkZdrId:     ckZdrId,
			CkShSj:      ckShSj,
			CkBz:        ckBz,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3SheZhXsHpMxt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}
func (r *repWl) DelZ3SheZhXsHpMxtSy(id string) error {
	err := goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3SheZhXsHpMxtSy, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//工厂即时台帐表
func (r *repWl) GetXtTz(lastUpdate time.Time) ([]*object.XtTz, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetXtTz,
		goToolCommon.GetDateTimeStrWithMillisecond(lastUpdate))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var dptId, brId, gsId int
	var gsQty float64
	var updateDate time.Time
	rList := make([]*object.XtTz, 0)
	for rows.Next() {
		err := rows.Scan(&dptId, &brId, &gsId, &gsQty, &updateDate)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetXtTz", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.XtTz{
			DptId:      dptId,
			BrId:       brId,
			GsId:       gsId,
			GsQty:      gsQty,
			UpdateDate: updateDate,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetXtTz", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//订单提货单
func (r *repWl) GetOoDxV1DdShCktSy() ([]string, error) {
	return r.getSyStr("GetOoDxV1DdShCktSy", sqlGetOoDxV1DdShCktSy)
}
func (r *repWl) GetOoDxV1DdShCkt(id string) ([]*object.OoDxV1DdShCkt, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetOoDxV1DdShCkt, id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var shLsh, shMxHh, shDdDjh, shDdLsh, shDdMxHh string
	var shHsr, shShSj time.Time
	var shFzJgId, shCkId, shBmId, shXzId, shSendRid, shKhId, shPhBj, shHpId, shJlDwId, shShrId int
	var shHsl, shDdZxSl, shJhJmSl, shShJmSl float64

	rList := make([]*object.OoDxV1DdShCkt, 0)
	for rows.Next() {
		err := rows.Scan(&shLsh, &shMxHh, &shHsr, &shFzJgId, &shCkId,
			&shBmId, &shXzId, &shSendRid, &shKhId, &shDdDjh,
			&shDdLsh, &shDdMxHh, &shPhBj, &shHpId, &shJlDwId,
			&shHsl, &shDdZxSl, &shJhJmSl, &shShJmSl, &shShrId,
			&shShSj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetOoDxV1DdShCkt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.OoDxV1DdShCkt{
			ShLsh:     shLsh,
			ShMxHh:    shMxHh,
			ShHsr:     shHsr,
			ShFzJgId:  shFzJgId,
			ShCkId:    shCkId,
			ShBmId:    shBmId,
			ShXzId:    shXzId,
			ShSendRid: shSendRid,
			ShKhId:    shKhId,
			ShDdDjh:   shDdDjh,
			ShDdLsh:   shDdLsh,
			ShDdMxHh:  shDdMxHh,
			ShPhBj:    shPhBj,
			ShHpId:    shHpId,
			ShJlDwId:  shJlDwId,
			ShHsl:     shHsl,
			ShDdZxSl:  shDdZxSl,
			ShJhJmSl:  shJhJmSl,
			ShShJmSl:  shShJmSl,
			ShShrId:   shShrId,
			ShShSj:    shShSj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetOoDxV1DdShCkt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}
func (r *repWl) DelOoDxV1DdShCktSy(id string) error {
	err := goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelOoDxV1DdShCktSy, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//门店订货代单
func (r *repWl) GetZ3MdDhDdDtSy() ([]string, error) {
	return r.getSyStr("GetZ3MdDhDdDtSy", sqlGetZ3MdDhDdDtSy)
}
func (r *repWl) GetZ3MdDhDdDt(id string) ([]*object.Z3MdDhDdDt, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlGetZ3MdDhDdDt, id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	var dhdLsh, dhDjh string
	var dhdHpId, dhdHwId, dhdXdrId int
	var dhXsRq, dhdXdSj time.Time
	var dhdHsl, dhdJlJmSjDhs, dhdDhj float64
	rList := make([]*object.Z3MdDhDdDt, 0)
	for rows.Next() {
		err := rows.Scan(&dhdLsh, &dhdHpId, &dhDjh, &dhXsRq, &dhdHwId,
			&dhdHsl, &dhdJlJmSjDhs, &dhdDhj, &dhdXdrId, &dhdXdSj)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3MdDhDdDt", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.Z3MdDhDdDt{
			DhdLsh:       dhdLsh,
			DhdHpId:      dhdHpId,
			DhDjh:        dhDjh,
			DhXsRq:       dhXsRq,
			DhdHwId:      dhdHwId,
			DhdHsl:       dhdHsl,
			DhdJlJmSjDhs: dhdJlJmSjDhs,
			DhdDhj:       dhdDhj,
			DhdXdrId:     dhdXdrId,
			DhdXdSj:      dhdXdSj,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetZ3MdDhDdDt", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}
func (r *repWl) DelZ3MdDhDdDtSy(id string) error {
	err := goToolMSSqlHelper.SetRowsBySQL2000(r.dbConfig, sqlDelZ3MdDhDdDtSy, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (r *repWl) getSyStr(key string, sqlStr string) ([]string, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL2000(r.dbConfig, sqlStr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]string, 0)
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", key, err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", key, rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

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

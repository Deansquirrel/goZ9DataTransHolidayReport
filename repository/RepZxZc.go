package repository

import (
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
)

const (
	sqlUpdateMdZ3XsCkDt = "" +
		"IF EXISTS (SELECT * FROM [z3xsckdt] WHERE [ckdmxhh]=?) " +
		"	BEGIN " +
		"		UPDATE [z3xsckdt] " +
		"		SET [ckdmxhh] = ?,[ckdjh] = ?,[ckdlsh] = ?,[ckdbj] = ?,[ckdcxbj] = ?, " +
		"			[ckdyyr] = ?,[ckdmdid] = ?,[ckdhpid] = ?,[ckddwid] = ?,[ckdhsl] = ?, " +
		"			[ckdjmsl] = ?,[ckdfjxx] = ?,[ckzdrid] = ?,[ckkhid] = ?,[ckbz] = ?, " +
		"			[ckzdzsj] = ?,[ckdckid] = ?,[ckdbqjxj] = ?,[ckdcjjxj] = ? " +
		"		WHERE [ckdmxhh]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3xsckdt]( " +
		"			[ckdmxhh],[ckdjh],[ckdlsh],[ckdbj],[ckdcxbj], " +
		"			[ckdyyr],[ckdmdid],[ckdhpid],[ckddwid],[ckdhsl], " +
		"			[ckdjmsl],[ckdfjxx],[ckzdrid],[ckkhid],[ckbz], " +
		"			[ckzdzsj],[ckdckid],[ckdbqjxj],[ckdcjjxj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?) " +
		"	END"

	sqlUpdateMdZ3XsTht = "" +
		"IF EXISTS (SELECT * FROM [z3xsthdt] WHERE [thhpmxhh] = ?) " +
		"	BEGIN " +
		"		UPDATE [z3xsthdt] " +
		"		SET [thhpmxhh] = ?,[thdjh] = ?,[thlsh] = ?,[thyyrq] = ?,[thmdid] = ?, " +
		"		   [thhpid] = ?,[thjldwid] = ?,[thzhl] = ?,[thjmsl] = ?,[thhpfjxx] = ?, " +
		"		   [thzdrid] = ?,[thkhid] = ?,[thbz] = ?,[thzdzsj] = ?,[thqtckid] = ?, " +
		"		   [thbqjjexj] = ?,[thzzcjjexj] = ? " +
		"		WHERE [thhpmxhh] = ? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3xsthdt]( " +
		"			[thhpmxhh],[thdjh],[thlsh],[thyyrq],[thmdid], " +
		"			[thhpid],[thjldwid],[thzhl],[thjmsl],[thhpfjxx], " +
		"			[thzdrid],[thkhid],[thbz],[thzdzsj],[thqtckid], " +
		"			[thbqjjexj],[thzzcjjexj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?) " +
		"	END"

	sqlUpdateMdZ3MdDbCkDt = "" +
		"IF EXISTS (SELECT * FROM [z3mddbckdt] WHERE [dbdmxhh]=?) " +
		"	BEGIN " +
		"		UPDATE [z3mddbckdt] " +
		"		SET [dbdmxhh] = ?,[dbddjh] = ?,[dbdlsh] = ?,[dbdckr] = ?,[dbddcjgid] = ?, " +
		"			[dbdhpid] = ?,[dbddwid] = ?,[dbdhsl] = ?,[dbdjmsl] = ?,[dbdzdrid] = ?, " +
		"			[dbdbz] = ?,[dbdshsj] = ?,[dbddcckid] = ?,[dbdlsj] = ?,[dbddhj] = ?, " +
		"			[dbrkshjgid] = ?,[dbrkppid] = ? " +
		"		WHERE [dbdmxhh] = ? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3mddbckdt]( " +
		"			[dbdmxhh],[dbddjh],[dbdlsh],[dbdckr],[dbddcjgid], " +
		"			[dbdhpid],[dbddwid],[dbdhsl],[dbdjmsl],[dbdzdrid], " +
		"			[dbdbz],[dbdshsj],[dbddcckid],[dbdlsj],[dbddhj], " +
		"			[dbrkshjgid],[dbrkppid]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?) " +
		"	END"

	sqlUpdateMdZ3MdDbTzDt = "" +
		"IF EXISTS (SELECT * FROM [z3mddbtzdt] WHERE [tzdmxhh]=?) " +
		"	BEGIN " +
		"		UPDATE [z3mddbtzdt] " +
		"		SET [tzdMxHh]=?,[tzdDjh]=?,[tzdLsh]=?,[tzdCkr]=?,[tzdJgId]=?, " +
		"			[tzdCkId]=?,[tzRkJgId]=?,[tzRkCkId]=?,[tzdHpId]=?,[tzdDwId]=?, " +
		"			[tzdHsl]=?,[tzdJmTzSl]=?,[tzdLsj]=?,[tzdDhj]=?,[tzdZdrId]=?, " +
		"			[tzdBz]=?,[tzdShSj]=? " +
		"		WHERE [tzdmxhh]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3mddbtzdt]( " +
		"			[tzdMxHh],[tzdDjh],[tzdLsh], [tzdCkr], [tzdJgId], " +
		"			[tzdCkId],[tzRkJgId],[tzRkCkId],[tzdHpId],[tzdDwId], " +
		"			[tzdHsl],[tzdJmTzSl],[tzdLsj],[tzdDhj],[tzdZdrId], " +
		"			[tzdBz],[tzdShSj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?) " +
		"	END"

	sqlUpdateZ3HpCkDjDt = "" +
		"IF EXISTS (SELECT * FROM [z3hpckdjdt] WHERE [ckdmxhh]=?) " +
		"	BEGIN " +
		"		UPDATE [z3hpckdjdt] " +
		"		SET [ckdMxHh]=?,[ckdDjh]=?,[ckdLsh]=?,[ckdYyr]=?,[ckdCkFzJgId]=?, " +
		"			[ckdCkId]=?,[ckdHpId]=?,[ckdDwId]=?,[ckdHsl]=?,[ckdJmSl]=?, " +
		"			[ckdLsj]=?,[ckdDhj]=?,[ckBz]=?,[ckShrId]=?,[ckShSj]=? " +
		"		WHERE [ckdmxhh]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3hpckdjdt]( " +
		"			[ckdMxHh],[ckdDjh],[ckdLsh],[ckdYyr],[ckdCkFzJgId], " +
		"			[ckdCkId],[ckdHpId],[ckdDwId],[ckdHsl],[ckdJmSl], " +
		"			[ckdLsj],[ckdDhj],[ckBz],[ckShrId],[ckShSj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?) " +
		"	END"

	sqoUpdateZ3HpRkDjDt = "" +
		"IF EXISTS (SELECT * FROM [z3hprkdjdt] WHERE [rkdmxhh]=?) " +
		"	BEGIN " +
		"		UPDATE [z3hprkdjdt] " +
		"		SET [rkdMxHh]=?,[rkdDjh]=?,[rkdLsh]=?,[rkdYyr]=?,[rkdRkFzJgId]=?, " +
		"			[rkdCkId]=?,[rkdHpId]=?,[rkdDwId]=?,[rkdHsl]=?,[rkdJmSl]=?, " +
		"			[rkdDhj]=?,[rkdLsj]=?,[rkBz]=?,[rkShrId]=?,[rkShSj]=? " +
		"		WHERE [rkdmxhh]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3hprkdjdt]( " +
		"			[rkdMxHh],[rkdDjh],[rkdLsh],[rkdYyr],[rkdRkFzJgId], " +
		"			[rkdCkId],[rkdHpId],[rkdDwId],[rkdHsl],[rkdJmSl], " +
		"			[rkdDhj],[rkdLsj],[rkBz],[rkShrId],[rkShSj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?) " +
		"	END"

	sqlUpdateZ3PkDjDt = "" +
		"IF EXISTS (SELECT * FROM [z3pkdjdt] WHERE [pkdmxhh]=?) " +
		"	BEGIN " +
		"		UPDATE [z3pkdjdt] " +
		"		SET [pkdMxHh]=?,[pkdDjh]=?,[pkdLsh]=?,[pkdYyr]=?,[pkdPdJgId]=?, " +
		"			[pkdPdCkId]=?,[pkdHpId]=?,[pkdDwId]=?,[pkdHsl]=?,[pkdJmSl]=?, " +
		"			[pkdLsj]=?,[pkdDhj]=?,[pkdBz]=?,[pkdShrId]=?,[pkdShSj]=? " +
		"		WHERE [pkdmxhh]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3pkdjdt]( " +
		"			[pkdMxHh],[pkdDjh],[pkdLsh],[pkdYyr],[pkdPdJgId], " +
		"			[pkdPdCkId],[pkdHpId],[pkdDwId],[pkdHsl],[pkdJmSl], " +
		"			[pkdLsj],[pkdDhj],[pkdBz],[pkdShrId],[pkdShSj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?) " +
		"	END"

	sqlUpdateZ3MdPsTzDt = "" +
		"IF EXISTS (SELECT * FROM [z3mdpstzdt] WHERE [tzdmxhh]=?) " +
		"	BEGIN " +
		"		UPDATE [z3mdpstzdt] " +
		"		SET [tzdMxHh]=?,[tzdDjh]=?,[tzdLsh]=?,[tzdCkr]=?,[tzdShJgId]=?, " +
		"			[tzdShCkId]=?,[tzdPsJgId]=?,[tzdPsCkId]=?,[tzdHpId]=?,[tzdDwId]=?, " +
		"			[tzdHsl]=?,[tzdJmSl]=?,[tzdDhj]=?,[tzdPsj]=?,[tzdBz]=?, " +
		"			[tzdShrId]=?,[tzdShSj]=? " +
		"		WHERE [tzdmxhh]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3mdpstzdt]( " +
		"			[tzdMxHh],[tzdDjh],[tzdLsh],[tzdCkr],[tzdShJgId], " +
		"			[tzdShCkId],[tzdPsJgId],[tzdPsCkId],[tzdHpId],[tzdDwId], " +
		"			[tzdHsl],[tzdJmSl],[tzdDhj],[tzdPsj],[tzdBz], " +
		"			[tzdShrId],[tzdShSj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?," +
		"			?,?) " +
		"	END"

	sqlUpdateZ3XsDdThDt = "" +
		"IF EXISTS (SELECT * FROM [z3xsddthdt] WHERE [thdmxhh]=?) " +
		"	BEGIN " +
		"		UPDATE [z3xsddthdt] " +
		"		SET [ThdMxHh]=?,[ThdDjh]=?,[ThdDdLsh]=?,[ThdLsh]=?,[ThdDjLx]=?, " +
		"			[ThdYyr]=?,[ThdJfMdId]=?,[ThdCkId]=?,[ThdKhPpid]=?,[ThPhBj]=?, " +
		"			[ThKhId]=?,[ThdHpId]=?,[ThdDwId]=?,[ThdHsl]=?,[ThdDZxSl]=?, " +
		"			[ThDdBqjXj]=?,[ThDdCjjXj]=?,[ThdJmSl]=?,[ThdBqjXj]=?,[ThdCjjXj]=?, " +
		"			[ThZdrId]=?,[ThZdSj]=?,[ThBz]=? " +
		"		WHERE [thdmxhh]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3xsddthdt]( " +
		"			[ThdMxHh],[ThdDjh],[ThdDdLsh],[ThdLsh],[ThdDjLx], " +
		"			[ThdYyr],[ThdJfMdId],[ThdCkId],[ThdKhPpid],[ThPhBj], " +
		"			[ThKhId],[ThdHpId],[ThdDwId],[ThdHsl],[ThdDZxSl], " +
		"			[ThDdBqjXj],[ThDdCjjXj],[ThdJmSl],[ThdBqjXj],[ThdCjjXj], " +
		"			[ThZdrId],[ThZdSj],[ThBz]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?) " +
		"	END"

	sqlUpdateZ3Hpa = "" +
		"IF EXISTS (SELECT * FROM [z3hpa] WHERE [hpid]=?) " +
		"	BEGIN " +
		"		UPDATE [z3hpa] " +
		"		SET [hpid]=?,[hpbm]=?,[hpszmc]=?,[hpppid]=?,[hpsym]=?, " +
		"			[hppym]=?,[hpzxdwid]=?,[hpzxdwmc]=?,[hpejflid]=?,[hpmdszbj]=?, " +
		"			[hpdkhdhhpbj]=?,[hpwddhhpbj]=?,[hpms]=?,[hpzybj]=?,[hpisforbidden]=?, " +
		"			[hpszid]=?,[hpnsflm]=?,[hpjrhp]=?,[hpdghp]=? " +
		"		WHERE [hpid]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3hpa]( " +
		"			[hpid],[hpbm],[hpszmc],[hpppid],[hpsym]," +
		" 			[hppym],[hpzxdwid],[hpzxdwmc],[hpejflid],[hpmdszbj], " +
		"			[hpdkhdhhpbj],[hpwddhhpbj],[hpms],[hpzybj],[hpisforbidden], " +
		"			[hpszid],[hpnsflm],[hpjrhp],[hpdghp]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?) " +
		"	END"
	sqlDelZ3Hpa = "" +
		"delete from [z3hpa] " +
		"where [hpid] = ?"

	sqlUpdateZ3HpDwFja = "" +
		"IF EXISTS (SELECT * FROM [z3hpdwfja] WHERE [dwfjhpid]=? and [dwfjdwid] = ?) " +
		"	BEGIN " +
		"		UPDATE [z3hpdwfja] " +
		"		SET [dwfjhpid]=?,[dwfjdwid]=?,[dwfjdwlx]=?,[dwfjhsl]=?,[dwfjtxm]=?,[dwfjisforbidden]=? " +
		"		WHERE [dwfjhpid]=? and [dwfjdwid] = ? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3hpdwfja]([dwfjhpid],[dwfjdwid],[dwfjdwlx],[dwfjhsl],[dwfjtxm],[dwfjisforbidden]) " +
		"		VALUES (?,?,?,?,?,?) " +
		"	END"
	sqlDelZ3HpDwFja = "" +
		"delete from [z3hpdwfja] " +
		"where [dwfjhpid]=? and [dwfjdwid] = ?"

	sqlUpdateZ3JlDwa = "" +
		"IF EXISTS (SELECT * FROM [z3jldwa] WHERE [dwid]=?) " +
		"	BEGIN " +
		"		UPDATE [z3jldwa] " +
		"		SET [dwid]=?,[dwszmc]=?,[dwglmc]=?,[dwsym]=?,[dwpym]=?,[dwisforbidden]=? " +
		"		WHERE [dwid]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3jldwa]([dwid],[dwszmc],[dwglmc],[dwsym],[dwpym],[dwisforbidden]) " +
		"		VALUES (?,?,?,?,?,?) " +
		"	END"
	sqlDelZ3JlDwa = "" +
		"delete from [z3jldwa] " +
		"where [dwid]=?"
)

type repZxZc struct {
	dbConfig *goToolMSSql.MSSqlConfig
}

func NewRepZxZc() (*repZxZc, error) {
	comm := NewCommon()
	dbConfig, err := comm.GetOnLineDbConfig()
	if err != nil {
		return nil, err
	}
	return &repZxZc{
		dbConfig: dbConfig,
	}, nil
}

//==============================================================================================

//销售出库货品明细
func (r *repZxZc) UpdateMdZ3XsCkDt(d *object.Z3XsCkDt) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateMdZ3XsCkDt,
		d.CkdMxHh,
		d.CkdMxHh, d.CkDjh, d.CkdLsh, d.CkdBj, d.CkdCxBj,
		d.CkdYyr, d.CkdMdId, d.CkdHpId, d.CkdDwId, d.CkdHsl,
		d.CkdJmSl, d.CkdFjXx, d.CkZdRid, d.CkKhId, d.CkBz,
		d.CkZdzSj, d.CkdCkId, d.CkdBqjXj, d.CkdCjjXj,
		d.CkdMxHh,
		d.CkdMxHh, d.CkDjh, d.CkdLsh, d.CkdBj, d.CkdCxBj,
		d.CkdYyr, d.CkdMdId, d.CkdHpId, d.CkdDwId, d.CkdHsl,
		d.CkdJmSl, d.CkdFjXx, d.CkZdRid, d.CkKhId, d.CkBz,
		d.CkZdzSj, d.CkdCkId, d.CkdBqjXj, d.CkdCjjXj,
	)
}

//门店销售退货明细
func (r *repZxZc) UpdateMdZ3XsTht(d *object.Z3XsTht) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateMdZ3XsTht,
		d.ThHpMxHh,
		d.ThHpMxHh, d.ThDjh, d.ThLsh, d.ThYyRq, d.ThMdid,
		d.ThHpId, d.ThJlDwId, d.ThZhl, d.ThJmSl, d.ThHpFjXx,
		d.ThZdrId, d.ThKhId, d.ThBz, d.ThZdzSj, d.ThQtCkId,
		d.ThBqjJeXj, d.ThZzCjJeXj,
		d.ThHpMxHh,
		d.ThHpMxHh, d.ThDjh, d.ThLsh, d.ThYyRq, d.ThMdid,
		d.ThHpId, d.ThJlDwId, d.ThZhl, d.ThJmSl, d.ThHpFjXx,
		d.ThZdrId, d.ThKhId, d.ThBz, d.ThZdzSj, d.ThQtCkId,
		d.ThBqjJeXj, d.ThZzCjJeXj)
}

//调拨出库明细
func (r *repZxZc) UpdateMdZ3MdDbCkDt(d *object.Z3MdDbCkDt) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateMdZ3MdDbCkDt,
		d.DbdMxHh,
		d.DbdMxHh, d.DbdDjh, d.DbdLsh, d.DbdCkr, d.DbdDcJgId,
		d.DbdHpId, d.DbdDwId, d.DbdHsl, d.DbdJmSl, d.DbdZdrId,
		d.DbdBz, d.DbdShSj, d.DbdDcCkId, d.DbdLsj, d.DbdDhj,
		d.DbrKShJgId, d.DbrKPpId,
		d.DbdMxHh,
		d.DbdMxHh, d.DbdDjh, d.DbdLsh, d.DbdCkr, d.DbdDcJgId,
		d.DbdHpId, d.DbdDwId, d.DbdHsl, d.DbdJmSl, d.DbdZdrId,
		d.DbdBz, d.DbdShSj, d.DbdDcCkId, d.DbdLsj, d.DbdDhj,
		d.DbrKShJgId, d.DbrKPpId)
}

//调拨调整明细
func (r *repZxZc) UpdateMdZ3MdDbTzDt(d *object.Z3MdDbTzDt) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateMdZ3MdDbTzDt,
		d.TzdMxHh,
		d.TzdMxHh, d.TzdDjh, d.TzdLsh, d.TzdCkr, d.TzdJgId,
		d.TzdCkId, d.TzRkJgId, d.TzRkCkId, d.TzdHpId, d.TzdDwId,
		d.TzdHsl, d.TzdJmTzSl, d.TzdLsj, d.TzdDhj, d.TzdZdrId,
		d.TzdBz, d.TzdShSj,
		d.TzdMxHh,
		d.TzdMxHh, d.TzdDjh, d.TzdLsh, d.TzdCkr, d.TzdJgId,
		d.TzdCkId, d.TzRkJgId, d.TzRkCkId, d.TzdHpId, d.TzdDwId,
		d.TzdHsl, d.TzdJmTzSl, d.TzdLsj, d.TzdDhj, d.TzdZdrId,
		d.TzdBz, d.TzdShSj)
}

//货品出库登记
func (r *repZxZc) UpdateMdZ3HpCkDjDt(d *object.Z3HpCkDjDt) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateZ3HpCkDjDt,
		d.CkdMxHh,
		d.CkdMxHh, d.CkdDjh, d.CkdLsh, d.CkdYyr, d.CkdCkFzJgId,
		d.CkdCkId, d.CkdHpId, d.CkdDwId, d.CkdHsl, d.CkdJmSl,
		d.CkdLsj, d.CkdDhj, d.CkBz, d.CkShrId, d.CkShSj,
		d.CkdMxHh,
		d.CkdMxHh, d.CkdDjh, d.CkdLsh, d.CkdYyr, d.CkdCkFzJgId,
		d.CkdCkId, d.CkdHpId, d.CkdDwId, d.CkdHsl, d.CkdJmSl,
		d.CkdLsj, d.CkdDhj, d.CkBz, d.CkShrId, d.CkShSj)
}

//货品入库登记
func (r *repZxZc) UpdateMdZ3HpRkDjDt(d *object.Z3HpRkDjDt) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqoUpdateZ3HpRkDjDt,
		d.RkdMxHh,
		d.RkdMxHh, d.RkdDjh, d.RkdLsh, d.RkdYyr, d.RkdRkFzJgId,
		d.RkdCkId, d.RkdHpId, d.RkdDwId, d.RkdHsl, d.RkdJmSl,
		d.RkdDhj, d.RkdLsj, d.RkBz, d.RkShrId, d.RksShSj,
		d.RkdMxHh,
		d.RkdMxHh, d.RkdDjh, d.RkdLsh, d.RkdYyr, d.RkdRkFzJgId,
		d.RkdCkId, d.RkdHpId, d.RkdDwId, d.RkdHsl, d.RkdJmSl,
		d.RkdDhj, d.RkdLsj, d.RkBz, d.RkShrId, d.RksShSj)
}

//盘亏登记
func (r *repZxZc) UpdateMdZ3PkDjDt(d *object.Z3PkDjDt) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateZ3PkDjDt,
		d.PkdMxHh,
		d.PkdMxHh, d.PkdDjh, d.PkdLsh, d.PkdYyr, d.PkdPdJgId,
		d.PkdPdCkId, d.PkdHpId, d.PkdDwId, d.PkdHsl, d.PkdJmSl,
		d.PkdLsj, d.PkdDhj, d.PkdBz, d.PkdShrId, d.PkdShSj,
		d.PkdMxHh,
		d.PkdMxHh, d.PkdDjh, d.PkdLsh, d.PkdYyr, d.PkdPdJgId,
		d.PkdPdCkId, d.PkdHpId, d.PkdDwId, d.PkdHsl, d.PkdJmSl,
		d.PkdLsj, d.PkdDhj, d.PkdBz, d.PkdShrId, d.PkdShSj)
}

//门店配送调整
func (r *repZxZc) UpdateMdZ3MdPsTzDt(d *object.Z3MdPsTzDt) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateZ3MdPsTzDt,
		d.TzdMxHh,
		d.TzdMxHh, d.TzdDjh, d.TzdLsh, d.TzdCkr, d.TzdShJgId,
		d.TzdShCkId, d.TzdPsJgId, d.TzdPsCkId, d.TzdHpId, d.TzdDwId,
		d.TzdHsl, d.TzdJmSl, d.TzdDhj, d.TzdPsj, d.TzdBz,
		d.TzdShrId, d.TzdShSj,
		d.TzdMxHh,
		d.TzdMxHh, d.TzdDjh, d.TzdLsh, d.TzdCkr, d.TzdShJgId,
		d.TzdShCkId, d.TzdPsJgId, d.TzdPsCkId, d.TzdHpId, d.TzdDwId,
		d.TzdHsl, d.TzdJmSl, d.TzdDhj, d.TzdPsj, d.TzdBz,
		d.TzdShrId, d.TzdShSj)
}

//提货货品明细
func (r *repZxZc) UpdateMdZ3XsDdThDt(d *object.Z3XsDdThDt) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateZ3XsDdThDt,
		d.ThdMxHh,
		d.ThdMxHh, d.ThdDjh, d.ThdDdLsh, d.ThdLsh, d.ThdDjLx,
		d.ThdYyr, d.ThdJfMdId, d.ThdCkId, d.ThdKhPpid, d.ThPhBj,
		d.ThKhId, d.ThdHpId, d.ThdDwId, d.ThdHsl, d.ThdDZxSl,
		d.ThDdBqjXj, d.ThDdCjjXj, d.ThdJmSl, d.ThdBqjXj, d.ThdCjjXj,
		d.ThZdrId, d.ThZdSj, d.ThBz,
		d.ThdMxHh,
		d.ThdMxHh, d.ThdDjh, d.ThdDdLsh, d.ThdLsh, d.ThdDjLx,
		d.ThdYyr, d.ThdJfMdId, d.ThdCkId, d.ThdKhPpid, d.ThPhBj,
		d.ThKhId, d.ThdHpId, d.ThdDwId, d.ThdHsl, d.ThdDZxSl,
		d.ThDdBqjXj, d.ThDdCjjXj, d.ThdJmSl, d.ThdBqjXj, d.ThdCjjXj,
		d.ThZdrId, d.ThZdSj, d.ThBz)
}

//==============================================================================================

//集团通用货品设置A
func (r *repZxZc) UpdateZ3Hpa(d *object.Z3Hpa) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateZ3Hpa,
		d.HpId,
		d.HpId, d.HpBm, d.HpSzMc, d.HpPpId, d.HpSym,
		d.HpPym, d.HpZxDwId, d.HpZxDwMc, d.HpEjFlId, d.HpMdSzBj,
		d.HpdDkhDhHpBj, d.HpWdDhHpBj, d.HpMs, d.HpZyNj, d.HpIsForbidden,
		d.HpSzId, d.HpNsFlm, d.HpJrHp, d.HpDgHp,
		d.HpId,
		d.HpId, d.HpBm, d.HpSzMc, d.HpPpId, d.HpSym,
		d.HpPym, d.HpZxDwId, d.HpZxDwMc, d.HpEjFlId, d.HpMdSzBj,
		d.HpdDkhDhHpBj, d.HpWdDhHpBj, d.HpMs, d.HpZyNj, d.HpIsForbidden,
		d.HpSzId, d.HpNsFlm, d.HpJrHp, d.HpDgHp)
}
func (r *repZxZc) DelZ3Hpa(id int64) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlDelZ3Hpa, id)
}

//货品计量单位附加表
func (r *repZxZc) UpdateZ3HpDwFja(d *object.Z3HpDwFja) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateZ3HpDwFja,
		d.DwFjHpId, d.DwFjDwId,
		d.DwFjHpId, d.DwFjDwId, d.DwFjDwLx, d.DwFjHsl, d.DwFjTxm, d.DwFjIsForbidden,
		d.DwFjHpId, d.DwFjDwId,
		d.DwFjHpId, d.DwFjDwId, d.DwFjDwLx, d.DwFjHsl, d.DwFjTxm, d.DwFjIsForbidden)
}
func (r *repZxZc) DelZ3HpDwFja(hpId, dwId int) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlDelZ3HpDwFja, hpId, dwId)
}

//计量单位设置
func (r *repZxZc) UpdateZ3JlDwa(d *object.Z3JlDwa) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlUpdateZ3JlDwa,
		d.DwId,
		d.DwId, d.DwSzMc, d.DwGlMc, d.DwSym, d.DwPym, d.DwIsForbidden,
		d.DwId,
		d.DwId, d.DwSzMc, d.DwGlMc, d.DwSym, d.DwPym, d.DwIsForbidden)
}
func (r *repZxZc) DelZ3JlDwa(id int64) error {
	comm := NewCommon()
	return comm.SetRowsBySQL(r.dbConfig, sqlDelZ3JlDwa, id)
}

//==============================================================================================

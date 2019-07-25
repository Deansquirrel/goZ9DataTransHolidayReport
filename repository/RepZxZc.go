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
func (r *repZxZc) UpdateZ3HpCkDjDt(d *object.Z3HpCkDjDt) error {
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
func (r *repZxZc) UpdateZ3HpRkDjDt(d *object.Z3HpRkDjDt) error {
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

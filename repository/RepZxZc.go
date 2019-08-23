package repository

import (
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"github.com/Deansquirrel/goZ9DataTransHolidayReport/object"
)

import log "github.com/Deansquirrel/goToolLog"

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

	sqlUpdateZ3HpEjFlt = "" +
		"IF EXISTS (SELECT * FROM [z3hpejflt] WHERE [ejflid]=?) " +
		"	BEGIN " +
		"		UPDATE [z3hpejflt] " +
		"		SET [ejflid]=?,[ejflmc]=?,[ejflbm]=?,[ejflppid]=?,[ejflsym]=?, " +
		"			[ejflpym]=?,[ejflhpyjflid]=? " +
		"		WHERE [ejflid]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3hpejflt]( " +
		"			[ejflid],[ejflmc],[ejflbm],[ejflppid],[ejflsym], " +
		"			[ejflpym],[ejflhpyjflid]) " +
		"		VALUES ( " +
		"			?,?,?,?,?," +
		" 			?,?) " +
		"	END"
	sqlDelZ3HpEjFlt = "" +
		"delete from [z3hpejflt] " +
		"where [ejflid]=?"

	sqlUpdateZ3KhDja = "" +
		"IF EXISTS (SELECT * FROM [z3khdja] WHERE [khid]=?) " +
		"	BEGIN " +
		"		UPDATE [z3khdja] " +
		"		SET [khid]=?,[khszmc]=?,[khqc]=?,[khbm]=?,[khnbygbj]=?, " +
		"			[khxb]=?,[khpym]=?,[khjb]=?,[khsjhm]=?,[khqtlsfs]=?, " +
		"			[khkhyt]=?,[khrcxyj]=?,[khrcxyjsydd]=?,[khrcyh]=?,[khflid]=?, " +
		"			[khlszz]=?,[khisforbidden]=? " +
		"		WHERE [khid]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3khdja]( " +
		"			[khid],[khszmc],[khqc],[khbm],[khnbygbj], " +
		"			[khxb],[khpym],[khjb],[khsjhm],[khqtlsfs], " +
		"			[khkhyt],[khrcxyj],[khrcxyjsydd],[khrcyh],[khflid]," +
		"			[khlszz],[khisforbidden]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?) " +
		"	END"

	sqlDelZ3KhDja = "" +
		"delete from [z3khdja] " +
		"where [khid]=?"

	sqlUpdateZ3JrDzqSzFja = "" +
		"IF EXISTS (SELECT * FROM [z3jrdzqszfja] WHERE [dzqid]=?) " +
		"	BEGIN " +
		"		UPDATE [z3jrdzqszfja] " +
		"		SET [dzqId]=?,[dzqDpqHpId]=?,[dzqJyBz]=? " +
		"		WHERE [dzqid]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3jrdzqszfja]([dzqId],[dzqDpqHpId],[dzqJyBz]) " +
		"		VALUES (?,?,?) " +
		"	END"
	sqlDelZ3JrDzqSzFja = "" +
		"delete from [z3jrdzqszfja] " +
		"where [dzqid]=?"

	sqlUpdateZ3FzJga = "" +
		"IF EXISTS (SELECT * FROM [z3fzjga] WHERE [fzjgid]=?) " +
		"	BEGIN " +
		"		UPDATE [z3fzjga] " +
		"		SET [fzjgid]=?,[fzjgszmc]=?,[fzjgqc]=?,[fzjglx]=?,[fzjgsxrq]=?, " +
		"			[fzjgpym]=?,[fzjgsym]=?,[fzjgtdm]=?,[fzjgxsm]=?,[fzjglszz]=?, " +
		"			[fzjgisforbidden]=? " +
		"		WHERE [fzjgid]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3fzjga]([fzjgid],[fzjgszmc],[fzjgqc],[fzjglx],[fzjgsxrq], " +
		"			[fzjgpym],[fzjgsym],[fzjgtdm],[fzjgxsm],[fzjglszz], " +
		"			[fzjgisforbidden]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?) " +
		"	END"
	sqlDelZ3FzJga = "" +
		"delete from [z3jrdzqszfja] " +
		"where [dzqid]=?"

	sqlUpdateZ3DzqSza = "" +
		"IF EXISTS (SELECT * FROM [z3dzqsza] WHERE [dzqid]=?) " +
		"	BEGIN " +
		"		UPDATE [z3dzqsza] " +
		"		SET [dzqid]=?,[dzqszmc]=?,[dzqdwxsmc]=?,[dzqbm]=?,[dzqpym]=?, " +
		"			[dzqjrqbj]=?,[dzqzzqmz]=?,[dzqyt]=?,[dzqhxhly]=?,[dzqdbbs]=?, " +
		"			[dzqhsxz]=?,[dzqctqid]=?,[dzqzdzkl]=?,[dzqqxzkyh]=?,[dzqqhsjzr]=?, " +
		"			[dzqjybz]=? " +
		"		WHERE [dzqid]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3dzqsza]([dzqid],[dzqszmc],[dzqdwxsmc],[dzqbm],[dzqpym], " +
		"			[dzqjrqbj],[dzqzzqmz],[dzqyt],[dzqhxhly],[dzqdbbs], " +
		"			[dzqhsxz],[dzqctqid],[dzqzdzkl],[dzqqxzkyh],[dzqqhsjzr], " +
		"			[dzqjybz]) 	" +
		"	VALUES ( " +
		"		?,?,?,?,?, " +
		"		?,?,?,?,?, " +
		"		?,?,?,?,?, " +
		"		?) " +
		"	END"
	sqlDelZ3DzqSza = "" +
		"delete from [z3dzqsza] " +
		"where [dzqid]=?"

	sqlUpdateZ3Ppa = "" +
		"IF EXISTS (SELECT * FROM [z3ppa] WHERE [ppid]=?) " +
		"	BEGIN " +
		"		UPDATE [z3ppa] " +
		"		SET [ppid]=?,[ppszmc]=?,[ppqc]=?,[ppyt]=?,[ppbm]=?, " +
		"			[ppsym]=?,[pppym]=?,[ppisforbidden]=?,[ppbz]=? " +
		"		WHERE [ppid]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [z3ppa]([ppid],[ppszmc],[ppqc],[ppyt],[ppbm], " +
		"			[ppsym],[pppym],[ppisforbidden],[ppbz]) " +
		"		VALUES ( " +
		"		?,?,?,?,?, " +
		"		?,?,?,?) " +
		"	END"
	sqlDelZ3Ppa = "" +
		"delete from [z3ppa] " +
		"where [ppid]=?"

	sqlUpdateZ3HxZzLsjt = "" +
		"IF EXISTS (SELECT * FROM [z3hxzzlsjt] WHERE [lsjjgzzid]=? and [lsjhpid]=?) " +
		"	Begin " +
		"		UPDATE [z3hxzzlsjt] " +
		"		SET [lsjjgzzid]=?,[lsjhpid]=?,[lsjyxzxlsj]=? " +
		"		WHERE [lsjjgzzid]=? and [lsjhpid]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		DELETE FROM [z3hxzzlsjt] WHERE [lsjhpid]=? " +
		"		INSERT INTO [z3hxzzlsjt]([lsjjgzzid],[lsjhpid],[lsjyxzxlsj]) " +
		"		VALUES (?,?,?) " +
		"	End"
	sqlDelZ3HxZzLsjt = "" +
		"delete from [z3hxzzlsjt] " +
		"where [lsjjgzzid]=? and [lsjhpid]=?"

	sqlUpdateZ3PsCkHpDt = "" +
		"IF EXISTS (SELECT * FROM [z3psckhpdt] WHERE [ckdlsh]=?) " +
		"	Begin " +
		"		UPDATE [z3psckhpdt] " +
		"		SET [ckdlsh]=?,[ckdhpid]=?,[ckddjh]=?,[ckddwid]=?,[ckdhsl]=?, " +
		"			[ckdcckfzjgid]=?,[ckdcckid]=?,[ckmdid]=?,[ckppid]=?,[ckdeddjlx]=?, " +
		"			[ckddhzxsl]=?,[ckdzxdhj]=?,[ckddhje]=?,[ckdzspsj]=?,[ckdhrq]=?, " +
		"			[ckdshsj]=?,[ckdpsjmsl]=?,[ckpsbz]=?,[ckfcczrid]=? " +
		"		WHERE [ckdlsh]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [z3psckhpdt]([ckdlsh],[ckdhpid],[ckddjh],[ckddwid],[ckdhsl], " +
		"			[ckdcckfzjgid],[ckdcckid],[ckmdid],[ckppid],[ckdeddjlx], " +
		"			[ckddhzxsl],[ckdzxdhj],[ckddhje],[ckdzspsj],[ckdhrq], " +
		"			[ckdshsj],[ckdpsjmsl],[ckpsbz],[ckfcczrid]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?) " +
		"	End"

	sqlUpdateZ3PsCkHpFjDt = "" +
		"IF EXISTS (SELECT * FROM [z3psckhpfjdt] WHERE [ckdlsh]=?) " +
		"	Begin " +
		"		UPDATE [z3psckhpfjdt] " +
		"		SET [ckdlsh]=?,[ckdhpid]=?,[ckghdmxhh]=?,[ckddwid]=?,[ckdhsl]=?, " +
		"			[ckzxddsl]=?,[ckzxpssl]=?,[ckphbj]=? " +
		"		WHERE [ckdlsh]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [z3psckhpfjdt]([ckdlsh],[ckdhpid],[ckghdmxhh],[ckddwid],[ckdhsl], " +
		"			[ckzxddsl],[ckzxpssl],[ckphbj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?) " +
		"	End"
	sqlUpdateZ3PsXzDt = "" +
		"IF EXISTS (SELECT * FROM [Z3PsXzDt] WHERE [xzdmxhh]=?) " +
		"	Begin " +
		"		UPDATE [Z3PsXzDt] " +
		"		SET [xzdmxhh]=?,[xzdjh]=?,[xzlsh]=?,[xzdxzr]=?,[xzdshjgid]=?, " +
		"			[xzdshckid]=?,[xzrkjgid]=?,[xzrkckid]=?,[xzdhpid]=?,[xzddwid]=?, " +
		"			[xzdhsl]=?,[xzdjmtzsl]=?,[xzdzxtzsl]=?,[xzdjmxzsl]=?,[xzddhj]=?, " +
		"			[xzdpsj]=?,[xzbz]=?,[xzrid]=?,[xzsj]=? " +
		"		WHERE [xzdmxhh]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [Z3PsXzDt]([xzdmxhh],[xzdjh],[xzlsh],[xzdxzr],[xzdshjgid], " +
		"			[xzdshckid],[xzrkjgid],[xzrkckid],[xzdhpid],[xzddwid], " +
		"			[xzdhsl],[xzdjmtzsl],[xzdzxtzsl],[xzdjmxzsl],[xzddhj], " +
		"			[xzdpsj],[xzbz],[xzrid],[xzsj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?) " +
		"	End"

	sqlUpdateZ3MdPsChDt = "" +
		"IF EXISTS (SELECT * FROM [z3mdpschdt] WHERE [chdlsh]=? AND [chdhpid]=?) " +
		"	Begin " +
		"		UPDATE [z3mdpschdt] " +
		"		SET [chdlsh]=?,[chdjh]=?,[chlsh]=?,[chdhpid]=?,[chdcckfzjgid]=?, " +
		"			[chdcckid]=?,[chdhsr]=?,[chdhsxzid]=?,[chdhsbmid]=?,[chddwid]=?, " +
		"			[chdhsl]=?,[chdpsjmsl]=?,[chdpsj]=?,[chbz]=?,[chshrid]=?, " +
		"			[chshsj]=? " +
		"		WHERE [chdlsh]=? AND [chdhpid]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [z3mdpschdt]([chdlsh],[chdjh],[chlsh],[chdhpid],[chdcckfzjgid], " +
		"			[chdcckid],[chdhsr],[chdhsxzid],[chdhsbmid],[chddwid], " +
		"			[chdhsl],[chdpsjmsl],[chdpsj],[chbz],[chshrid], " +
		"			[chshsj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?) " +
		"	End"

	sqlUpdateShengChWgRkDt = "" +
		"IF EXISTS (SELECT * FROM [shengchwgrkdt] WHERE [wgrkmxhh]=?) " +
		"	Begin " +
		"		UPDATE [shengchwgrkdt] " +
		"		SET [wgrkmxhh]=?,[wgrkdjh]=?,[wgrklsh]=?,[wgrkhsrq]=?,[wgrkhpid]=?, " +
		"			[wgrkdwid]=?,[wgrkhsl]=?,[wgrkrfzjgid]=?,[wgrkrckid]=?,[wgrkjmsl]=?, " +
		"			[wgrkcpczdj]=?,[wgrkjg]=?,[wgrkzdrid]=?,[wgrkshsj]=?,[wgrkbz]=? " +
		"		WHERE [wgrkmxhh]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [shengchwgrkdt]([wgrkmxhh],[wgrkdjh],[wgrklsh],[wgrkhsrq],[wgrkhpid], " +
		"			[wgrkdwid],[wgrkhsl],[wgrkrfzjgid],[wgrkrckid],[wgrkjmsl], " +
		"			[wgrkcpczdj],[wgrkjg],[wgrkzdrid],[wgrkshsj],[wgrkbz]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?) " +
		"	End"

	sqlUpdateZ3SheZhXsHpMxt = "" +
		"IF EXISTS (SELECT * FROM [z3shezhxshpmxt] WHERE [ckmxhh]=?) " +
		"	Begin " +
		"		UPDATE [z3shezhxshpmxt] " +
		"		SET [ckmxhh]=?,[ckdjh]=?,[cklsh]=?,[ckhsrq]=?,[ckcckfzjgid]=?, " +
		"			[ckcckid]=?,[ckhpid]=?,[ckdwid]=?,[ckhsl]=?,[ckjmdj]=?, " +
		"			[ckjmsl]=?,[ckcjjexj]=?,[ckjmlsj]=?,[cklsjexj]=?,[ckzdrid]=?, " +
		"			[ckshsj]=?,[ckbz]=? " +
		"		WHERE [ckmxhh]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [z3shezhxshpmxt]([ckmxhh],[ckdjh],[cklsh],[ckhsrq],[ckcckfzjgid], " +
		"			[ckcckid],[ckhpid],[ckdwid],[ckhsl],[ckjmdj], " +
		"			[ckjmsl],[ckcjjexj],[ckjmlsj],[cklsjexj],[ckzdrid], " +
		"			[ckshsj],[ckbz]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?) " +
		"	End"

	sqlUpdateXtTz = "" +
		"IF EXISTS (SELECT * FROM [xttz] WHERE [dptid]=? AND [brid]=? AND [gsid]=?) " +
		"	BEGIN " +
		"		UPDATE [xttz] " +
		"		SET [gsqty]=?,[updatedate]=? " +
		"		WHERE [dptid]=? AND [brid]=? AND [gsid]=? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO [xttz]([dptid],[brid],[gsid],[gsqty],[updatedate]) " +
		"		VALUES (?,?,?,?,?) " +
		"	END"

	sqlUpdateOoDxV1DdShCkt = "" +
		"IF EXISTS (SELECT * FROM [oodxv1ddshckt] WHERE [shmxhh]=?) " +
		"	Begin " +
		"		UPDATE [oodxv1ddshckt] " +
		"		SET [shlsh]=?,[shmxhh]=?,[shhsr]=?,[shfzjgid]=?,[shckid]=?, " +
		"			[shbmid]=?,[shxzid]=?,[shsendrid]=?,[shkhid]=?,[shdddjh]=?, " +
		"			[shddlsh]=?,[shddmxhh]=?,[shphbj]=?,[shhpid]=?,[shjldwid]=?, " +
		"			[shhsl]=?,[shddzxsl]=?,[shjhjmsl]=?,[shshjmsl]=?,[shshrid]=?, " +
		"			[shshsj]=? " +
		"		WHERE [shmxhh]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [oodxv1ddshckt]([shlsh],[shmxhh],[shhsr],[shfzjgid],[shckid], " +
		"			[shbmid],[shxzid],[shsendrid],[shkhid],[shdddjh], " +
		"			[shddlsh],[shddmxhh],[shphbj],[shhpid],[shjldwid], " +
		"			[shhsl],[shddzxsl],[shjhjmsl],[shshjmsl],[shshrid], " +
		"			[shshsj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?) " +
		"	End"

	sqlUpdateZ3MdDhDdDt = "" +
		"IF EXISTS (SELECT * FROM [z3mddhdddt] WHERE [dhdlsh]=? AND [dhdhpid]=?) " +
		"	Begin " +
		"		UPDATE [z3mddhdddt] " +
		"		SET [dhdlsh]=?,[dhdhpid]=?,[dhdjh]=?,[dhxsrq]=?,[dhdhwid]=?, " +
		"			[dhdhsl]=?,[dhdjljmsjdhs]=?,[dhddhj]=?,[dhdxdrid]=?,[dhdxdsj]=? " +
		"		WHERE [dhdlsh]=?  AND [dhdhpid]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [z3mddhdddt]([dhdlsh],[dhdhpid],[dhdjh],[dhxsrq],[dhdhwid], " +
		"			[dhdhsl],[dhdjljmsjdhs],[dhddhj],[dhdxdrid],[dhdxdsj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?) " +
		"	End"

	sqlUpdateZ3MdDhDdt = "" +
		"IF EXISTS (SELECT * FROM [z3mddhddt] WHERE [dhdmxhh]=?) " +
		"	Begin " +
		"		UPDATE [z3mddhddt] " +
		"		SET [dhdmxhh]=?,[dhddjh]=?,[dhdlsh]=?,[dhdyyrq]=?,[dhdmdid]=?, " +
		"			[dhdhpid]=?,[dhddwid]=?,[dhdhsl]=?,[dhdjljmdhs]=?,[dhdzxpsjg]=?, " +
		"			[dhdbz]=?,[dhdshrid]=?,[dhdshsj]=? " +
		"		WHERE [dhdmxhh]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [z3mddhddt]([dhdmxhh],[dhddjh],[dhdlsh],[dhdyyrq],[dhdmdid], " +
		"			[dhdhpid],[dhddwid],[dhdhsl],[dhdjljmdhs],[dhdzxpsjg], " +
		"			[dhdbz],[dhdshrid],[dhdshsj]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?,?,?) " +
		"	End"
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
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateMdZ3XsCkDt,
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
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//门店销售退货明细
func (r *repZxZc) UpdateMdZ3XsTht(d *object.Z3XsTht) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateMdZ3XsTht,
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
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//调拨出库明细
func (r *repZxZc) UpdateMdZ3MdDbCkDt(d *object.Z3MdDbCkDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateMdZ3MdDbCkDt,
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
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//调拨调整明细
func (r *repZxZc) UpdateMdZ3MdDbTzDt(d *object.Z3MdDbTzDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateMdZ3MdDbTzDt,
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
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//货品出库登记
func (r *repZxZc) UpdateMdZ3HpCkDjDt(d *object.Z3HpCkDjDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3HpCkDjDt,
		d.CkdMxHh,
		d.CkdMxHh, d.CkdDjh, d.CkdLsh, d.CkdYyr, d.CkdCkFzJgId,
		d.CkdCkId, d.CkdHpId, d.CkdDwId, d.CkdHsl, d.CkdJmSl,
		d.CkdLsj, d.CkdDhj, d.CkBz, d.CkShrId, d.CkShSj,
		d.CkdMxHh,
		d.CkdMxHh, d.CkdDjh, d.CkdLsh, d.CkdYyr, d.CkdCkFzJgId,
		d.CkdCkId, d.CkdHpId, d.CkdDwId, d.CkdHsl, d.CkdJmSl,
		d.CkdLsj, d.CkdDhj, d.CkBz, d.CkShrId, d.CkShSj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//货品入库登记
func (r *repZxZc) UpdateMdZ3HpRkDjDt(d *object.Z3HpRkDjDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqoUpdateZ3HpRkDjDt,
		d.RkdMxHh,
		d.RkdMxHh, d.RkdDjh, d.RkdLsh, d.RkdYyr, d.RkdRkFzJgId,
		d.RkdCkId, d.RkdHpId, d.RkdDwId, d.RkdHsl, d.RkdJmSl,
		d.RkdDhj, d.RkdLsj, d.RkBz, d.RkShrId, d.RksShSj,
		d.RkdMxHh,
		d.RkdMxHh, d.RkdDjh, d.RkdLsh, d.RkdYyr, d.RkdRkFzJgId,
		d.RkdCkId, d.RkdHpId, d.RkdDwId, d.RkdHsl, d.RkdJmSl,
		d.RkdDhj, d.RkdLsj, d.RkBz, d.RkShrId, d.RksShSj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//盘亏登记
func (r *repZxZc) UpdateMdZ3PkDjDt(d *object.Z3PkDjDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3PkDjDt,
		d.PkdMxHh,
		d.PkdMxHh, d.PkdDjh, d.PkdLsh, d.PkdYyr, d.PkdPdJgId,
		d.PkdPdCkId, d.PkdHpId, d.PkdDwId, d.PkdHsl, d.PkdJmSl,
		d.PkdLsj, d.PkdDhj, d.PkdBz, d.PkdShrId, d.PkdShSj,
		d.PkdMxHh,
		d.PkdMxHh, d.PkdDjh, d.PkdLsh, d.PkdYyr, d.PkdPdJgId,
		d.PkdPdCkId, d.PkdHpId, d.PkdDwId, d.PkdHsl, d.PkdJmSl,
		d.PkdLsj, d.PkdDhj, d.PkdBz, d.PkdShrId, d.PkdShSj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//门店配送调整
func (r *repZxZc) UpdateMdZ3MdPsTzDt(d *object.Z3MdPsTzDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3MdPsTzDt,
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
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//提货货品明细
func (r *repZxZc) UpdateMdZ3XsDdThDt(d *object.Z3XsDdThDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3XsDdThDt,
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
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//==============================================================================================

// 配送出库
func (r *repZxZc) UpdateZ3PsCkHpDt(d *object.Z3PsCkHpDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3PsCkHpDt,
		d.CkdLsh,
		d.CkdLsh, d.CkdHpId, d.CkdDjh, d.CkdDwId, d.CkdHsl,
		d.CkdCckFzJgId, d.CkdCckId, d.CkMdId, d.CkPpId, d.CkdEdDjLx,
		d.CkdDhZxSl, d.CkdZxDhj, d.CkdDhJe, d.CkdZsPsj, d.CkdHrq,
		d.CkdShSj, d.CkdPsJmSl, d.CkPsBz, d.CkFcCzrId,
		d.CkdLsh,
		d.CkdLsh, d.CkdHpId, d.CkdDjh, d.CkdDwId, d.CkdHsl,
		d.CkdCckFzJgId, d.CkdCckId, d.CkMdId, d.CkPpId, d.CkdEdDjLx,
		d.CkdDhZxSl, d.CkdZxDhj, d.CkdDhJe, d.CkdZsPsj, d.CkdHrq,
		d.CkdShSj, d.CkdPsJmSl, d.CkPsBz, d.CkFcCzrId)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

// 配送出库二级明细表
func (r *repZxZc) UpdateZ3PsCkHpFjDt(d *object.Z3PsCkHpFjDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3PsCkHpFjDt,
		d.CkdLsh,
		d.CkdLsh, d.CkdHpId, d.CkGhdMxHh, d.CkdDwId, d.CkdHsl,
		d.CkZxDdSl, d.CkZxPsSl, d.CkPhBj,
		d.CkdLsh,
		d.CkdLsh, d.CkdHpId, d.CkGhdMxHh, d.CkdDwId, d.CkdHsl,
		d.CkZxDdSl, d.CkZxPsSl, d.CkPhBj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//配送修正
func (r *repZxZc) UpdateZ3PsXzDt(d *object.Z3PsXzDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3PsXzDt,
		d.XzdMxHh,
		d.XzdMxHh, d.XzDjh, d.XzLsh, d.XzdXzr, d.XzdShJgId,
		d.XzdShCkId, d.XzRkJgId, d.XzRkCkId, d.XzdHpId, d.XzdDwId,
		d.XzdHsl, d.XzdJmTzSl, d.XzdZxTzSl, d.XzdJmXzSl, d.XzdDhj,
		d.XzdPsj, d.XzBz, d.XzrId, d.XzSj,
		d.XzdMxHh,
		d.XzdMxHh, d.XzDjh, d.XzLsh, d.XzdXzr, d.XzdShJgId,
		d.XzdShCkId, d.XzRkJgId, d.XzRkCkId, d.XzdHpId, d.XzdDwId,
		d.XzdHsl, d.XzdJmTzSl, d.XzdZxTzSl, d.XzdJmXzSl, d.XzdDhj,
		d.XzdPsj, d.XzBz, d.XzrId, d.XzSj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//配送冲红单
func (r *repZxZc) UpdateZ3MdPsChDt(d *object.Z3MdPsChDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3MdPsChDt,
		d.ChdLsh, d.ChdHpId,
		d.ChdLsh, d.ChDjh, d.ChLsh, d.ChdHpId, d.ChdCckFzJgId,
		d.ChdCckId, d.ChdHsr, d.ChdHsXzId, d.ChdHsBmId, d.ChdDwId,
		d.ChdHsl, d.ChdPsJmSl, d.ChdPsj, d.ChBz, d.ChShrId,
		d.ChShSj,
		d.ChdLsh, d.ChdHpId,
		d.ChdLsh, d.ChDjh, d.ChLsh, d.ChdHpId, d.ChdCckFzJgId,
		d.ChdCckId, d.ChdHsr, d.ChdHsXzId, d.ChdHsBmId, d.ChdDwId,
		d.ChdHsl, d.ChdPsJmSl, d.ChdPsj, d.ChBz, d.ChShrId,
		d.ChShSj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//完工入库
func (r *repZxZc) UpdateShengChWgRkDt(d *object.ShengChWgRkDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateShengChWgRkDt,
		d.WgRkMxHh,
		d.WgRkMxHh, d.WgRkDjh, d.WgRkLsh, d.WgRkHsRq, d.WgRkHpId,
		d.WgRkDwId, d.WgRkHsl, d.WgRkrFzJgId, d.WgRkrCkId, d.WgRkJmSl,
		d.WgRkCpCzDj, d.WgRkJg, d.WgRkZdrId, d.WgRkShSj, d.WgRkBz,
		d.WgRkMxHh,
		d.WgRkMxHh, d.WgRkDjh, d.WgRkLsh, d.WgRkHsRq, d.WgRkHpId,
		d.WgRkDwId, d.WgRkHsl, d.WgRkrFzJgId, d.WgRkrCkId, d.WgRkJmSl,
		d.WgRkCpCzDj, d.WgRkJg, d.WgRkZdrId, d.WgRkShSj, d.WgRkBz)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//赊账销售出库
func (r *repZxZc) UpdateZ3SheZhXsHpMxt(d *object.Z3SheZhXsHpMxt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3SheZhXsHpMxt,
		d.CkMxHh,
		d.CkMxHh, d.CkDjh, d.CkLsh, d.CkHsRq, d.CkCckFzJgId,
		d.CkCckId, d.CkHpId, d.CkDwId, d.CkHsl, d.CkJmDj,
		d.CkJmSl, d.CkCjJeXj, d.CkJmLsj, d.CkLsJeXj, d.CkZdrId,
		d.CkShSj, d.CkBz,
		d.CkMxHh,
		d.CkMxHh, d.CkDjh, d.CkLsh, d.CkHsRq, d.CkCckFzJgId,
		d.CkCckId, d.CkHpId, d.CkDwId, d.CkHsl, d.CkJmDj,
		d.CkJmSl, d.CkCjJeXj, d.CkJmLsj, d.CkLsJeXj, d.CkZdrId,
		d.CkShSj, d.CkBz)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//工厂即时台帐表
func (r *repZxZc) UpdateXtTz(d *object.XtTz) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateXtTz,
		d.DptId, d.BrId, d.GsId,
		d.GsQty, d.UpdateDate,
		d.DptId, d.BrId, d.GsId,
		d.DptId, d.BrId, d.GsId, d.GsQty, d.UpdateDate)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//订单提货单
func (r *repZxZc) UpdateOoDxV1DdShCkt(d *object.OoDxV1DdShCkt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateOoDxV1DdShCkt,
		d.ShMxHh,
		d.ShLsh, d.ShMxHh, d.ShHsr, d.ShFzJgId, d.ShCkId,
		d.ShBmId, d.ShXzId, d.ShSendRid, d.ShKhId, d.ShDdDjh,
		d.ShDdLsh, d.ShDdMxHh, d.ShPhBj, d.ShHpId, d.ShJlDwId,
		d.ShHsl, d.ShDdZxSl, d.ShJhJmSl, d.ShShJmSl, d.ShShrId,
		d.ShShSj,
		d.ShMxHh,
		d.ShLsh, d.ShMxHh, d.ShHsr, d.ShFzJgId, d.ShCkId,
		d.ShBmId, d.ShXzId, d.ShSendRid, d.ShKhId, d.ShDdDjh,
		d.ShDdLsh, d.ShDdMxHh, d.ShPhBj, d.ShHpId, d.ShJlDwId,
		d.ShHsl, d.ShDdZxSl, d.ShJhJmSl, d.ShShJmSl, d.ShShrId,
		d.ShShSj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//门店订货代单
func (r *repZxZc) UpdateZ3MdDhDdDt(d *object.Z3MdDhDdDt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3MdDhDdDt,
		d.DhdLsh, d.DhdHpId,
		d.DhdLsh, d.DhdHpId, d.DhDjh, d.DhXsRq, d.DhdHwId,
		d.DhdHsl, d.DhdJlJmSjDhs, d.DhdDhj, d.DhdXdrId, d.DhdXdSj,
		d.DhdLsh, d.DhdHpId,
		d.DhdLsh, d.DhdHpId, d.DhDjh, d.DhXsRq, d.DhdHwId,
		d.DhdHsl, d.DhdJlJmSjDhs, d.DhdDhj, d.DhdXdrId, d.DhdXdSj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//门店订货单
func (r *repZxZc) UpdateZ3MdDhDdt(d *object.Z3MdDhDdt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3MdDhDdt,
		d.DhdMxHh,
		d.DhdMxHh, d.DhdDjh, d.DhdLsh, d.DhdYyRq, d.DhdMdId,
		d.DhdHpId, d.DhdDwId, d.DhdHsl, d.DhdJlJmDhs, d.DhdZxPsJg,
		d.DhdBz, d.DhdShrId, d.DhdShSj,
		d.DhdMxHh,
		d.DhdMxHh, d.DhdDjh, d.DhdLsh, d.DhdYyRq, d.DhdMdId,
		d.DhdHpId, d.DhdDwId, d.DhdHsl, d.DhdJlJmDhs, d.DhdZxPsJg,
		d.DhdBz, d.DhdShrId, d.DhdShSj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//==============================================================================================

//集团通用货品设置A
func (r *repZxZc) UpdateZ3Hpa(d *object.Z3Hpa) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3Hpa,
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
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func (r *repZxZc) DelZ3Hpa(id int64) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelZ3Hpa, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//货品计量单位附加表
func (r *repZxZc) UpdateZ3HpDwFja(d *object.Z3HpDwFja) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3HpDwFja,
		d.DwFjHpId, d.DwFjDwId,
		d.DwFjHpId, d.DwFjDwId, d.DwFjDwLx, d.DwFjHsl, d.DwFjTxm, d.DwFjIsForbidden,
		d.DwFjHpId, d.DwFjDwId,
		d.DwFjHpId, d.DwFjDwId, d.DwFjDwLx, d.DwFjHsl, d.DwFjTxm, d.DwFjIsForbidden)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func (r *repZxZc) DelZ3HpDwFja(hpId, dwId int) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelZ3HpDwFja, hpId, dwId)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//计量单位设置
func (r *repZxZc) UpdateZ3JlDwa(d *object.Z3JlDwa) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3JlDwa,
		d.DwId,
		d.DwId, d.DwSzMc, d.DwGlMc, d.DwSym, d.DwPym, d.DwIsForbidden,
		d.DwId,
		d.DwId, d.DwSzMc, d.DwGlMc, d.DwSym, d.DwPym, d.DwIsForbidden)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func (r *repZxZc) DelZ3JlDwa(id int64) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelZ3JlDwa, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//货品二级分类
func (r *repZxZc) UpdateZ3HpEjFlt(d *object.Z3HpEjFlt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3HpEjFlt,
		d.EjFlId,
		d.EjFlId, d.EjFlMc, d.EjFlBm, d.EjFlPpId, d.EjFlSym,
		d.EjFlPym, d.EjFlHpYjFlId,
		d.EjFlId,
		d.EjFlId, d.EjFlMc, d.EjFlBm, d.EjFlPpId, d.EjFlSym,
		d.EjFlPym, d.EjFlHpYjFlId)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func (r *repZxZc) DelZ3HpEjFlt(id int64) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelZ3HpEjFlt, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//客户登记通用信息表
func (r *repZxZc) UpdateZ3KhDja(d *object.Z3KhDja) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3KhDja,
		d.KhId,
		d.KhId, d.KhSzMc, d.KhQc, d.KhBm, d.KhNbYgBj,
		d.KhXb, d.KhPym, d.KhJb, d.KhSjHm, d.KhQtLsFs,
		d.KhKHyt, d.KhRcXyj, d.KhRcXyjSyDd, d.KhrRcYh, d.KhFlId,
		d.KhLsZz, d.KhIsForbidden,
		d.KhId,
		d.KhId, d.KhSzMc, d.KhQc, d.KhBm, d.KhNbYgBj,
		d.KhXb, d.KhPym, d.KhJb, d.KhSjHm, d.KhQtLsFs,
		d.KhKHyt, d.KhRcXyj, d.KhRcXyjSyDd, d.KhrRcYh, d.KhFlId,
		d.KhLsZz, d.KhIsForbidden)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func (r *repZxZc) DelZ3KhDja(id int64) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelZ3KhDja, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//节日电子券设置附加表
func (r *repZxZc) UpdateZ3JrDzqSzFja(d *object.Z3JrDzqSzFja) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3JrDzqSzFja,
		d.DzqId,
		d.DzqId, d.DzqDpqHpId, d.DzqJyBz,
		d.DzqId,
		d.DzqId, d.DzqDpqHpId, d.DzqJyBz)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func (r *repZxZc) DelZ3JrDzqSzFja(id int64) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelZ3JrDzqSzFja, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//机构表A
func (r *repZxZc) UpdateZ3FzJga(d *object.Z3FzJga) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3FzJga,
		d.FzJgId,
		d.FzJgId, d.FzJgSzMc, d.FzJgQc, d.FzJgLx, d.FzJgsSxRq,
		d.FzJgPym, d.FzJgSym, d.FzJgTdm, d.FzJgXsm, d.FzJgLsZz,
		d.FzJgIsForbidden,
		d.FzJgId,
		d.FzJgId, d.FzJgSzMc, d.FzJgQc, d.FzJgLx, d.FzJgsSxRq,
		d.FzJgPym, d.FzJgSym, d.FzJgTdm, d.FzJgXsm, d.FzJgLsZz,
		d.FzJgIsForbidden)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func (r *repZxZc) DelZ3FzJga(id int64) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelZ3FzJga, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//核心组织零售价
func (r *repZxZc) UpdateZ3HxZzLsjt(d *object.Z3HxZzLsjt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3HxZzLsjt,
		d.LsjJgZzId, d.LsjHpid,
		d.LsjJgZzId, d.LsjHpid, d.LsjYxZxLsj,
		d.LsjJgZzId, d.LsjHpid,
		d.LsjHpid,
		d.LsjJgZzId, d.LsjHpid, d.LsjYxZxLsj)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func (r *repZxZc) DelZ3HxZzLsjt(id *object.Z3HxZzLsjSyt) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelZ3HxZzLsjt, id.LsjZzId, id.LsjHpId)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//电子券设置
func (r *repZxZc) UpdateZ3DzqSza(d *object.Z3DzqSza) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3DzqSza,
		d.DzqId,
		d.DzqId, d.DzqSzMc, d.DzqDwXsMc, d.DzqBm, d.DzqPym,
		d.DzqJrqBj, d.DzqZzqMz, d.DzqYt, d.DzqHxhLy, d.DzqDbBs,
		d.DzqHsXz, d.DzqCtqId, d.DzqZdZkl, d.DzqQxZkYh, d.DzqQhsJzr,
		d.DzqJyBz,
		d.DzqId,
		d.DzqId, d.DzqSzMc, d.DzqDwXsMc, d.DzqBm, d.DzqPym,
		d.DzqJrqBj, d.DzqZzqMz, d.DzqYt, d.DzqHxhLy, d.DzqDbBs,
		d.DzqHsXz, d.DzqCtqId, d.DzqZdZkl, d.DzqQxZkYh, d.DzqQhsJzr,
		d.DzqJyBz)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func (r *repZxZc) DelZ3DzqSza(id int64) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelZ3DzqSza, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//货品品牌设置
func (r *repZxZc) UpdateZ3Ppa(d *object.Z3Ppa) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZ3Ppa,
		d.PpId,
		d.PpId, d.PpSzMc, d.PpQc, d.PpYt, d.PpBm,
		d.PpSym, d.PpPym, d.PpIsForbidden, d.PpBz,
		d.PpId,
		d.PpId, d.PpSzMc, d.PpQc, d.PpYt, d.PpBm,
		d.PpSym, d.PpPym, d.PpIsForbidden, d.PpBz)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
func (r *repZxZc) DelZ3Ppa(id int64) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelZ3Ppa, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

//==============================================================================================

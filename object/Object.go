package object

import (
	"time"
)

//==============================================================================================

//销售出库
type Z3XsCkDt struct {
	CkdMxHh  string
	CkDjh    string
	CkdLsh   string
	CkdBj    int
	CkdCxBj  int
	CkdYyr   time.Time
	CkdMdId  int
	CkdCkId  int
	CkdHpId  int
	CkdDwId  int
	CkdHsl   float64
	CkdJmSl  float64
	CkdBqjXj float64
	CkdCjjXj float64
	CkdFjXx  string
	CkZdRid  int
	CkKhId   int
	CkBz     string
	CkZdzSj  time.Time
}

//销售退货
type Z3XsTht struct {
	ThHpMxHh   string
	ThDjh      string
	ThLsh      string
	ThYyRq     time.Time
	ThMdid     int
	ThQtCkId   int
	ThHpId     int
	ThJlDwId   int
	ThZhl      float64
	ThJmSl     float64
	ThBqjJeXj  float64
	ThZzCjJeXj float64
	ThHpFjXx   string
	ThZdrId    int
	ThKhId     int
	ThBz       string
	ThZdzSj    time.Time
}

//调拨出库
type Z3MdDbCkDt struct {
	DbdMxHh    string
	DbdDjh     string
	DbdLsh     string
	DbdCkr     time.Time
	DbdDcJgId  int
	DbdDcCkId  int
	DbrKShJgId int
	DbrKPpId   int
	DbdHpId    int
	DbdDwId    int
	DbdHsl     float64
	DbdJmSl    float64
	DbdLsj     float64
	DbdDhj     float64
	DbdZdrId   int
	DbdBz      string
	DbdShSj    time.Time
}

//调拨调整
type Z3MdDbTzDt struct {
	TzdMxHh   string
	TzdDjh    string
	TzdLsh    string
	TzdCkr    time.Time
	TzdJgId   int
	TzdCkId   int
	TzRkJgId  int
	TzRkCkId  int
	TzdHpId   int
	TzdDwId   int
	TzdHsl    float64
	TzdJmTzSl float64
	TzdLsj    float64
	TzdDhj    float64
	TzdZdrId  int
	TzdBz     string
	TzdShSj   time.Time
}

//货品出库登记
type Z3HpCkDjDt struct {
	CkdMxHh     string
	CkdDjh      string
	CkdLsh      string
	CkdYyr      time.Time
	CkdCkFzJgId int
	CkdCkId     int
	CkdHpId     int
	CkdDwId     int
	CkdHsl      float64
	CkdJmSl     float64
	CkdLsj      float64
	CkdDhj      float64
	CkBz        string
	CkShrId     int
	CkShSj      time.Time
}

//货品入库登记
type Z3HpRkDjDt struct {
	RkdMxHh     string
	RkdDjh      string
	RkdLsh      string
	RkdYyr      time.Time
	RkdRkFzJgId int
	RkdCkId     int
	RkdHpId     int
	RkdDwId     int
	RkdHsl      float64
	RkdJmSl     float64
	RkdDhj      float64
	RkdLsj      float64
	RkBz        string
	RkShrId     int
	RksShSj     time.Time
}

//盘亏登记
type Z3PkDjDt struct {
	PkdMxHh   string
	PkdDjh    string
	PkdLsh    string
	PkdYyr    time.Time
	PkdPdJgId int
	PkdPdCkId int
	PkdHpId   int
	PkdDwId   int
	PkdHsl    float64
	PkdJmSl   float64
	PkdLsj    float64
	PkdDhj    float64
	PkdBz     string
	PkdShrId  int
	PkdShSj   time.Time
}

//门店配送调整
type Z3MdPsTzDt struct {
	TzdMxHh   string
	TzdDjh    string
	TzdLsh    string
	TzdCkr    time.Time
	TzdShJgId int
	TzdShCkId int
	TzdPsJgId int
	TzdPsCkId int
	TzdHpId   int
	TzdDwId   int
	TzdHsl    float64
	TzdJmSl   float64
	TzdDhj    float64
	TzdPsj    float64
	TzdBz     string
	TzdShrId  int
	TzdShSj   time.Time
}

//提货货品明细
type Z3XsDdThDt struct {
	ThdMxHh   string
	ThdDjh    string
	ThdDdLsh  string
	ThdLsh    string
	ThdDjLx   int
	ThdYyr    time.Time
	ThdJfMdId int
	ThdCkId   int
	ThdKhPpid int
	ThPhBj    int
	ThKhId    int
	ThdHpId   int
	ThdDwId   int
	ThdHsl    float64
	ThdDZxSl  float64
	ThDdBqjXj float64
	ThDdCjjXj float64
	ThdJmSl   float64
	ThdBqjXj  float64
	ThdCjjXj  float64
	ThZdrId   int
	ThZdSj    time.Time
	ThBz      string
}

//==============================================================================================

//配送出库货品明细
type Z3PsCkHpDt struct {
	CkdLsh       string
	CkdHpId      int
	CkdDjh       string
	CkdDwId      int
	CkdHsl       float64
	CkdCckFzJgId int
	CkdCckId     int
	CkMdId       int
	CkPpId       int
	CkdEdDjLx    int
	CkdPhBj      int
	CkdDhZxSl    float64
	CkdZxDhj     float64
	CkdDhJe      float64
	CkdZsPsj     float64
	CkdHrq       time.Time
	CkdShSj      time.Time
	CkdPsJmSl    float64
	CkPsBz       string
	CkFcCzrId    int
}

//==============================================================================================

//集团通用货品设置A
type Z3Hpa struct {
	HpId          int
	HpBm          string
	HpSzMc        string
	HpPpId        int
	HpSym         string
	HpPym         string
	HpZxDwId      int
	HpZxDwMc      string
	HpEjFlId      int
	HpMdSzBj      int
	HpdDkhDhHpBj  int
	HpWdDhHpBj    int
	HpMs          string
	HpZyNj        int
	HpIsForbidden int
	HpSzId        int
	HpNsFlm       string
	HpJrHp        int
	HpDgHp        int
}

//货品计量单位附加表索引
type Z3HpDwFjSyt struct {
	DwFjHpId int
	DwFjDwId int
}

//货品计量单位附加表
type Z3HpDwFja struct {
	DwFjHpId        int
	DwFjDwId        int
	DwFjDwLx        int
	DwFjHsl         float64
	DwFjTxm         string
	DwFjIsForbidden int
}

//计量单位设置（A）
type Z3JlDwa struct {
	DwId          int
	DwSzMc        string
	DwGlMc        string
	DwSym         string
	DwPym         string
	DwIsForbidden int
}

//货品二级分类
type Z3HpEjFlt struct {
	EjFlId       int
	EjFlMc       string
	EjFlBm       string
	EjFlPpId     int
	EjFlSym      string
	EjFlPym      string
	EjFlHpYjFlId int
}

//==============================================================================================

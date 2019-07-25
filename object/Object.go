package object

import (
	"time"
)

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

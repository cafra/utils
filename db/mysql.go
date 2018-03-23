package db

import (
	"errors"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var NoData = errors.New("no data")

type MsqlExtraCfg struct {
	ShowSQL      bool
	MaxOpenConns int
	MaxIdleConns int
}

type MysqlDao struct {
	debug  bool
	engine *xorm.Engine
}

func NewMysqlDao(cfgStr string, extraCfg *MsqlExtraCfg) (dao *MysqlDao, err error) {
	dao = new(MysqlDao)
	dao.engine, err = xorm.NewEngine("mysql", cfgStr)
	if err != nil {
		return
	}
	dao.engine.ShowSQL(extraCfg.ShowSQL)
	dao.engine.SetMaxOpenConns(extraCfg.MaxOpenConns)
	dao.engine.SetMaxIdleConns(extraCfg.MaxIdleConns)

	return
}

func (d *MysqlDao) Insert(data interface{}) (id int64, err error) {
	id, err = d.engine.InsertOne(data)
	return
}
func (d *MysqlDao) InsertAll(data interface{}) (id int64, err error) {
	id, err = d.engine.Insert(data)
	return
}

func (d *MysqlDao) Get(data interface{}) (bool, error) {
	return d.engine.Get(data)
}

func (d *MysqlDao) GetById(id int64, data interface{}) (err error) {
	ok, err := d.engine.Id(id).Get(data)
	if !ok {
		err = NoData
	}
	return
}
func (d *MysqlDao) GetById_I(id interface{}, data interface{}) (err error) {
	ok, err := d.engine.Id(id).Get(data)
	if !ok {
		err = NoData
	}
	return
}

func (d *MysqlDao) Query(sql string, params interface{}) ([]map[string][]byte, error) {
	return d.engine.Query(sql, params)
}

func (d *MysqlDao) GetByWhere(data interface{}, where string, args ...interface{}) (err error) {
	ok, err := d.engine.Where(where, args...).Get(data)
	if !ok {
		err = NoData
	}
	return
}

func (d *MysqlDao) GetList(l interface{}) (err error) {
	err = d.engine.Find(l)
	if err != nil {
		return
	}
	return
}
func (d *MysqlDao) GetListDesc(l interface{}) (err error) {
	err = d.engine.Desc("id").Find(l)
	if err != nil {
		return
	}
	return
}
func (d *MysqlDao) GetListOrderBy(l interface{}, IsDes bool, key ...string) (err error) {
	if IsDes {
		err = d.engine.Desc(key...).Find(l)
	} else {
		err = d.engine.Asc(key...).Find(l)
	}
	if err != nil {
		return
	}
	return
}

func (d *MysqlDao) GetListWhere(l interface{}, where string, args ...interface{}) (err error) {
	err = d.engine.Where(where, args...).Find(l)
	if err != nil {
		return
	}

	return
}

func (d *MysqlDao) GetCount(l interface{}) (total int64, err error) {
	total, err = d.engine.Count(l)
	if err != nil {
		return
	}
	return
}

func (d *MysqlDao) Count(sql string, data interface{}, args ...interface{}) (int64, error) {
	return d.engine.SQL(sql, args...).Count(data)
}

func (d *MysqlDao) GetCountWhere(l interface{}, where string, args ...interface{}) (total int64, err error) {
	total, err = d.engine.Where(where, args...).Count(l)
	if err != nil {
		return
	}
	return
}

func (d *MysqlDao) GetSumWhere(l interface{}, cols string, where string, args ...interface{}) (total float64, err error) {
	total, err = d.engine.Where(where, args...).Sum(l, cols)
	if err != nil {
		return
	}
	return
}

//获取最新一条数据
func (d *MysqlDao) GetLastOneById(l interface{}) (exist bool, err error) {
	exist, err = d.engine.Desc("id").Limit(1, 0).Get(l)
	if err != nil {
		return
	}
	return
}

func (d *MysqlDao) Update(data interface{}, where string, args ...interface{}) (err error) {
	_, err = d.engine.Where(where, args...).Update(data)
	return
}

func (d *MysqlDao) UpdateById(id int64, data interface{}) (err error) {
	_, err = d.engine.Id(id).Update(data)
	return
}

func (d *MysqlDao) UpdateByCols(id int64, data interface{}, Cols ...string) (err error) {
	_, err = d.engine.Id(id).Cols(Cols...).Update(data)
	return
}

func (d *MysqlDao) UpdateByCond(been interface{}, cond interface{}) (int64, error) {
	id, err := d.engine.Update(been, cond)
	if err != nil {
		return id, err
	}
	if id == 0 {
		return id, NoData
	}
	return id, err
}

func (d *MysqlDao) Delete(data interface{}) (id int64, err error) {
	id, err = d.engine.Delete(data)
	return
}

//GetAll 查询所有的数据
func (d *MysqlDao) GetAll(res interface{}) error {
	return d.engine.Find(res)
}

//GetByCond 根据条件查询数据
func (d *MysqlDao) GetByCond(res interface{}, cond interface{}) error {
	return d.engine.Find(res, cond)
}

func (d *MysqlDao) Engine() *xorm.Engine {
	return d.engine
}

package db

import (
	//"fmt"
	"testing"
	//"utils"
)

var (
	md  *MysqlDao
	err error
)

func init() {
	md, err = NewMysqlDao("root:111111@tcp(localhost:3306)/cm_yn_loan?timeout=3s&parseTime=true&loc=Local&charset=utf8",
		&MsqlExtraCfg{
			ShowSQL:      true,
			MaxIdleConns: 5,
			MaxOpenConns: 10,
		})
	if err != nil {
		panic(err)
	}
}

func test() error {
	//u := new(model.UserInfo)
	//return md.GetById(1, u)
	return nil
}
func test2() error {
	//u := new(model.UserInfo)
	//err := md.GetById(1, u)
	//fmt.Printf("%#v", u)
	//err = md.UpdateById(1, model.UserInfo{Name: "czcz"})
	//md.GetById(1, u)
	//fmt.Printf("%#v", u)

	return err
}

func TestMysqlDao_Get(t *testing.T) {
	t.Log(test2())
}

func BenchmarkMysqlDao_GetById(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := test(); err != nil {
			b.Log(err)
			b.Fail()
		}
	}
}

type UserInfo struct {
	Id    int64  `json:"-"`
	Name  string `xorm:"VARCHAR(40)" json:"name"`
	Phone int64  `xorm:"BIGINT(20) unique" json:"phone,string"`
}

func TestInsert(t *testing.T) {
	u := &UserInfo{Phone: 13683566670}
	t.Log(md.Engine().InsertOne(u))
	t.Log(u)
}

type AuthStatus struct {
	Site  string `xorm:"site"`
	Value int    `xorm:"value"`
}

func TestGroup(t *testing.T) {
	ats := make([]*AuthStatus, 0)
	t.Log(md.Engine().SQL("select site,count(1) value from auth where uid=1 group by site ").Find(&ats))

	t.Log(len(ats))
	for _, info := range ats {
		t.Log(info.Site, info.Value)
	}

}

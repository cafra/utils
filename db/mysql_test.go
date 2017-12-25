package db

import "testing"

func TestNewMysqlDao(t *testing.T) {
	md, err := NewMysqlDao("root:111111@tcp(localhost:3306)/mysql?timeout=3s&parseTime=true&loc=Local&charset=utf8",
		&MsqlExtraCfg{
			ShowSQL:      true,
			MaxIdleConns: 100,
			MaxOpenConns: 100,
		})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(md.Engine().Exec("select 1"))
}

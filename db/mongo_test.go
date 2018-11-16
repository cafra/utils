package db
//
//import (
//	"github.com/cafra/utils"
//	"github.com/globalsign/mgo/bson"
//	"github.com/spf13/viper"
//	"log"
//	"testing"
//	"time"
//)
//
//var (
//	dao MongoDB
//)
//
//func init() {
//	viper.Reset()
//	viper.SetConfigFile("./cfg.yml")
//	err := viper.ReadInConfig()
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	viper.WatchConfig()
//
//	dao = MustOpen("uangme_loan")
//}
//
//func TestOpen(t *testing.T) {
//	dao = MustOpen("uangme_loan")
//}
//
//func TestInsert4Mgo(t *testing.T) {
//	utils.CurrentTest(func() error {
//		c := dao.C("test")
//		err = c.Insert(bson.D{{Name: "uid", Value: time.Now().UnixNano()}, {Name: "name", Value: "cz"}})
//		if err != nil {
//			t.Log("=====", err)
//		}
//		return err
//	}, 10000, 100)
//}
//
//func TestCount(t *testing.T) {
//	utils.CurrentTest(func() error {
//		c := dao.C("test")
//		t.Log(c.Find(bson.D{{Name: "uid", Value: 1534909216178647811}}).Count())
//		return err
//	}, 10000, 100)
//}

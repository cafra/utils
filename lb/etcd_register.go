package lb

import (
	"context"
	"fmt"
	"strings"

	"github.com/coreos/etcd/clientv3"
)

var (
	Prefix     = "etcd3_naming"
	deregister = make(chan interface{})
)

// endpoints 服务的地址 host:port
// etcds etcd 集群地址列表，"," 隔开的字符串
func Register(name, endpoints, etcds string, ttl int64) (err error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(etcds, ","),
	})
	if err != nil {
		return err
	}
	//defer client.Close() 坑，deregister 可以让client 逃逸到堆上，从而延长生命周期

	key := fmt.Sprintf("/%v/%v/%v", Prefix, name, endpoints)
	//grant
	lease, err := client.Grant(context.TODO(), ttl)
	if err != nil {
		return err
	}
	//put lease
	if _, err = client.Put(context.TODO(), key, endpoints, clientv3.WithLease(lease.ID)); err != nil {
		return err
	}
	//keep alive heartbeat
	if _, err = client.KeepAlive(context.TODO(), lease.ID); err != nil {
		return
	}
	go func() {
		<-deregister
		// 坑 既然都是lease  则删除的时候应该使用Revoke，不是del,结束时关闭client
		client.Revoke(context.TODO(), lease.ID)
		client.Close()
		deregister <- struct{}{}
	}()
	return nil
}
func Deregister() {
	deregister <- struct{}{}
	<-deregister
}

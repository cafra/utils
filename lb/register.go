package lb

import (
	"strings"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/proxy/grpcproxy"
)

var Prefix = "etcd3_naming"

// endpoints 服务的地址 host:port
// etcds etcd集群地址列表，"," 隔开的字符串
func Register(name, endpoints, etcds string) (c <-chan struct{}, err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(etcds, ","),
	})
	if err != nil {
		return
	}
	prefix := Prefix + name
	//cli.Grant()
	return grpcproxy.Register(cli, prefix, endpoints, 15), nil
}

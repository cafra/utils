package lb

import (
	"errors"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc/naming"
	"strings"
)

type Resolver struct {
	serviceName string
}

func NewResolver(serviceName string) *Resolver {
	return &Resolver{serviceName}
}

func (re *Resolver) Resolve(target string) (naming.Watcher, error) {
	if re.serviceName == "" {
		return nil, errors.New("lb: no service name provided")
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(target, ","),
	})
	if err != nil {
		return nil, err
	}

	return &watcher{re: re, client: client}, nil
}

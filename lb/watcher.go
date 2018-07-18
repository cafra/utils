package lb

import (
	"context"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"google.golang.org/grpc/naming"
	"log"
)

type watcher struct {
	isInitialized bool
	re            *Resolver
	client        *clientv3.Client
}

func (w *watcher) Next() ([]*naming.Update, error) {
	prefix := fmt.Sprintf("/%v/%v/", Prefix, w.re.serviceName)

	if !w.isInitialized {
		log.Print("watcher initing")
		resp, err := w.client.Get(context.Background(), prefix, clientv3.WithPrefix())
		if err != nil {
			return nil, err
		}
		w.isInitialized = true
		//	初始化
		addrs := extractAddrs(resp)
		updates := make([]*naming.Update, 0, len(addrs))

		for _, addr := range addrs {
			updates = append(updates, &naming.Update{Op: naming.Add, Addr: addr})
		}
		return updates, nil
	}

	for wresp := range w.client.Watch(context.Background(), prefix, clientv3.WithPrefix()) {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				return []*naming.Update{{Op: naming.Add, Addr: string(ev.Kv.Value)}}, nil
			case mvccpb.DELETE:
				return []*naming.Update{{Op: naming.Delete, Addr: string(ev.Kv.Value)}}, nil
			}
		}
	}

	return nil, nil
}

// Close closes the Watcher.
func (w *watcher) Close() {

}

func extractAddrs(resp *clientv3.GetResponse) (addrs []string) {
	if resp == nil || resp.Kvs == nil {
		return nil
	}
	addrs = make([]string, resp.Count)
	for _, kv := range resp.Kvs {
		addrs = append(addrs, string(kv.Value))
	}

	return
}

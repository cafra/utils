package lb

import (
	"google.golang.org/grpc/naming"
	"log"
)

type pseudoWatcher struct {
	addrs []string
}

func (w *pseudoWatcher) Next() ([]*naming.Update, error) {
	updates := make([]*naming.Update, 0)
	for _, addr := range w.addrs {
		updates = append(updates, &naming.Update{Op: naming.Add, Addr: addr})
	}
	log.Print("pseudoWatcher Next")
	return updates, nil
}

func (w *pseudoWatcher) Close() {

}

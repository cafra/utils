package lb

import (
	"errors"

	"google.golang.org/grpc/naming"
)

type PseudoResolver struct {
	addrs []string
}

func NewPseudoResolver(addrs []string) *PseudoResolver {
	return &PseudoResolver{addrs}
}

func (re *PseudoResolver) Resolve(target string) (naming.Watcher, error) {
	if len(re.addrs) == 0 {
		return nil, errors.New("lb: no addrs provided")
	}
	w := &pseudoWatcher{
		updatesChan: make(chan []*naming.Update, 1),
	}

	updates := make([]*naming.Update, 0, len(re.addrs))
	for _, addr := range re.addrs {
		updates = append(updates, &naming.Update{Op: naming.Add, Addr: addr})
	}
	w.updatesChan <- updates
	return w, nil
}

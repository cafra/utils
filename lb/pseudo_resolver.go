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

	return &pseudoWatcher{re.addrs}, nil
}

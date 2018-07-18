package lb

import (
	"errors"
	"google.golang.org/grpc/naming"
)

type pseudoWatcher struct {
	updatesChan chan []*naming.Update
}

func (w *pseudoWatcher) Next() ([]*naming.Update, error) {
	uc, ok := <-w.updatesChan
	if !ok {
		return nil, errors.New("updatesChan closed")
	}
	return uc, nil
}

func (w *pseudoWatcher) Close() {
	close(w.updatesChan)
}

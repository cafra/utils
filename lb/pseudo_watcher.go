package lb

import (
	"errors"
	"google.golang.org/grpc/naming"
	"log"
)

type pseudoWatcher struct {
	updatesChan chan []*naming.Update
}

func (w *pseudoWatcher) Next() ([]*naming.Update, error) {
	uc, ok := <-w.updatesChan
	if !ok {
		log.Print("pseudoWatcher Next !ok")
		return nil, errors.New("updatesChan closed")
	}
	log.Print("pseudoWatcher Next")
	return uc, nil
}

func (w *pseudoWatcher) Close() {
	close(w.updatesChan)
}

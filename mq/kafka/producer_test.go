package kafka

import (
	"testing"
)

func TestNewProducer(t *testing.T) {
	p, err := NewProducer("10.60.81.181:9092")
	if err != nil {
		t.Error(err)
		return
	}
	p.Write("test", "hello")

	p.Close()
}

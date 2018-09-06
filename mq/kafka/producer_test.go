package kafka

import (
	"testing"
)

func TestNewProducer(t *testing.T) {
	p, err := NewProducer(KAddrs)
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < 10; i++ {
		t.Log(p.Write(KTopic, "hello"))
		t.Log(p.Write(KTopic, struct {
			Name string
			Age  int
		}{
			Name: "cz",
			Age:  i,
		}))
	}

	p.Close()
}

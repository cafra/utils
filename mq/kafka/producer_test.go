package kafka

import (
	"fmt"
	"testing"
	"time"
)

func TestNewProducer(t *testing.T) {
	p, err := NewProducer(KAddrs)
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < 1; i++ {
		t.Log(p.Write(KTopic, fmt.Sprint(i)))
		//t.Log(p.Write(KTopic, struct {
		//	Name string
		//	Age  int
		//}{
		//	Name: "cz",
		//	Age:  i,
		//}))
	}
	time.Sleep(time.Second * 10)
	p.Close()
}

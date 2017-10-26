package utils

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	//Start(logger)
	Start(logger, 10000, 10000)
}
func logger() error {
	fmt.Sprintf("test")
	return nil
}

package example

import (
	"github.com/go-liam/cache/service"
	"testing"
)

func TestHealth(t *testing.T) {
	v1 := service.Health()
	println("v1=",v1)
}


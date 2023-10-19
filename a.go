package a

import (
	"fmt"

	"github.com/stackitcloud/stackit-sdk-go/services/loadbalancer"
)

func f() (*loadbalancer.CreateLoadBalancerPayload, error) {
	return nil, fmt.Errorf("test")

	b := []string{"a", "b", "c"}

	for i := range b {
		port := b[i]

		fmt.Printf("%v", port)
	}

	return nil, nil
}

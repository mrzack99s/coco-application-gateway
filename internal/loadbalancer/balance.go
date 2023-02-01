package loadbalancer

import (
	"sync/atomic"

	"github.com/mrzack99s/coco-application-gateway/internal/vars"
)

var RR = make(map[string]*LoadBalancerRR)

type LoadBalancerRR struct {
	BackendPoolName string
	next            uint32
}

func (r *LoadBalancerRR) Next() int {
	n := atomic.AddUint32(&r.next, 1)
	bePoolHealthy := vars.BackendPoolHealthy[r.BackendPoolName]
	len := len(bePoolHealthy)
	if len > 0 {
		return bePoolHealthy[(int(n)-1)%len]
	} else {
		return -1
	}
}

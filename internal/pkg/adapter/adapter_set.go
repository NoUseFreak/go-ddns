package adapter

import (
	"fmt"

	"github.com/NoUseFreak/go-ddns/internal/pkg/adapter/route53"
	"github.com/NoUseFreak/go-ddns/internal/pkg/config"
)

type AdapterSet struct {
}

func (as *AdapterSet) GetAdapter(config config.ConfigSet) (Adapter, error) {
	switch config.GetType() {
	case route53.TYPE:
		return &route53.Route53Adapter{}, nil
	}

	return nil, fmt.Errorf("Unable to find an adapter for %s", config.GetType())
}

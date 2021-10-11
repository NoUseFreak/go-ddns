package adapter

import "github.com/nousefreak/go-ddns/internal/pkg/config"

type Adapter interface {
	SetIP(ip string, config *config.ConfigSet) error
}

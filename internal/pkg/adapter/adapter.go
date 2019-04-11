package adapter

import "github.com/NoUseFreak/go-ddns/internal/pkg/config"

type Adapter interface {
	SetIP(ip string, config *config.ConfigSet) error
}

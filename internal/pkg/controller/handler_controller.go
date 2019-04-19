package controller

import (
	"sync"
	"time"

	"github.com/NoUseFreak/go-ddns/internal/pkg/adapter"
	"github.com/NoUseFreak/go-ddns/internal/pkg/config"
	"github.com/NoUseFreak/go-ddns/internal/pkg/utils/dns"
	log "github.com/sirupsen/logrus"
)

type HandlerController struct {
}

func (c *HandlerController) Handle(cf config.Config) {
	adaptorSet := adapter.AdapterSet{}

	var wg sync.WaitGroup
	wg.Add(len(cf.GetSets()))
	for _, set := range cf.GetSets() {
		ad, err := adaptorSet.GetAdapter(set)
		if err != nil {
			log.Errorf(err.Error())
			return
		}
		go func(cs config.ConfigSet, ad adapter.Adapter) {
			defer wg.Done()

			for {
				c.HandleSet(cs, ad)
				time.Sleep(time.Duration(cs.GetTimeout()) * time.Second)
			}
		}(set, ad)
	}
	wg.Wait()

}

func (c *HandlerController) HandleSet(cs config.ConfigSet, adapter adapter.Adapter) {
	util := dns.DnsUtil{}
	domain := cs.GetRecord()

	ip, err := util.GetMyIP()
	if err != nil {
		log.WithField("message", err.Error()).Warn("Failed to get remote ip")
		return
	}
	if update, _ := util.NeedsUpdate(ip, domain); update == true {
		log.Infof("Updating domain '%s'", domain)
		adapter.SetIP(ip, &cs)
		log.Infof("Updating domain complete '%s'", domain)
	} else {
		log.Debugf("Domain '%s' already at correct ip", domain)
	}
}

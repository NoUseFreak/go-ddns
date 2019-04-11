package config

type ConfigSet struct {
	setType string
	record  string
	timeout int
	options map[string]string
}

func (cs *ConfigSet) GetOptions() map[string]string {
	return cs.options
}

func (cs *ConfigSet) GetRecord() string {
	return cs.record
}

func (cs *ConfigSet) GetTimeout() int {
	return cs.timeout
}

func (cs *ConfigSet) GetType() string {
	return cs.setType
}

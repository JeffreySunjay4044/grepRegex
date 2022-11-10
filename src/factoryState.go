package main

type FactoryState struct {
	Dev Dev `mapstructure:"dev"`
}

type Dev struct {
	DefaultFactoryAddress string             `mapstructure:"defaultFactoryAddress"`
	DefaultFactoryData    DefaultFactoryData `mapstructure:"defaultFactoryData"`
	Proxies               []Proxies          `mapstructure:"proxies"`
}

type DefaultFactoryData struct {
	Address string `mapstructure:"address"`
	Owner   string `mapstructure:"owner"`
	Gas     Gas    `mapstructure:"gas"`
}

type Gas struct {
	Type string `mapstructure:"type"`
	Hex  string `mapstructure:"hex"`
}

type Proxies struct {
	FactoryAddress string `mapstructure:"factoryAddress"`
	LogicName      string `mapstructure:"logicName"`
	LogicAddress   string `mapstructure:"logicAddress"`
	Salt           string `mapstructure:"salt"`
	ProxyAddress   string `mapstructure:"proxyAddress"`
	InitCallData   string `mapstructure:"initCallData"`
	Gas            Gas    `mapstructure:"gas"`
}

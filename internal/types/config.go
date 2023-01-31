package types

type ConfigType struct {
	Properties struct {
		SSL struct {
			Enable   bool   `yaml:"enable"`
			KeyPath  string `yaml:"keyPath"`
			CertPath string `yaml:"certPath"`
		} `yaml:"ssl"`
		Port struct {
			Http  int `yaml:"http"`
			Https int `yaml:"https"`
		} `yaml:"port"`
	} `yaml:"properties"`
}

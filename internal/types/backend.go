package types

type Backend struct {
	Name      string `yaml:"name"`
	Hostname  string `yaml:"hostname"`
	IsHealthy bool
}

type BackendPool struct {
	Name               string    `yaml:"name"`
	HealthProbeName    string    `yaml:"healthProbeName"`
	Servers            []Backend `yaml:"servers"`
	ServerHealthyIndex []int
}

type BackendPoolConfig struct {
	Pools []BackendPool `yaml:"pools"`
}

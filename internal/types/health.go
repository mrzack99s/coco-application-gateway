package types

type HealthProbe struct {
	Name       string `yaml:"name"`
	Path       string `yaml:"path"`
	Https      bool   `yaml:"https"`
	StatusCode int    `yaml:"statusCode"`
}

type HealthProbeConfig struct {
	Probes []HealthProbe `yaml:"probes"`
}

package types

type RuleEndpointType struct {
	PoolName         string `yaml:"poolName"`
	LoadBalancerMode string `yaml:"loadBalancerMode"`
	Https            bool   `yaml:"https"`
}

type RuleType struct {
	Hostname string `yaml:"hostname"`
	Features struct {
		WAFEnable   bool          `yaml:"wafEnable"`
		RateLimit   RateLimit     `yaml:"rateLimit"`
		IPWhiteList []IPWhiteList `yaml:"ipWhiteList"`
	} `yaml:"features"`
	Backend RuleEndpointType `yaml:"backend"`
}

type RoutingTypeConfig struct {
	HTTP  []RuleType `yaml:"http"`
	HTTPS []RuleType `yaml:"https"`
}

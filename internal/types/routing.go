package types

type RouteEndpointType struct {
	Path            string `yaml:"path"`
	Action          string `yaml:"action"`
	BackendPoolName string `yaml:"backendPoolName"`
	To              string `yaml:"to"`
	Https           bool   `yaml:"https"`
	Response        struct {
		StatusCode int    `yaml:"code"`
		Message    string `yaml:"message"`
	} `yaml:"response"`
}

type RouteType struct {
	Hostname string `yaml:"hostname"`
	Features struct {
		WAFEnable   bool          `yaml:"wafEnable"`
		RateLimit   RateLimit     `yaml:"rateLimit"`
		IPWhiteList []IPWhiteList `yaml:"ipWhiteList"`
	} `yaml:"features"`
	Routes []RouteEndpointType `yaml:"routes"`
}

type RoutingTypeConfig struct {
	HTTP  []RouteType `yaml:"http"`
	HTTPS []RouteType `yaml:"https"`
}

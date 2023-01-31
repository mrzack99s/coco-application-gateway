package types

type RouteEndpointType struct {
	Path     string `yaml:"path"`
	Action   string `yaml:"action"`
	To       string `yaml:"to"`
	Response struct {
		StatusCode int    `yaml:"code"`
		Message    string `yaml:"message"`
	} `yaml:"response"`
}

type RoutingType struct {
	HTTP []struct {
		Hostname string              `yaml:"hostname"`
		Routes   []RouteEndpointType `yaml:"routes"`
	} `yaml:"http"`
	HTTPS []struct {
		Hostname string              `yaml:"hostname"`
		Routes   []RouteEndpointType `yaml:"routes"`
	} `yaml:"https"`
}

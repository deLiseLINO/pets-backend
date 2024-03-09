package config

import "fmt"

// HTTPServer holds the configuration for the HTTP server.
type HTTPServer struct {
	Host string `yaml:"host"` // Host specifies the server hostname or IP address.
	Port int    `yaml:"port"` // Port specifies the server port to listen on.

	// ShutdownTimeout specifies allowed time to graceful shutdown
	ShutdownTimeout int `yaml:"shutdown_timeout"`

	// ReadTimeout specifies the maximum duration for reading the entire request, including the body
	ReadTimeout int `yaml:"read_timeout"`

	// WriteTimeout specifies the maximum duration before timing out writes of the response
	WriteTimeout int `yaml:"write_timeout"`

	// Release mode for Gin application
	Release bool `yaml:"release"`
}

func (s HTTPServer) ListenAddress() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

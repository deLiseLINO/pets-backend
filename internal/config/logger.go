package config

const (
	OutputStdout = "stdout"
	OutputStderr = "stderr"
)

// Logger holds the configuration for the logger.
type Logger struct {
	Output string `yaml:"output"` // Output specifies the logger output path or type.
	Level  int    `yaml:"level"`  // Level specifies the logger verbosity level.
}

package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds the global configuration
type Config struct {
	ProjectsRoot string `mapstructure:"projects_root"`
	TicketsRoot  string `mapstructure:"tickets_root"`
	TicketNaming string `mapstructure:"ticket_naming"`
}

// Load initializes and loads the configuration
func Load() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home dir: %w", err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(home, ".yard"))
	viper.AddConfigPath(filepath.Join(home, ".config", "yardmaster"))
	viper.AddConfigPath(".")

	viper.SetDefault("projects_root", filepath.Join(home, ".yard", "projects"))
	viper.SetDefault("tickets_root", filepath.Join(home, ".yard", "tickets"))
	viper.SetDefault("ticket_naming", "{{.ID}}")

	viper.SetEnvPrefix("YARD")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
		// Config file not found is okay, use defaults
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

    // Expand tilde
    cfg.ProjectsRoot = expandPath(cfg.ProjectsRoot, home)
    cfg.TicketsRoot = expandPath(cfg.TicketsRoot, home)

	return &cfg, nil
}

func expandPath(path, home string) string {
    if path == "~" {
        return home
    }
    if len(path) > 1 && path[:2] == "~/" {
        return filepath.Join(home, path[2:])
    }
    return path
}

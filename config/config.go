package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)


// Config is a application configuration structure
type Config struct {
	Database struct {
		Host string `yaml:"host" env:"DB_HOST,
		HOST" env-description:"Server host" env-default:"localhost"`
		Port string `yaml:"port" env:"DB_PORT,
		PORT" env-description:"Server port" env-default:"8080"`
	} `yaml:"database"`
}

// Args command-line parameters
type Args struct {
	ConfigPath string
}

func GetConfig() Config{
	var config Config
	args := ProcessArgs(&config)
	// read configuration from the file and environment variables
	if err := cleanenv.ReadConfig(args.ConfigPath, &config); err != nil {
		fmt.Println(err)
		return Config{}
	}
	return config
}

// ProcessArgs processes and handles CLI arguments
func ProcessArgs(cfg interface{}) Args {
	var a Args

	f := flag.NewFlagSet("DB server", 1)
	f.StringVar(&a.ConfigPath, "c", "config.yml", "./")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])
	return a
}

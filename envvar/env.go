package envVariables

import (
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	isLocal    bool
	dotEnvPath string
}

type Props struct {
	DotEnvPath string
}

func New(props Props) (*Env, error) {
	isLocalEnv := false
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "local" || appEnv == "" {
		dotEnvPath := ".env"
		if props.DotEnvPath != "" {
			dotEnvPath = props.DotEnvPath
		}

		viper.SetConfigFile(dotEnvPath)

		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}

		isLocalEnv = true
	}
	return &Env{
		isLocal: isLocalEnv,
	}, nil
}

func (e *Env) Get(key string) string  {
	if e.isLocal {
		return viper.Get(key).(string)
	}
	return os.Getenv(key)
}
package config

type (
	IConfig interface {
		Callback() string
	}

	AppConfig struct {
	}
)

func Get() IConfig {
	c := &AppConfig{}
	return c
}

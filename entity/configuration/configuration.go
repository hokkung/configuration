package configuration

type Configuration struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func (c Configuration) EntID() string {
	return c.Key
}

func New(key, val string) *Configuration {
	return &Configuration{
		Key: key,
		Val: val,
	}
}

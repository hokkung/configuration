package configuration

import "fmt"

type Configuration struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func (c *Configuration) GetKey() string {
	return fmt.Sprintf("%s:%s", c.Key, c.Val)
}

func NewConfiguration(key, val string) *Configuration {
	return &Configuration{
		Key: key,
		Val: val,
	}
}

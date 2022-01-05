package configs

import (
	"github.com/spf13/viper"
)

var c *viper.Viper

func ConfigureFromJsonFile(configFile string) error {
	v := viper.New()
	v.SetConfigFile(configFile)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	c = v

	return nil
}

func Get(k string) interface{} {
	return c.Get(k)
}

func GetString(k string) string {
	return c.GetString(k)
}

func GetBool(k string) bool {
	return c.GetBool(k)
}

func GetInt(k string) int {
	return c.GetInt(k)
}

func GetInt64(k string) int64 {
	return c.GetInt64(k)
}

func GetFloat64(k string) float64 {
	return c.GetFloat64(k)
}

func GetStringSlice(k string) []string {
	return c.GetStringSlice(k)
}

func GetStringMapString(k string) map[string]string {
	return c.GetStringMapString(k)
}

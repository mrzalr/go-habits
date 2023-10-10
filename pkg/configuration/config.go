package configuration

import "github.com/spf13/viper"

func (c *Configuration) SetConfig(path string) error {
	viper.AddConfigPath(".")
	viper.SetConfigName(path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

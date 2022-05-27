package ConfigManger

import (
	"github.com/spf13/viper"
)

var (
	configObj *Config
	isLoad    bool
)

type Config struct {
	ScreenModel       string
	HSPort            int
	MeetingRewardTime int
}

func Load(path string) error {
	if isLoad {
		return nil
	}

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	if len(path) == 0 {
		path = "."
	}

	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	configObj = &Config{}
	err = viper.Unmarshal(configObj)

	isLoad = true
	return err
}

func GetScreenKey() string {
	return configObj.ScreenModel
}

func GetHSPort() int {
	return configObj.HSPort
}

func GetMeetingRewardTime() int {
	return configObj.MeetingRewardTime
}

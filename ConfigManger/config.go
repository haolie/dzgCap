package ConfigManger

import (
	"context"

	"github.com/spf13/viper"
)

var (
	configObj  *Config
	isLoad     bool
	moduleName = "configManger"
)

type Config struct {
	ScreenModel       string
	HSPort            int
	MeetingRewardTime int
	ConfigPath        string
}

func init() {
}

func LoadHandler(ctx context.Context) (errList []error) {
	err := load("")
	if err != nil {
		errList = append(errList, err)
	}

	return
}

func load(path string) error {
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

func GetConfigCopy() Config {
	return *configObj
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

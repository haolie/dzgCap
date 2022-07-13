package ConfigManger

type SetOption func(opt *Config)

var (
	optionSetList = make([]SetOption, 0, 8)
)

func RegisterOptionSet(setter SetOption) {
	optionSetList = append(optionSetList, setter)
}

func SetHSPort(port int) SetOption {
	return func(opt *Config) {
		opt.HSPort = port
	}
}

func SetRewardDuration(duration int) SetOption {
	return func(opt *Config) {
		opt.MeetingRewardTime = duration
	}
}

func DoSet() {
	for _, setter := range optionSetList {
		setter(configObj)
	}
}

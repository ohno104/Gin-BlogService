package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	//vp.AddConfigPath("configs/")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}

	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	//return &Setting{vp}, nil

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			fmt.Println("watcher.Enevts")
			_ = s.ReloadAllSection()
		})
	}()
}

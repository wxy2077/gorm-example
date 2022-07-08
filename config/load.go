package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type LoadCfg struct {
	reloadFunc func(v interface{})
}

func NewLoad() *LoadCfg {
	return &LoadCfg{}
}

func (l *LoadCfg) LoadCfg(path string, v interface{}) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %s ", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(v); err != nil {
		log.Fatal(fmt.Errorf("unmarshal global conf failed, err:%s ", err))
	}

	go func() {
		// 监控配置文件变化
		viper.WatchConfig()

		viper.OnConfigChange(func(in fsnotify.Event) {
			fmt.Println("global config file has been changed !!!")
			if err := viper.Unmarshal(v); err != nil {
				log.Fatal(fmt.Errorf("unmarshal conf failed, err:%s", err))
			}
			fmt.Println("re-init !!!")
		})
	}()

}

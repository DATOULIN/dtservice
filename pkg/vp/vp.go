package vp

import (
	"github.com/spf13/viper"
)

type Viper struct {
	vp *viper.Viper
}

type Options struct {
	ConfigName string
	ConfigPath string
	ConfigType string
}

func (vp Viper) ReadSection(k string, v interface{}) error {
	err := vp.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

func NewViper(opts *Options) (*Viper, error) {
	vp := viper.New()
	vp.SetConfigName(opts.ConfigName)
	vp.AddConfigPath(opts.ConfigPath)
	vp.SetConfigType(opts.ConfigType)
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Viper{vp}, nil
}

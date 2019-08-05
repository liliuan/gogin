package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	MgoUrl  string `env:"MgoUrl" envDefault:"mongodb://user:password@127.0.0.1:37017/mydb"`
	MgoName string `env:"MgoName" envDefault:"mydb"`
}

var Cfg Config

func init() {

	err := env.Parse(&Cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CONFIG AS: %+v\n", Cfg)
}

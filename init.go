package visimail

import (
	"log"

	"github.com/vrischmann/envconfig"
)

func init() {
	if err := envconfig.InitWithPrefix(&env, configPrefix); err != nil {
		log.Fatal(err)
	}
}

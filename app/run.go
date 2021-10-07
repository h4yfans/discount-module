package app

import (
	"log"

	"github.com/h4yfans/discount-module/rule"
	"github.com/h4yfans/discount-module/service"
)

func Run() {
	e := service.NewEntityService()
	c, err := rule.NewConfig()
	if err != nil {
		log.Fatalf("config could not set: %v", err)
	}

	r := rule.NewRules(c)
	client := service.NewClient(e, r)
	err = client.Output()
	if err != nil {
		log.Fatalf("Something went wrong: err -> %v", err)
	}
}

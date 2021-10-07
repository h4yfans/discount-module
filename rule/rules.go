package rule

import (
	"github.com/h4yfans/discount-module/entity"
)

type Ruler interface {
	ApplyRule(s RuleSet) *entity.Shipment
}

func RuleList(configRules *ConfigRules) []Ruler {
	return []Ruler{
		NewLowestPackagePrice(configRules),
		NewFreeShipmentPerMonth(configRules),
		NewMaximumDiscountPerMonth(configRules),
	}
}

type RuleSet struct {
	ProviderList *entity.ProviderList
	Shipment     *entity.Shipment
	ConfigRules  *ConfigRules
}

type Rules struct {
	Config   *Config
	RuleList []Ruler
}

func NewRules(config *Config) *Rules {
	r := &Rules{Config: config}
	return r.SetRules()
}

func (r *Rules) SetRules() *Rules {
	r.RuleList = RuleList(r.Config.ConfigRules)
	return r
}

func (r *Rules) ApplyAllRules(s RuleSet) *entity.Shipment {
	s.ConfigRules = r.Config.ConfigRules
	shipment := &entity.Shipment{}
	for _, rule := range r.RuleList {
		shipment = rule.ApplyRule(s)
	}
	return shipment
}

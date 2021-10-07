package rule

import (
	"encoding/json"

	"github.com/h4yfans/discount-module/entity"
)

type FreeShipmentPerMonth struct {
	ConfigRules *ConfigRules
	shipmentMap map[string]map[string]map[string]int
}

func NewFreeShipmentPerMonth(r *ConfigRules) *FreeShipmentPerMonth {
	return &FreeShipmentPerMonth{
		ConfigRules: r,
		shipmentMap: make(map[string]map[string]map[string]int)}
}

func (f *FreeShipmentPerMonth) ApplyRule(s RuleSet) *entity.Shipment {
	var data map[string]map[string]int
	byteData, _ := json.Marshal(f.ConfigRules.Rules.FreeShippingPerMonthRule.Providers)
	_ = json.Unmarshal(byteData, &data)

	month := s.Shipment.GetDate().Format("2006-01")
	providerName := s.Shipment.GetProvider().GetName()
	shipmentSize := s.Shipment.GetPackage().GetSize()

	_, ok := f.shipmentMap[month]
	if !ok {
		f.shipmentMap[month] = map[string]map[string]int{}
	}

	_, ok = f.shipmentMap[month][providerName]
	if !ok {
		f.shipmentMap[month][providerName] = map[string]int{}
	}

	f.shipmentMap[month][providerName][shipmentSize] += 1

	if f.shipmentMap[month][providerName][shipmentSize] == data[providerName][shipmentSize] {
		s.Shipment.SetDiscount(s.Shipment.GetPrice() + s.Shipment.GetDiscount())
		s.Shipment.SetPrice(0)
	}

	return s.Shipment
}

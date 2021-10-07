package rule

import (
	"github.com/h4yfans/discount-module/common"
	"github.com/h4yfans/discount-module/entity"
)

type LowestPackagePrice struct {
	ConfigRules *ConfigRules
}

func NewLowestPackagePrice(r *ConfigRules) Ruler {
	return &LowestPackagePrice{ConfigRules: r}
}

func (l *LowestPackagePrice) ApplyRule(s RuleSet) *entity.Shipment {
	_, ok := common.Find(l.ConfigRules.Rules.LowestPriceRule.Packages, s.Shipment.GetPackage().GetSize())
	if ok {
		price := s.Shipment.GetPrice()
		lowestPrice := s.Shipment.GetPrice()
		for _, provider := range *s.ProviderList {
			providerPrice, ok := provider.GetProviderPackageShipmentPrice(s.Shipment.GetPackage().GetSize())
			if ok && providerPrice < lowestPrice {
				lowestPrice = providerPrice
			}
		}

		if price > lowestPrice {
			discount := price - lowestPrice
			s.Shipment.SetPrice(lowestPrice)
			s.Shipment.SetDiscount(discount)
		}
	}
	return s.Shipment
}

package rule

import (
	"github.com/h4yfans/discount-module/entity"
)

type MaximumDiscountPerMonth struct {
	ConfigRules     *ConfigRules
	appliedDiscount map[string]float64
}

func NewMaximumDiscountPerMonth(r *ConfigRules) *MaximumDiscountPerMonth {
	return &MaximumDiscountPerMonth{
		ConfigRules:     r,
		appliedDiscount: make(map[string]float64),
	}
}

func (f *MaximumDiscountPerMonth) ApplyRule(s RuleSet) *entity.Shipment {
	month := s.Shipment.GetDate().Format("2006-01")

	f.appliedDiscount[month] += s.Shipment.GetDiscount()

	if f.appliedDiscount[month] > s.ConfigRules.Rules.MaxDiscountPerMonthRule.MaximumDiscount {
		discount := f.appliedDiscount[month] - s.ConfigRules.Rules.MaxDiscountPerMonthRule.MaximumDiscount

		s.Shipment.SetPrice(s.Shipment.GetPrice() + discount)
		s.Shipment.SetDiscount(s.Shipment.GetDiscount() - discount)

		f.appliedDiscount[month] -= discount
	}

	return s.Shipment
}

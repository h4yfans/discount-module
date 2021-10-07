package entity

import (
	"github.com/h4yfans/discount-module/common"
)

type ProviderPackage struct {
	pkg   *Package
	price float64
}

type ProviderPackageList map[string]*ProviderPackage

func NewProviderPackage(p *Package, price string) (*ProviderPackage, error) {
	fixedPrice, err := common.ToFloat64(price)
	if err != nil {
		return nil, err
	}
	return &ProviderPackage{pkg: p, price: fixedPrice}, nil
}

func (p ProviderPackageList) GetItem(size string) *ProviderPackage {
	providerPackage, ok := p[size]
	if !ok {
		return nil
	}
	return providerPackage
}

func NewProviderPackageList() ProviderPackageList {
	return ProviderPackageList{}
}

func (p ProviderPackage) GetPrice() float64 {
	return p.price
}

package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/h4yfans/discount-module/common"
)

type Shipment struct {
	date     time.Time
	pkg      *Package
	provider *Provider
	price    float64
	discount float64
}

func NewShipment() *Shipment {
	return &Shipment{}
}

func (s *Shipment) FromString(line, delimiter string, providerList *ProviderList) (*Shipment, error) {
	shipmentData := strings.Split(line, delimiter)
	if len(shipmentData) != 3 {
		return nil, nil
	}

	shipmentDate := strings.TrimSpace(shipmentData[0])
	shipmentPackageSize := strings.TrimSpace(shipmentData[1])
	shipmentProvider := strings.TrimSpace(shipmentData[2])

	pkg := NewPackage(shipmentPackageSize)
	if !pkg.IsSizeValid() {
		return nil, fmt.Errorf("invalid package size %s", pkg.GetSize())
	}

	provider := providerList.GetItem(shipmentProvider)

	err := s.setDate(shipmentDate)
	if err != nil {
		return nil, fmt.Errorf("invalid date time. err: %v", err)
	}
	s.setPackage(pkg)
	s.setProvider(provider)

	if provider != nil {
		s.SetPrice(provider.GetProviderPackageList().GetItem(pkg.GetSize()).GetPrice())
	}

	return s, nil
}

func (s *Shipment) setDate(date string) error {
	t, err := time.Parse(common.DateLayout, date)
	if err != nil {
		return err
	}
	s.date = t
	return nil
}

func (s *Shipment) GetDate() time.Time {
	return s.date
}

func (s *Shipment) setPackage(pkg *Package) {
	s.pkg = pkg
}

func (s *Shipment) GetPackage() *Package {
	return s.pkg
}

func (s *Shipment) setProvider(provider *Provider) {
	s.provider = provider
}

func (s *Shipment) GetProvider() *Provider {
	return s.provider
}

func (s *Shipment) SetPrice(price float64) {
	s.price = price
}

func (s *Shipment) GetPrice() float64 {
	return s.price
}

func (s *Shipment) SetDiscount(discount float64) {
	s.discount = discount
}

func (s *Shipment) GetDiscount() float64 {
	return s.discount
}

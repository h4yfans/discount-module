package service

import (
	"os"
	"strings"

	"github.com/h4yfans/discount-module/entity"
)

type EntityListService struct {
	ProviderDataFileLocation string
}

// NewEntityService return a new entity instance
func NewEntityService() *EntityListService {
	e := EntityListService{}
	return e.SetProviderDataFileLocation()
}

// GetProviderList return available provider list with related entities
func (e *EntityListService) GetProviderList() (*entity.ProviderList, error) {
	providerList := entity.NewProviderList()

	fileService := NewFileService(e.ProviderDataFileLocation)
	file, err := fileService.Open()
	if err != nil {
		return nil, err
	}

	for _, v := range file.Lines {
		providerData := strings.Split(v, " ")
		providerName := strings.TrimSpace(providerData[0])
		packageSize := strings.TrimSpace(providerData[1])
		providerShippingPackagePrice := strings.TrimSpace(providerData[2])
		shippingPackage := entity.NewPackage(packageSize)

		providerPackage, err := entity.NewProviderPackage(shippingPackage, providerShippingPackagePrice)
		if err != nil {
			return nil, err
		}

		provider := entity.NewProvider(providerName)

		p := providerList.GetItem(provider.GetName())
		if p == nil {
			providerList.AddItem(provider)
		}

		p = providerList.GetItem(provider.GetName()).SetProviderPackageList(providerPackage)
		providerList.AddItem(p)
	}
	return providerList, nil
}

func (e *EntityListService) SetProviderDataFileLocation() *EntityListService {
	e.ProviderDataFileLocation = os.Getenv("PROVIDERS_PATH")
	return e
}

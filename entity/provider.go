package entity

type Provider struct {
	name        string
	packageList ProviderPackageList
}

func NewProvider(name string) *Provider {
	return &Provider{name: name, packageList: NewProviderPackageList()}
}

func (p *Provider) GetName() string {
	return p.name
}

func (p *Provider) SetProviderPackageList(pkg *ProviderPackage) *Provider {
	p.packageList[pkg.pkg.GetSize()] = pkg
	return p
}

func (p *Provider) GetProviderPackageList() ProviderPackageList {
	return p.packageList
}

func (p *Provider) GetProviderPackageShipmentPrice(size string) (float64, bool) {
	pkg := p.packageList.GetItem(size)
	if pkg == nil {
		return 0, false
	}

	return pkg.GetPrice(), true
}

type ProviderList map[string]Provider

func NewProviderList() *ProviderList {
	return &ProviderList{}
}

func (p ProviderList) GetItem(name string) *Provider {
	provider, ok := p[name]
	if !ok {
		return nil
	}
	return &provider
}

func (p ProviderList) AddItem(provider *Provider) {
	p[provider.name] = *provider
}

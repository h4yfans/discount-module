package entity

const (
	SIZE_S = "S"
	SIZE_M = "M"
	SIZE_L = "L"
)

type Package struct {
	size string
}

func NewPackage(size string) *Package {
	return &Package{size: size}
}

func (p *Package) GetSize() string {
	return p.size
}

func (p *Package) IsSizeValid() bool {
	if p.size == SIZE_S || p.size == SIZE_M || p.size == SIZE_L {
		return true
	}
	return false
}

package application

type ApplicationType string

const (
	TypeAll      ApplicationType = "all"
	TypeOneClick ApplicationType = "one-click"
	TypeMarket   ApplicationType = "market"
)

func (t ApplicationType) IsValid() bool {
	return t == TypeAll || t == TypeOneClick || t == TypeMarket
}

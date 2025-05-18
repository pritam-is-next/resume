package components

type DropDown struct {
	Name string
	Href string
}
type NavItems struct {
	Name       string
	Href       string
	Disabled   bool
	IsDropDown bool
	DropDown   []DropDown
}

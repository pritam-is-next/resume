package components

type DropDown struct {
	Name string
	Href string
}

var NavItems = []struct {
	Name     string
	Href     string
	Disabled bool
	DropDown []DropDown
}{
	{Name: "Home", Href: "#home", Disabled: false},
	{Name: "About Me", Href: "#about-me", Disabled: false},
	{Name: "Skills", Href: "#skills", Disabled: false},
	{Name: "Experience", Href: "#experience", Disabled: false},
	{Name: "Projects", Href: "#projects", Disabled: false},
	{Name: "Contact", Href: "#contact", Disabled: false},
	// {Name: "DropDown", Href: "#", Disabled: false, DropDown: []DropDown{
	// 	{Name: "Dropdown1", Href: "#"},
	// 	{Name: "Dropdown2", Href: "#"},
	// }},
}

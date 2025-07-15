package components

type CallToAction struct {
	Text string
	Href string
}

var Hero = struct {
	Heading       string
	SubHeading    string
	CallToActions []CallToAction
}{
	Heading:    "Pritam Dutta",
	SubHeading: "Engineer at heart, builder by choice.From bootstrapping a Go server to streamlining cloud deployments — I make things that make a difference.",
	CallToActions: []CallToAction{
		{Text: "View Project", Href: "#projects"},
		{Text: "Get in Touch", Href: "#contact-me"},
	},
}

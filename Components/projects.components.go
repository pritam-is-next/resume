package components

var Projects = []struct {
	Title            string
	SmallDescription string
	Link             string
	Image            string
}{
	// 🏆 Most Unique & Impactful
	{
		Title:            "MissedIT (CS:GO Hack)",
		SmallDescription: "C++ game mod with custom UI framework, theming, and 600+ active users at peak. Combines reverse engineering with polished UX design.",
		Link:             "https://github.com/HackerPolice/MissedIT",
		Image:            "/images/missedit.png",
	},
	{
		Title:            "Tortoise PHP Framework",
		SmallDescription: "Lightweight MVC framework (300+ ★) with build-mode optimization. Simplifies web dev with intuitive routing and DB management.",
		Link:             "https://github.com/vrianta/Tortois",
		Image:            "/images/tortoise-framework.png",
	},
	{
		Title:            "Go HTTP Server Handler",
		SmallDescription: "Production-grade server with encrypted sessions, fast data processing, and custom storage. Demonstrates Go's concurrency power.",
		Link:             "https://github.com/vrianta/Server/tree/golang-dev-2.0",
		Image:            "/images/go-server.png",
	},

	// 🔧 Technical Depth
	{
		Title:            "CIL (Compiled Interpreter Language)",
		SmallDescription: "Custom interpreter with lexer/parser implementation for educational purposes. Teaches compiler design fundamentals.",
		Link:             "https://github.com/tutoraddicts/CIL",
		Image:            "/images/cil-interpreter.png",
	},
}

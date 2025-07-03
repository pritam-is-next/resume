package components

var Projects = []struct {
	Title            string
	SmallDescription string
	Link             string
	Image            string
	TechStack        []string // Technologies used
	Category         string   // Project category (e.g., "Worker Management", "Automation")
	Stars            int      // GitHub stars (if applicable)
	IsFeatured       bool     // Whether it's a featured project
}{
	// 🏆 Most Unique & Impactful Projects
	{
		Title:            "Go HTTP Server Handler",
		SmallDescription: "Production-grade server with encrypted sessions, fast data processing, and custom storage. Demonstrates Go's concurrency power.",
		Link:             "https://github.com/vrianta/Server/tree/golang-dev-2.0",
		Image:            "/Static/img/go-http.avif",
		TechStack:        []string{"Go", "Gin", "Redis"},
		Category:         "Backend",
		Stars:            120,
		IsFeatured:       true,
	},
	{
		Title:            "Workers Hub App (Flutter)",
		SmallDescription: "A Flutter-based worker management app for tracking shifts, tasks, and team coordination. Built for efficient workforce organization.",
		Link:             "https://github.com/vrianta/workers_hub_app_flutter",
		Image:            "/Static/img/flutter-wo1.png",
		TechStack:        []string{"Flutter", "Firebase", "Dart", "Provider"},
		Category:         "Mobile App",
		Stars:            45,
		IsFeatured:       true,
	},
	{
		Title:            "MissedIT (CS:GO Hack)",
		SmallDescription: "C++ game mod with custom UI framework, theming, and 600+ active users at peak. Combines reverse engineering with polished UX design.",
		Link:             "https://github.com/HackerPolice/MissedIT",
		Image:            "/Static/img/missedit.gif",
		TechStack:        []string{"C++", "DirectX", "ImGui"},
		Category:         "Game Mod",
		Stars:            320,
		IsFeatured:       false,
	},
	{
		Title:            "Tortoise PHP Framework",
		SmallDescription: "Lightweight MVC framework (300+ ★) with build-mode optimization. Simplifies web dev with intuitive routing and DB management.",
		Link:             "https://github.com/vrianta/Tortois",
		Image:            "/Static/img/php-mvc.webp",
		TechStack:        []string{"PHP", "MySQL", "Composer"},
		Category:         "Web Framework",
		Stars:            310,
		IsFeatured:       true,
	},
}

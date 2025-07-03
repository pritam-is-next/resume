package components

var Skills = []struct {
	Name  string
	Level int // Percentage (0-100)
}{
	// Core Languages
	{Name: "Go (Golang)", Level: 95}, // Demonstrated by Server Handler & production systems
	{Name: "C/C++", Level: 90},       // MissedIT hack & system-level projects
	{Name: "Python", Level: 85},      // Automation scripts & tools
	{Name: "PHP", Level: 88},         // Tortoise framework
	{Name: "Java", Level: 75},        // Enterprise experience

	// Infrastructure/DevOps
	{Name: "Docker", Level: 92},
	{Name: "Kubernetes", Level: 90},
	{Name: "Terraform", Level: 88},
	{Name: "Azure", Level: 85},

	// Web Technologies
	{Name: "HTTP/HTTPS", Level: 93}, // Custom server implementation
	{Name: "REST APIs", Level: 90},
	{Name: "MVC Architecture", Level: 88}, // Tortoise framework

	// Databases
	{Name: "MySQL/MariaDB", Level: 85},
	{Name: "Session Management", Level: 90}, // Go server feature

	// System Programming
	{Name: "Multithreading", Level: 88},
	{Name: "Memory Management", Level: 85},
	{Name: "Performance Optimization", Level: 90},

	// Tools
	{Name: "Git", Level: 95}, // Active GitHub profile
	{Name: "Linux/Unix", Level: 88},
	{Name: "Bash/PowerShell", Level: 85},

	// Specialized Skills
	{Name: "Compiler Design", Level: 80}, // CIL interpreter
	{Name: "Game Modding", Level: 83},    // MissedIT project
	{Name: "CI/CD Pipelines", Level: 90},
}

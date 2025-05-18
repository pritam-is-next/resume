module github.com/pritam-is-next/resume

go 1.24.3

replace (
	github.com/pritam-is-next/resume/Controllers => ./Controllers
	github.com/pritam-is-next/resume/Server => ../Server
)

require github.com/pritam-is-next/resume/Server v0.0.0-00010101000000-000000000000

require golang.org/x/crypto v0.38.0 // indirect

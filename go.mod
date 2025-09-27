module github.com/pritam-is-next/resume

go 1.24.3

replace github.com/vrianta/agai => ../agai

require (
	github.com/go-sql-driver/mysql v1.9.3
	github.com/vrianta/agai v0.2.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	golang.org/x/crypto v0.40.0 // indirect
)

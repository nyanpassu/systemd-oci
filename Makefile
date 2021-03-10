.PHONY: binary

binary:
	go build -o systemd-runc cmd/runc/runc.go
	go build -o systemd-rununit cmd/rununit/rununit.go

.PHONY: binary

binary:
	go build -o eru-systemd-runc cmd/runc/runc.go
	go build -o eru-systemd-rununit cmd/rununit/rununit.go

install:
	cp eru-systemd-runc /usr/local/bin/
	cp eru-systemd-rununit /usr/local/bin/
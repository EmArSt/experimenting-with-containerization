SCRDIR=./src/
BUILDDIR=./bin/

MAIN=$(SCRDIR)main.go

BINARY=$(BUILDDIR)main


build:$(BINARY)

$(BINARY):$(MAIN)
	go build -o $(BINARY) $(MAIN)

run-bash: build
	sudo $(BINARY) run /bin/bash

.ONESHELL:
get_ubuntu:
	docker export $$(docker create ubuntu) -o rootfs.tar.gz
	mkdir rootfs
	cd rootfs
	tar --no-same-owner --no-same-permissions --owner=0 --group=0 -mxf ../rootfs.tar.gz

clean:
	sudo rm -rf $(BINARY) rootfs rootfs.tar.gz
	sudo umount rootfs

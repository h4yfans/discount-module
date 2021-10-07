SHELL = /bin/sh


.PHONY: build
build:
	docker build --rm -t discount-module .

.PHONY: run
run:
	docker run discount-module


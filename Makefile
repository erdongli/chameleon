ts = $(shell date +%s)

docker:
	docker build . -t chameleon:$(ts) -t chameleon:latest

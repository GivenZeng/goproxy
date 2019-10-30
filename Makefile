make:
	go build -o goproxy *.go
run:
	cp -r etc_sample etc
	make
	./goproxy
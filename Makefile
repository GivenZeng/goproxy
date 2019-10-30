make:
	go build -o goproxy *.go
run:
	rm -rf etc
	cp -r etc_sample etc
	make
	./goproxy
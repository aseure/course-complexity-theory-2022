test:
	go test -v ./...

bench:
	go test -v -bench=. ./...

clean:
	rm -rf *.{dot,png,svg}

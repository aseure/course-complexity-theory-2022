.PHONY: maze
maze:
	go test -v ./maze/...

.PHONY: test
test:
	go test -v ./linear/...
	go test -v ./sorting/...
	go test -v ./tree/...

.PHONY: bench
bench:
	go test -v -bench=. ./linear/...

.PHONY: clean
clean:
	rm -rf *.{dot,png,svg}

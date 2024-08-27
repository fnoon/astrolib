#OPT := -ldflags="-s -w"

.PHONY: all
all: astrolib

astrolib: main.go
	go vet $<
	go fmt $<
	go build -o $@ $(OPT) $<

.PHONY: clean
clean:
	go clean

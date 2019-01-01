GOCMD=go
GOBUILD=$(GOCMD) build -a -tags netgo -installsuffix netgo $(LDFLAGS)
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mackerel-plugin-proc-net-arp

.PYHON: all
all: test build

.PYHON: build
build: bin/linux/386/$(BINARY_NAME) bin/linux/amd64/$(BINARY_NAME) bin/darwin/amd64/$(BINARY_NAME)
bin/linux/386/$(BINARY_NAME): main.go $(shell find lib -type f)
	CGO_ENALBED=0 GOOS=linux GOARCH=386 $(GOBUILD) -o $@ -v

bin/linux/amd64/$(BINARY_NAME): main.go $(shell find lib -type f)
	CGO_ENALBED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $@ -v

bin/darwin/amd64/$(BINARY_NAME): main.go $(shell find lib -type f)
	CGO_ENBALED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $@ -v

.PHONY: test
test: main.go $(shell find lib -type f)
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	find bin -name $(BINARY_NAME) -exec rm -f {} \;

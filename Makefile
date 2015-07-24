TARGET_OSARCH="linux/386"

.PHONY: all clean gox deps build install

all: clean deps build

deps:
	go get -d -v ./...
	go get github.com/mitchellh/gox
	go get github.com/ziutek/mymysql/mysql
	go get github.com/ziutek/mymysql/native

gox:
	go get github.com/mitchellh/gox
	gox -build-toolchain -osarch=$(TARGET_OSARCH)

build: deps
	mkdir -p ./build; \
	for i in mackerel-check-*; do \
		gox -osarch=$(TARGET_OSARCH) -output ./build/$$i github.com/hiroakis/mackerel-agent-checks-plugins/$$i; \
	done

install:
	install -o root -g root -m 0755 ./build/* /usr/local/bin/

clean:
	rm -rf ./build



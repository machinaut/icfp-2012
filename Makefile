#!/usr/bin/env make

ICFP = icfp-96682152

all: icfp install lifter src misc
	tar czf $(ICFP).tgz -C $(ICFP) install lifter PACKAGES PACKAGES-TESTING README src

icfp:
	mkdir -p $(ICFP)

install: icfp
	cd install && go build
	cp install/install  $(ICFP)

lifter: icfp
	cd lifter && go build
	cp lifter/lifter $(ICFP)

src: icfp
	mkdir -p $(ICFP)/src
	cp -r install $(ICFP)/src
	cp -r lifter $(ICFP)/src

misc: icfp
	cp PACKAGES $(ICFP)
	cp PACKAGES $(ICFP)/PACKAGES-TESTING
	cp README $(ICFP)

clean:
	cd lifter && go clean
	cd install && go clean
	rm -rf $(ICFP)
	rm -f $(ICFP).tgz

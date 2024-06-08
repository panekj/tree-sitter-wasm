all install uninstall clean:
	$(MAKE) -C wat $@
	$(MAKE) -C wast $@

test:
	$(TS) test
	$(TS) parse examples/* --quiet --time

.PHONY: all install uninstall clean test update

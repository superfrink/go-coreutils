SUBDIRS = arch base64 basename cat date dirname echo env exit expr factor false groups head logname ls md5sum mkdir mv pwd rm rmdir sha1sum sha224sum sha256sum sha384sum sha512sum sleep stat sync tail tee touch true tsort uname uptime wc whoami yes

all:
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir; \
	done

clean:
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir clean; \
	done

check: test

test:
	echo FIXME write tests

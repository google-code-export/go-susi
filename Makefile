BINARIES=run-tests go-susi encrypt decrypt sibridge

all:
	main/makeversion
	go build main/run-tests.go
	go build main/go-susi.go
	go build main/encrypt.go
	go build main/decrypt.go
	go build main/sibridge.go
	strip $(BINARIES)
	ln -snf go-susi gosa-si-server

frankensusi:
	main/frankensusi $$(find . /srv/www -name "gosa-si-server_*_i386.deb" -print)

debug:
	rm -f run-tests
	go build -a -gcflags '-N -l' main/run-tests.go

test: all
	./run-tests --unit --system=./go-susi

clean:
	rm -f $(BINARIES) gosa-si-server go-susi_*.deb
	hg revert --no-backup testdata/ldif

deb: all
	main/makedeb
	
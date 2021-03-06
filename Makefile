all: binaries
	
binaries: dnsmasqmgr dnsmasqmgrd

clean:
	rm -f cmd/dnsmasqmgr/dnsmasqmgr

dnsmasqmgr: vendor
	cd cmd/dnsmasqmgr && go build -v .

dnsmasqmgrd: vendor
	cd cmd/dnsmasqmgrd && go build -v .

vendor:
	dep ensure

proto: pkg/dnsmasqmgr/dnsmasqmgr.pb.go

pkg/dnsmasqmgr/dnsmasqmgr.pb.go: pkg/dnsmasqmgr/dnsmasqmgr.proto
	cd pkg && protoc -I dnsmasqmgr/ dnsmasqmgr/dnsmasqmgr.proto --go_out=plugins=grpc:dnsmasqmgr

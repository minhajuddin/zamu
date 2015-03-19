# ==================================================================
#	Makefile for zamu
#	builds the binary and tests it
#	Author: Khaja Minhajuddin <minhajuddink@gmail.com>
# ==================================================================

#FOO_KEY := 2323

# run is the default target
run:	build errcheck
	./zamu -logtostderr

build:	test errcheck vet fmt
	godep go build -o zamu

test: errcheck
	GOENV=test godep go test zamu/.../

errcheck:
	errcheck zamu/...

#pipe through cat to get rid of the grep -v exit code
vet:
	@godep go tool vet -all=true . 2>&1 | grep -v 'Godeps/_workspace' | cat

fmt:
	@godep go fmt zamu/...

.PHONY: test test-short build vmserver clean

VM_VERSION := $(shell git describe --tags --long --dirty --always)

clean:
	go clean -cache -testcache

build:
	go build ./...

vmserver:
ifndef VMSERVER_PATH
	$(error VMSERVER_PATH is undefined)
endif
	go build -o ./cmd/vmserver/vmserver ./cmd/vmserver
	cp ./cmd/vmserver/vmserver ${VMSERVER_PATH}

test: clean
	go test -count=1 ./...

test-short:
	go test -short -count=1 ./...

build-test-contracts:
	drtpy contract build ./test/contracts/erc20
	drtpy contract build ./test/contracts/counter

	drtpy contract build ./test/contracts/init-correct
	drtpy contract build ./test/contracts/init-simple
	drtpy contract build ./test/contracts/init-wrong
	drtpy contract build ./test/contracts/misc
	drtpy contract build ./test/contracts/signatures
	drtpy contract build ./test/contracts/vmhooks
	drtpy contract build ./test/contracts/breakpoint
	drtpy contract build --no-optimization ./test/contracts/num-with-fp

	drtpy contract build ./test/contracts/exec-same-ctx-simple-parent
	drtpy contract build ./test/contracts/exec-same-ctx-simple-child
	drtpy contract build ./test/contracts/exec-same-ctx-child
	drtpy contract build ./test/contracts/exec-same-ctx-parent
	drtpy contract build ./test/contracts/exec-dest-ctx-parent
	drtpy contract build ./test/contracts/exec-dest-ctx-child
	drtpy contract build ./test/contracts/exec-same-ctx-recursive
	drtpy contract build ./test/contracts/exec-same-ctx-recursive-parent
	drtpy contract build ./test/contracts/exec-same-ctx-recursive-child
	drtpy contract build ./test/contracts/exec-dest-ctx-recursive
	drtpy contract build ./test/contracts/exec-dest-ctx-recursive-parent
	drtpy contract build ./test/contracts/exec-dest-ctx-recursive-child
	drtpy contract build ./test/contracts/async-call-parent
	drtpy contract build ./test/contracts/async-call-child
	drtpy contract build ./test/contracts/exec-same-ctx-builtin
	drtpy contract build ./test/contracts/deployer
	drtpy contract build ./test/contracts/exchange
	drtpy contract build ./test/contracts/promises
	drtpy contract build ./test/contracts/promises-train
	drtpy contract build ./test/contracts/promises-tracking
	drtpy contract build ./test/contracts/timelocks
	drtpy contract build ./test/contracts/async-call-builtin


build-delegation:
ifndef SANDBOX
	$(error SANDBOX variable is undefined)
endif
	rm -rf ${SANDBOX}/drt-rs-sc-delegation
	git clone --depth=1 --branch=master https://github.com/TerraDharitri/drt-rs-sc-delegation.git ${SANDBOX}/drt-rs-sc-delegation
	rm -rf ${SANDBOX}/drt-rs-sc-delegation/.git
	drtpy contract build ${SANDBOX}/drt-rs-sc-delegation
	drtpy contract test --directory="tests" ${SANDBOX}/drt-rs-sc-delegation
	cp ${SANDBOX}/drt-rs-sc-delegation/output/delegation.wasm ./test/delegation/delegation.wasm


build-dns:
ifndef SANDBOX
	$(error SANDBOX variable is undefined)
endif
	rm -rf ${SANDBOX}/drt-rs-sc-dns
	git clone --depth=1 --branch=master https://github.com/TerraDharitri/drt-rs-sc-dns.git ${SANDBOX}/drt-rs-sc-dns
	rm -rf ${SANDBOX}/drt-rs-sc-dns/.git
	drtpy contract build ${SANDBOX}/drt-rs-sc-dns
	drtpy contract test --directory="tests" ${SANDBOX}/drt-rs-sc-dns
	cp ${SANDBOX}/drt-rs-sc-dns/output/dns.wasm ./test/dns/dns.wasm


build-sc-examples:
ifndef SANDBOX
	$(error SANDBOX variable is undefined)
endif
	rm -rf ${SANDBOX}/sc-examples

	drtpy contract new --template=erc20-c --directory ${SANDBOX}/sc-examples erc20-c
	drtpy contract build ${SANDBOX}/sc-examples/erc20-c
	cp ${SANDBOX}/sc-examples/erc20-c/output/wrc20.wasm ./test/erc20/contracts/erc20-c.wasm


build-sc-examples-rs:
ifndef SANDBOX
	$(error SANDBOX variable is undefined)
endif
	rm -rf ${SANDBOX}/sc-examples-rs
	
	drtpy contract new --template=simple-coin --directory ${SANDBOX}/sc-examples-rs simple-coin
	drtpy contract new --template=adder --directory ${SANDBOX}/sc-examples-rs adder
	drtpy contract build ${SANDBOX}/sc-examples-rs/adder
	drtpy contract build ${SANDBOX}/sc-examples-rs/simple-coin
	drtpy contract test ${SANDBOX}/sc-examples-rs/adder
	drtpy contract test ${SANDBOX}/sc-examples-rs/simple-coin
	cp ${SANDBOX}/sc-examples-rs/adder/output/adder.wasm ./test/adder/adder.wasm
	cp ${SANDBOX}/sc-examples-rs/simple-coin/output/simple-coin.wasm ./test/erc20/contracts/simple-coin.wasm

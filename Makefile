
SPIN ?= spin
TG ?= tinygo
LOGDIR ?= ./_scratch/logs

.PHONY: build
build:
	$(TG) build -o favicon.wasm -target=wasi -wasm-abi=generic -gc=leaking main.go
	ls -lah favicon.wasm
	wasm-opt -O favicon.wasm -o favicon.wasm
	ls -lah favicon.wasm

.PHONY: serve
serve: build
serve:
	RUST_LOG=spin=trace $(SPIN) up --file spin.toml -L $(LOGDIR)
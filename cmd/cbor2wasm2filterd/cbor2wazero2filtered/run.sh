#!/bin/sh

export ENV_WASM_MODULE_DIR=./modules.d/out.d

cat sample.d/input.jsonl |
	json2arr2cbor |
	./cbor2wazero2filtered |
	python3 \
		-m uv \
		tool \
		run \
		cbor2 \
		--sequence

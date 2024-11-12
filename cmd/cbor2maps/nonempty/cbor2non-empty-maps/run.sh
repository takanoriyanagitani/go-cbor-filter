#!/bin/sh

cat ./sample.d/input.jsonl |
  json2map2cbor |
  ./cbor2non-empty-maps |
  python3 \
  	-m uv \
	tool \
	run \
	cbor2 \
	--sequence

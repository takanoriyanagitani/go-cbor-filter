#!/bin/zsh

ijson=./sample.d/input.jsonl

bfilter=removed=b:true
ufilter=iheight=u:634
ifilter=id=s:-1
ffilter=dist=f:42.195
sfilter=date=S:2024/11/12

cat "${ijson}" |
	json2map2cbor |
	./cbor2primitive2filtered \
		"${sfilter}" \
		"${bfilter}" |
	python3 \
		-m uv \
		tool \
		run \
		cbor2 \
		--sequence

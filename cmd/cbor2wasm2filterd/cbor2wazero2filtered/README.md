## Example

input json(actual input: CBOR):

```json
["hw",42.195,634,true,false,null]
["HW",0.599,333,false,true,""]
[]
["HW",3.776,42,true,false,null]
```

output json(actual output: CBOR):

```json
["hw", 42.195, 634.0, true, false, null]
["HW", 0.599, 333.0, false, true, ""]
["HW", 3.776, 42.0, true, false, null]
```

command(*):

```sh
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
```

(*) Requirements:

- json2arr2cbor: github.com/takanoriyanagitani/go-json2cbor/tree/main/cmd/json2arr2cbor
- python
- uv(python lib)
- cbor2(python lib)

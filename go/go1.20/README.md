# Go 1.20

## How to run

```sh
source ./utils.sh

compose-up
run-comparable

for i in ./type-convertions/*; do
    echo ">>> run $i"
    GO run $i
done

GO test -v ./type-convertions/05-string-slice-magic

# cleanup
compose-down
```

## Reference

- https://go.dev/doc/go1.20
- https://stackoverflow.com/questions/59209493/how-to-use-unsafe-get-a-byte-slice-from-a-string-without-memory-copy

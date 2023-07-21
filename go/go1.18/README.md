# Go 1.18

## How to run

```sh
source ./utils.sh

compose-up
GO run ./generics
GO run ./fuzz
GO test ./fuzz
GO test -run=FuzzReverse/e386ddb8ccb05b77 ./fuzz
GO test -fuzz=Fuzz ./fuzz

for i in ./generics/limitations/*; do
    echo ">>> run $i"
    GO run $i
done

# cleanup
compose-down
```

## Reference

- https://go.dev/doc/tutorial/generics
- https://go.dev/doc/tutorial/fuzz

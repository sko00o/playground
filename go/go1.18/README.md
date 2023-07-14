# Go 1.18

## How to run

```sh
docker compose up -d --remove-orphans

GO() {
    docker compose exec -it dev go $@
}
GO run ./generics
GO run ./fuzz
GO test ./fuzz
GO test -run=FuzzReverse/e386ddb8ccb05b77 ./fuzz
GO test -fuzz=Fuzz ./fuzz

# cleanup
docker compose down --remove-orphans --volumes
```

## Reference

- https://go.dev/doc/tutorial/generics
- https://go.dev/doc/tutorial/fuzz

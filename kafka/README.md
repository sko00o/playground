# How to connect Kafka in docker?

## Client (or another broker) in a container, on the same host

```sh
cd client-on-same-host && ./run.sh
```

## Client on same machine, not in a container

```sh
cd client-on-same-machine && ./run.sh
```

## Client on another machine (or broker on remote host)

```sh
cd broker-on-remote-host && ./run.sh
```

## Reference

- https://stackoverflow.com/a/51634499

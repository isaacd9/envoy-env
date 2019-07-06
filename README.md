## My Envoy environment

To use, first download Envoy and build according to the [developer
docs](https://github.com/envoyproxy/envoy/blob/master/bazel/README.md):

```
$ bazel run //source/exe:envoy-static -- -c `pwd`/config.yaml
```

in another shell run

```
$ go run http-server.go
```

you should then be able to see content on `localhost:9000`

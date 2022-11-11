# xk6-udp
A k6 extension for sending datagrams for udp protocol like the one for [tcp](https://github.com/NAlexandrov/xk6-tcp)

## Build

To build a `k6` binary with this plugin, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:

  ```shell
  go install github.com/k6io/xk6/cmd/xk6@latest
  ```

2. Build the binary:

  ```shell
  xk6 build master \
    --with github.com/pedrosaraiva/xk6-udp
  ```

3. Execute k6:

  ```shell
  ./k6 run examples/udp.ts
  ```
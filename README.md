# go-lunatic

go-lunatic is an experimental package to write WASI modules in Go
for use by [lunatic].

[lunatic]: https://lunatic.solutions/

## Examples

| Name             |         Go         |       TinyGo       | Notes |
| ---------------- | :----------------: | :----------------: | ----- |
| [hello]          | :heavy_check_mark: | :heavy_check_mark: |       |
| [net]            | :heavy_check_mark: |        n/a         |       |
| [simple-process] |       fails        |       fails        |       |
| [sleep]          | :heavy_check_mark: | :heavy_check_mark: |       |
| [spawn]          |       fails        |       fails        |       |
| [version]        | :heavy_check_mark: | :heavy_check_mark: |       |

[hello]: ./examples/hello/
[net]: ./examples/net/
[simple-process]: ./examples/simple-process/
[sleep]: ./examples/sleep/
[spawn]: ./examples/spawn/
[version]: ./examples/version/

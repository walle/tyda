[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/walle/tyda)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/walle/tyda/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/walle/tyda)](http:/goreportcard.com/report/walle/tyda)

# tyda.go

CLI tool for searching website tyda.se.

Requires the [tyda-api](https://github.com/walle/tyda-api) binary to be installed.

If you have your `$GOPATH` set up, and `$GOPATH/bin` in you path
```
$ go get github.com/walle/tyda-api/...
```

## Installation

```shell
$ go get github.com/walle/tyda.go/...
```

## Usage

```shell
usage: tyda [--simple] [--languages LANGUAGES] QUERY

positional arguments:
  query

options:
  --simple, -s           Simple output
  --languages LANGUAGES, -l LANGUAGES
                         Languages to translate to (en fr de es la nb sv) [default: [en]]
  --help, -h             display this help and exit
```

## Testing

Use the `go test` tool.

```shell
$ go test -cover
```

## Contributing

All contributions are welcome! See [CONTRIBUTING](CONTRIBUTING.md) for more
info.

## License

The code is under the MIT license. See [LICENSE](LICENSE) for more
information.

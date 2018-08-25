# Serve

[![GitHub release](https://img.shields.io/github/release/z0mbie42/serve.svg)](https://github.com/z0mbie42/serve/releases/latest)
[![Build Status](https://travis-ci.org/z0mbie42/serve.svg?branch=master)](https://travis-ci.org/z0mbie42/serve)

The simplest directory serving web server

1. [Installation](#installation)
2. [Usage](#usage)

-------------------

## Installation

### Using go (nightly)

```bash
$ go get -u github.com/z0mbie42/serve
```

### Binary releases

[https://github.com/z0mbie42/serve/releases/latest](https://github.com/z0mbie42/serve/releases/latest)



## Usage

```bash
$ serve
```

```bash
$ serve help
The simplest directory serving web server. See https://github.com/z0mbie42/serve

Usage:
  serve [flags]
  serve [command]

Available Commands:
  help        Help about any command
  version     Display the version and build information

Flags:
  -a, --address string     address to listen to (default "0.0.0.0")
  -d, --directory string   directory to serve (default ".")
  -f, --format string      logging format. Valid values are: [console, json] (default "console")
  -h, --help               help for serve
  -p, --port string        port to listen to (default "8080")

Use "serve [command] --help" for more information about a command.
```

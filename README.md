
# gocos

Commandline utilities for interacting with Tencent COS service.

## Installation

```bash
go get github.com/evshiron/gocos
```

## Usage

```bash
# setup environment variables
export COS_SECRET_ID=
export COS_SECRET_KEY=
export COS_BUCKET=
export COS_REGION=

# put a file
~/go/bin/gocos put path/to/local/file path/to/remote/file
```

## License

Apache License

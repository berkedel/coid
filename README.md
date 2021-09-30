# Coid

Coid is a utility for compressing or decompressing any UUID generated from Cocos Creator.

## Install

```shell
$ go mod download
```

## Build

```shell
$ go build -ldflags "-s -w"
```

## Usage

To decompress uuid:

```shell
$ util fcmR3XADNLgJ1ByKhqcC5Z
fc991dd7-0033-4b80-9d41-c8a86a702e59 
```

To compress uuid:

```shell
$ util -c fc991dd7-0033-4b80-9d41-c8a86a702e59
fcmR3XADNLgJ1ByKhqcC5Z
```

## License

[MIT](LICENSE)
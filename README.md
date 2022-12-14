# connartist

[![Go Report Card](https://goreportcard.com/badge/github.com/adrianosela/connartist)](https://goreportcard.com/report/github.com/adrianosela/connartist)
[![Documentation](https://godoc.org/github.com/adrianosela/connartist?status.svg)](https://godoc.org/github.com/adrianosela/connartist)
[![GitHub issues](https://img.shields.io/github/issues/adrianosela/connartist.svg)](https://github.com/adrianosela/connartist/issues)
[![license](https://img.shields.io/github/license/adrianosela/connartist.svg)](https://github.com/adrianosela/connartist/blob/master/LICENSE)

A small framework that makes it easy to extend, override, and mock [net.Conn](https://pkg.go.dev/net#Conn) interface implementations.

### Usage

Easily override the functions of a [net.Conn](https://pkg.go.dev/net#Conn) implementation by wrapping it in a `connartist.Conn` and setting overrides as follows:

```
c, err := l.Accept()
if err != nil {
	fmt.Println(err)
	return
}

conn := connartist.New(c).
	WithRead(readHandler).
	WithWrite(writeHandler)
defer conn.Close()
```

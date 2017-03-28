# keygen

[![Build Status](https://travis-ci.org/northbright/keygen.svg?branch=master)](https://travis-ci.org/northbright/keygen)
[![Go Report Card](https://goreportcard.com/badge/github.com/northbright/keygen)](https://goreportcard.com/report/github.com/northbright/keygen)
[![GoDoc](https://godoc.org/github.com/northbright/keygen?status.svg)](https://godoc.org/github.com/northbright/keygen)

keygen is a [Golang](http://golang.org) package which provides helper functions to generate symmetric key(Ex: HMAC).

#### Create Symmetric Algorithm Key for HMAC SHA-256

    var err error
    size := 256  // Key size = 256 bits.
    key := make([]byte, size)

    if key, err = keygen.GenSymmetricKey(size); err != nil {
        fmt.Printf("GenSymmetricKey(%v) error: %v\n", size, err)
    }

    // Compare a zero-value byte array to see if key is generated.
    fmt.Printf("%v", !bytes.Equal(key, make([]byte, size)))

#### [CLI to generate symmetric key](https://github.com/northbright/keygen/tree/master/cli/gensymmetrickey)
* [gensymmetrickey](https://github.com/northbright/keygen/tree/master/cli/gensymmetrickey)

  * Generate symmetric key(Ex: HMAC).
  * Usage:  
      `gensymmetrickey -b=<key size in bits> -o=<output file>`
      
          Ex:
          gensymmetrickey -b=512 -o="mykey.dat"
          
  * [Binary Release](https://github.com/northbright/keygen/releases)

#### Documentation
* [API Reference](http://godoc.org/github.com/northbright/keygen)

#### License
* [MIT License](./LICENSE)

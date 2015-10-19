# keygen

Package keygen is a helper to generate symmetric key(Ex: HMAC).

#### Create Symmetric Algorithm Key for HMAC SHA-256

    var err error
    size := 256  // Key size = 256 bits.
    key := make([]byte, size)

    if key, err = keygen.GenSymmetricKey(size); err != nil {
        fmt.Printf("GenSymmetricKey(%v) error: %v\n", size, err)
    }

    // Compare a zero-value byte array to see if key is generated.
    fmt.Printf("%v", !bytes.Equal(key, make([]byte, size)))

#### [CLI](https://github.com/northbright/keygen/tree/master/cli/gensymmetrickey) to generate symmetric key
* gensymmetrickey

  * Generate symmetric key(Ex: HMAC).
  * Usage:  
      `gensymmetrickey -b=<key size in bits> -o=<output file>`
      
          Ex:
          gensymmetrickey -b=512 -o="mykey.dat"
          

#### Documentation
* [API Reference](http://godoc.org/github.com/northbright/keygen)

#### License
* [MIT License](./LICENSE)

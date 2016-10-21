package keygen

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/northbright/pathhelper"
)

const (
	keySizeError string = "Key size error" // Key size error message.
)

// GenSymmetricKey generates key byte array for symmetric key alogorithm like HMAC.
//
//   Params:
//       size: Key size in bits. Ex: 256 bits for HMAC SHA-256 and 512 bits for HMAC SHA-512.
//   Return:
//       k: Key stored in byte array.
//       err: error.
func GenSymmetricKey(bits int) (k []byte, err error) {
	if bits <= 0 || bits%8 != 0 {
		return nil, fmt.Errorf(keySizeError)
	}

	size := bits / 8
	k = make([]byte, size)
	if _, err = rand.Read(k); err != nil {
		return nil, err
	}

	return k, nil
}

// GenSymmetricKeyFile generates key file for symmetric key alogrithm like HMAC.
//
//   Params:
//       size: Key size in bits. Ex: 256 bits for HMAC SHA-256 and 512 bits for HMAC SHA-512.
//       outputFile: Output key file.
//       perm: Permission of key file. The final permission of file will be perm - umask. Ex: 0666 - 002(centos) = 0664.
//   Return:
//       err: error.
func GenSymmetricKeyFile(bits int, outputFile string, perm os.FileMode) (err error) {
	if len(outputFile) == 0 {
		return fmt.Errorf("Empty output file")
	}

	p := ""
	if p, err = pathhelper.GetAbsPath(outputFile); err != nil {
		return err
	}

	k, err := GenSymmetricKey(bits)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(p, k, perm)
}

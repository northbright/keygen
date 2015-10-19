package keygen

import (
	"crypto/rand"
	"errors"
	"github.com/northbright/errorhelper"
	"github.com/northbright/pathhelper"
	"io/ioutil"
	"os"
)

const (
	KeySizeError string = "Key size error" // Key size error message.
)

// GenSymmetricKey() generates key byte array for symmetric key alogorithm like HMAC.
//
//   Params:
//       size: Key size in bits. Ex: 256 bits for HMAC SHA-256 and 512 bits for HMAC SHA-512.
//   Return:
//       k: Key stored in byte array.
//       err: error.
func GenSymmetricKey(bits int) (k []byte, err error) {
	if bits <= 0 || bits%8 != 0 {
		return nil, errors.New(KeySizeError)
	}

	size := bits / 8
	k = make([]byte, size)
	if _, err = rand.Read(k); err != nil {
		return nil, err
	}

	return k, nil
}

// GenSymmetricKeyFile() generates key file for symmetric key alogrithm like HMAC.
//
//   Params:
//       size: Key size in bits. Ex: 256 bits for HMAC SHA-256 and 512 bits for HMAC SHA-512.
//       outputFile: Output key file.
//       perm: Permission of key file. The final permission of file will be perm - umask. Ex: 0666 - 002(centos) = 0664.
//   Return:
//       err: error.
func GenSymmetricKeyFile(bits int, outputFile string, perm os.FileMode) (err error) {
	if err = errorhelper.GenEmptyStringError(outputFile, "outputFile"); err != nil {
		return err
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
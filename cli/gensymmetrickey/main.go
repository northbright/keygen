package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/northbright/keygen"
)

const (
	defaultKeySize    int         = 512
	defaultOutputFile string      = "./key.dat"
	defaultPerm       os.FileMode = 0666
)

var (
	keySize    int    // Key Size
	outputFile string // Output Key File
)

func main() {

	flag.StringVar(&outputFile, "o", defaultOutputFile, "Output key file. Ex: -a='./mykey.dat'. Default: ./key.dat")
	flag.IntVar(&keySize, "b", defaultKeySize, "Key Size(bits). Ex: -b='256'. Default: 512 bits")

	flag.Parse()

	if err := keygen.GenSymmetricKeyFile(keySize, outputFile, defaultPerm); err != nil {
		fmt.Printf("Generate symmetric key error: %v\n", err)
	} else {
		fmt.Printf("Generate symmetric key successfully.\n")
	}
}

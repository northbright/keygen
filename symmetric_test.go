package keygen_test

import (
	"bytes"
	"fmt"
	"os"
	"os/user"

	"github.com/northbright/keygen"
	"github.com/northbright/pathhelper"
)

func ExampleGenSymmetricKey() {
	var err error
	size := 256 // Key size = 256 bits.
	key := make([]byte, size)

	if key, err = keygen.GenSymmetricKey(size); err != nil {
		fmt.Printf("GenSymmetricKey(%v) error: %v\n", size, err)
	}

	// Compare a zero-value byte array to see if key is generated.
	fmt.Printf("%v", !bytes.Equal(key, make([]byte, size)))
	// Output:
	// true
}

func ExampleGenSymmetricKeyFile() {
	var err error
	size := 512 // Key size = 256 bits.

	u, _ := user.Current()
	username := u.Username
	keyFile := fmt.Sprintf("/home/%v/my.key", username) // make absolute key path for current user.

	p, err := pathhelper.GetAbsPath(keyFile)
	if err != nil {
		fmt.Printf("GetAbsPath(%v) error: %v\n", keyFile, err)
	}

	fmt.Fprintf(os.Stderr, "Key file: %v\n", p)

	if err = keygen.GenSymmetricKeyFile(size, p, 0660); err != nil {
		fmt.Printf("GenSymmetricKeyFile(%v) error: %v\n", size, err)
	}

	// Output:
}

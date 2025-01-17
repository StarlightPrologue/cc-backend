// Copyright (C) 2022 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	// rand.Reader uses /dev/urandom on Linux
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "JWT_PUBLIC_KEY=%#v\nJWT_PRIVATE_KEY=%#v\n",
		base64.StdEncoding.EncodeToString(pub),
		base64.StdEncoding.EncodeToString(priv))
}

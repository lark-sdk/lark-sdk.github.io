package main

import (
  "github.com/chyroc/lark"
)

func main() {
  _ = lark.New(lark.WithAppCredential("<appID>", "<appSecret>"))
}

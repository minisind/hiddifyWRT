package main

import (
	_ "github.com/minisind/hiddifyWRT/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}

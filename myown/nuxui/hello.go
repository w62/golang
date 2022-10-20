package main

import (
    "nuxui.org/nuxui/nux"
    _ "nuxui.org/nuxui/ui"
)

func main() {
	nux.App().Init(manifest)
	nux.App().Run()

}

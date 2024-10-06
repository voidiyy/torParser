package settings

import (
	"fmt"
	"os"
)

func PrintHelp() {
	fmt.Println(
		`Modules:
    -domain     -> Domain name
    -tag        -> Tag name
    -attr       -> Attribute name
	-help       -> Display help
Example:
    $ ./parse google.com --tag a --attr href`)
	os.Exit(0)
}

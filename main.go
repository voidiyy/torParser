package main

import (
	"flag"
	"torParser/parser"
	"torParser/settings"
)

func main() {

	domain := flag.String("domain", "no", "domain name")
	htmlTag := flag.String("tag", "", "html tag")
	htmlAttr := flag.String("attr", "", "html attribute")
	help := flag.Bool("help", false, "help")

	if *help {
		settings.PrintHelp()
	}

	flag.Parse()

	sett := parser.SetSettings(*domain, *htmlTag, *htmlAttr)

	sett.PrintResult(sett.ParseUrl(), "\n")
}

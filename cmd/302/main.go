package main

import (
	"flag"
	"fmt"
	"os"

	redirect "github.com/nbari/302"
)

func exit1(err error) {
	fmt.Println(err)
	os.Exit(1)
}

var version string

func main() {
	var (
		d = flag.String("d", "302.db", "database file to store URL's")
		c = flag.String("c", "", "`config` file")
		p = flag.Int("p", 8000, "port to listen on")
		v = flag.Bool("v", false, fmt.Sprintf("Print version: %s", version))
	)

	flag.Parse()

	if *v {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	if *d == "" {
		exit1(fmt.Errorf("Missing database file, use(\"%s -h\") for help.\n", os.Args[0]))
	}

	if *c == "" {
		exit1(fmt.Errorf("Missing config file, use(\"%s -h\") for help.\n", os.Args[0]))
	}

	if *p < 1 || *p > 65535 {
		exit1(fmt.Errorf("Invalid port, use(\"%s -h\") for help.\n", os.Args[0]))
	}

	r, err := redirect.New(*c, *d)
	if err != nil {
		exit1(err)
	}

	err = r.Start(*p)
	if err != nil {
		exit1(err)
	}
}

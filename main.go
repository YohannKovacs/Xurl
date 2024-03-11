package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	exit = os.Exit

	flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	help         = flags.Bool("help", false, "Print usage instructions and exit.")
	printVersion = flags.Bool("version", false, "Print version information and exit.")
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	if *printVersion {
		// TODO(Print version information)
	}

	if *help {
		usage()
		exit(0)
	}

	args := flags.Args()
	if len(args) == 0 {

	}

	fHeadersOnly := UriSet.Bool("e", false, "Display only headers.")

	if len(os.Args) < 2 {
		panic("Please provide a subcommand: help or uri")
	}
	switch os.Args[1] {
	case "help":
		HelpSet.Parse(os.Args[2:])
		if *fHelp {
			HelpSet.Usage()
		} else if *fExtHelp {
			fmt.Println("Extended help")
			os.Exit(0)
		}
	case "uri":
		UriSet.Parse(os.Args[2:])
		if *fUri == "" {
			panic("Please provide a URI from using the -u flag")
		}
	}

	if strings.Split(*fUri, ".")[0] == "www" {
		*fUri, _ = strings.CutPrefix(*fUri, "www.")
		*fUri = strings.Join([]string{"http", *fUri}, "://")
	}

	resp, err := http.Get(*fUri)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Url not available.")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if *fHeadersOnly {
		headers := resp.Header
		for k, v := range headers {
			fmt.Println(k, ":\t", v)
		}
	} else {
		fmt.Println(string(body))
	}
}

// TODO(Usage information(help))
func usage() {

}

func fail(err error, msg string, args ...interface{}) {
	if err != nil {
		
	}
}

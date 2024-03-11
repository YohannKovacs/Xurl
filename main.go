package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	//Declaring flags

	//fFail := false
	//fData := false

	HelpSet := flag.NewFlagSet("help", flag.ExitOnError)
	fHelp := HelpSet.Bool("h", false, "Display help.")
	fExtHelp := HelpSet.Bool("x", false, "Display extended help.")

	UriSet := flag.NewFlagSet("uri", flag.ExitOnError)
	fUri := UriSet.String("u", "", "Request uri.")
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

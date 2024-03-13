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
	version = flags.Bool("version", false, "Print version information and exit.")
	headerOnly   = flags.Bool("headersOnly", false, "Print only headers and exit.")
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	if *version {
		fmt.Fprintln(os.Stdout, prettify(`
		Version : 0.0.1
		`))
		exit(0)
	}

	if *help {
		usage()
		exit(0)
	}

	args := flags.Args()
	if len(args) == 0 {
		fail(nil, "Too few arguments.")
	}

	addr := args[0]
	if addr == "" {
		fail(nil, "No host or port specified")
	}

	// FIXME(Can we handle www better?)
	if strings.Split(addr, ".")[0] == "www" {
		addr, _ = strings.CutPrefix(addr, "www.")
		addr = strings.Join([]string{"http", addr}, "://")
	}

	resp, err := http.Get(addr)
	if err != nil {
		fail(err, "Error while getting a response.")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fail(nil, "Uri provided is not available.")
	}

	if *headerOnly {
		for k, v := range resp.Header {
			fmt.Fprintln(os.Stdout, k, v)
		}
		exit(0)
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fail(err, "Error while reading response body.")
		}
		fmt.Fprintln(os.Stdout, string(body))
		exit(0)
	}
}

// TODO(Usage information(help))
func usage() {
	printLogo()
	fmt.Fprintf(os.Stderr, prettify(`
	Usage:
	%s [flags] [address]

	Xurl is a curl alternative cli program. The main purpose of xurl is to send data 
	to stdout to make it easy to chain it with other tools to do what you want.
	`), os.Args[0])
	fmt.Fprintln(os.Stderr)
	flags.PrintDefaults()
}

func fail(err error, msg string, args ...interface{}) {
	if err != nil {
		msg += ": %v"
		args = append(args, msg)
	}
	fmt.Fprintf(os.Stderr, msg, args)
	fmt.Fprintln(os.Stderr)
	if err != nil {
		exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "Try %s -help for more details.\n", os.Args[0])
		exit(2)
	}
}

func prettify(docString string) string {
	parts := strings.Split(docString, "\n")
	j := 0

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		parts[j] = part
		j++
	}
	return strings.Join(parts[:j], "\n")
}

func printLogo() {
	logo, err := os.ReadFile("logo.txt")
	if err != nil {
		fail(err, "Error while reading the logo file.")
	}

	fmt.Fprintln(os.Stderr, string(logo))
}
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	help       = flags.Bool("help", false, "Print usage instructions and exit.")
	version    = flags.Bool("version", false, "Print version information and exit.")
	headerOnly = flags.Bool("headersOnly", false, "Print only response headers and exit.")
	data       = flags.Bool("data", false, "Add Json data to the request object.")
)

func main() {
	var (
		addr    string
		scheme  string
		reqData []byte
	)

	flags.Usage = usage
	flags.Parse(os.Args[1:])

	if *version {
		fmt.Fprintln(os.Stdout, prettify(`
		Version : 0.0.1
		`))
		os.Exit(0)
	}

	if *help {
		usage()
		os.Exit(0)
	}

	args := flags.Args()
	if len(args) == 0 {
		fail(nil, "Too few arguments.")
	}

	addr = strings.TrimSpace(args[len(args)-1])
	if addr == "" {
		fail(nil, "No host or port specified")
	}

	if strings.Split(addr, ".")[0] == "www" {
		addr, _ = strings.CutPrefix(addr, "www.")
		addr = strings.Join([]string{"http", addr}, "://")
	}
	scheme = strings.TrimSpace(strings.Split(addr, "://")[0])

	if *data {
		dataArg := args[:len(args)-1]
		if strings.ContainsRune(dataArg[0], '@') {
			var err error
			reqData, err = os.ReadFile(strings.Split(dataArg[0], "@")[1])
			if err != nil {
				fail(err, "Error while reading request data file")
			}
		}
	}

	if scheme != "" {
		switch scheme {
		case "http", "https":
			fmt.Println(reqData)
			sch := &DefaultScheme{
				reqData: &reqData,
				addr:    addr,
				method:  "//TODO()",
			}
			sch.MakeRequest()
		case "file":
		case "ws":
			sch := &WebsocketScheme{
				url:    addr,
				origin: "http://localhost",
				data:	&reqData,
			}
			sch.MakeRequest()
		case "ftp":
		case "telnet":
		case "imap":
		default:
		}
	}
}

func usage() {
	logo, err := os.ReadFile("logo.txt")
	if err != nil {
		fail(err, "Error while reading the logo file.")
	}
	fmt.Fprintln(os.Stderr, string(logo))

	fmt.Fprintf(os.Stderr, prettify(`
	Usage:
	%s [flags] [data] [address]

	Xurl is a curl alternative cli program. The main purpose of xurl is to send data 
	to stdout to make it easy to chain it with other tools to do what you want.
	`), os.Args[0])
	fmt.Fprintln(os.Stderr)
	flags.PrintDefaults()
}

func fail(err error, msg string, args ...interface{}) {
	if err != nil {
		msg += ": %v"
		args = append(args, err)
	}
	fmt.Fprintf(os.Stderr, msg, args)
	fmt.Fprintln(os.Stderr)
	if err != nil {
		os.Exit(1)
	} else {
		// nil err means it was cli usage issue
		fmt.Fprintf(os.Stderr, "Try %s -help for more details.\n", os.Args[0])
		os.Exit(2)
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

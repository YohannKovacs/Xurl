package main

import (
	"flag"
)

// TODO(Figure out how to handle flags)

type ProgramFlags struct {
	UrlSet     *flag.FlagSet
	HelpSet    *flag.FlagSet
	Help       *bool
	ExtHelp    *bool
	Fail       *bool
	Data       *bool
	Uri        *string
	HeaderOnly *bool
}

func (f *ProgramFlags) DeclareFlags() {
	f.HelpSet = flag.NewFlagSet("Help Flagset", flag.ExitOnError)
	f.HelpSet.BoolVar(f.Help, "h", false, "Display help.")
	f.HelpSet.BoolVar(f.ExtHelp, "x", false, "Display extended help.")

	f.UrlSet = flag.NewFlagSet("Url Flagset", flag.ExitOnError)
	f.UrlSet.StringVar(f.Uri, "u", "", "Request Url")
	f.UrlSet.BoolVar(f.HeaderOnly, "e", false, "Display only headers of the request")

	flag.Parse()
}

func (f *ProgramFlags) ProcessHelpFlagSet() {
	if *f.Help {
		flag.Usage()
	} else if *f.ExtHelp {
		// Not sure what this does
		flag.CommandLine.Usage()
	}

}


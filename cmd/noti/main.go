package main

import (
	"log"
	"os"

	"github.com/variadico/noti/cmd/noti/cli"
	"github.com/variadico/noti/cmd/noti/cli/banner"
	"github.com/variadico/noti/cmd/noti/cli/root"
	"github.com/variadico/noti/cmd/noti/cli/slack"
	"github.com/variadico/noti/cmd/noti/cli/speech"
	"github.com/variadico/noti/cmd/noti/cli/version"
)

func main() {
	log.SetFlags(0)

	noti := root.NewCommand().(*root.Command)
	if err := noti.Parse(os.Args[1:]); err != nil {
		log.Println("Error:", err)
		log.Fatalln("Try 'noti -help' for more information.")
	}

	noti.Cmds = map[string]cli.Cmd{
		"version": version.NewCommand(),
		"banner":  banner.NewCommand().(cli.Cmd),
		"speech":  speech.NewCommand().(cli.Cmd),
		"slack":   slack.NewCommand().(cli.Cmd),
	}

	if len(noti.Args()) == 0 {
		// noti was called by itself.
		if err := noti.Run(); err != nil {
			log.Fatalln("Error:", err)
		}
		return
	}

	var cmd cli.Cmd
	var found bool
	cmd, found = noti.Cmds[noti.Args()[0]]

	if found {
		// Command is something like: noti foo ls
		if err := cmd.Parse(noti.Args()[1:]); err != nil {
			log.Println("Error:", err)
			log.Fatalf("Try 'noti %s -help' for more information.", noti.Args()[0])
		}
	} else {
		// Command is something like: noti ls
		cmd = noti
	}

	if err := cmd.Run(); err != nil {
		log.Fatalln("Error:", err)
	}
}

package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	text "github.com/multiverse-os/cli-framework/text"
	color "github.com/multiverse-os/cli-framework/text/color"
)

// TODO: Lets not use global variables, it just doesnt feel right

// TODO: Why are commands VisibleCategories? This section is practically unreadable and very hard to customize
// TODO: All lines should be checked for length of 80 and broken into new line if so with the correct tab spacing prefixing it
// TODO: Use table library code to improve the structure of this and do better alignment of values
func (self *CLI) PrintHelp() {
	if self.Description != "" {
		fmt.Println(color.Strong(self.Description))
	}
	if self.Usage != "" {
		if self.NoANSI {
			fmt.Println("Usage")
		} else {
			fmt.Println(color.Strong("Usage"))
		}
		fmt.Print(text.Repeat(" ", 4))

		if self.NoANSI {
			fmt.Print(self.Name)
		} else {
			fmt.Print(color.Header(self.Name))
		}
		if self.HasVisibleFlags() {
			fmt.Print(" [options]")
		} else {
		}
		if self.HasVisibleCommands() {
			fmt.Print(" command [command options]")
		}
		if self.ArgsUsage != "" {
			fmt.Println(" " + self.ArgsUsage)
		} else {
			fmt.Println("[arguments...]")
		}
		if self.HasVisibleFlags() {
			fmt.Println("\n" + color.Strong("Options"))
			for _, flag := range self.VisibleFlags() {
				fmt.Println("flag: ", flag)
			}
		}
	}
}

var CLIHelpTemplate = `{{range $index, $option := .VisibleFlags}}{{if $index}}{{"\n"}}{{end}}{{"\t\t"}}{{$option}}{{end}}{{"\n"}}{{if .VisibleCategories}}{{"\n"}}` +
	fmt.Sprintf(color.STRONG) + `Commands` + fmt.Sprintf(color.RESET) + `{{range .VisibleCategories}}{{if .Name}}{{"\n"}}{{.Name}}:{{end}}{{end}}{{range .VisibleCommands}}{{"\n\t "}}` + fmt.Sprintf(color.H1) + ` {{join .Names ", "}}` + fmt.Sprintf(color.RESET) + `{{"\t"}}{{.Usage}}{{end}}{{end}}{{"\n"}}`

var CommandHelpTemplate = fmt.Sprintf(color.H1) + `{{.Name}}` + fmt.Sprintf(color.RESET) + ` - {{.Usage}}{{"\n"}}` + fmt.Sprintf(color.H1) + `Usage` + fmt.Sprintf(color.RESET) +
	`{{"\n"}}{{if .UsageText}}{{.UsageText}}{{else}}{{.Name}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}

` + fmt.Sprintf(color.H2) + `Category` + fmt.Sprintf(color.RESET) + `
   {{.Category}}{{end}}{{if .Description}}

` + fmt.Sprintf(color.H2) + `Description` + fmt.Sprintf(color.RESET) + `
   {{.Description}}{{end}}{{if .VisibleFlags}}

` + fmt.Sprintf(color.H2) + `Options` + fmt.Sprintf(color.RESET) + `
   {{range .VisibleFlags}}{{.}}{{end}}{{end}}
`

var SubcommandHelpTemplate = `Name
   ` + fmt.Sprintf(color.H1) + `{{.HelpName}}` + fmt.Sprintf(color.RESET) + ` - {{if .Description}}{{.Description}}{{else}}{{.Usage}}{{end}}

` + fmt.Sprintf(color.H2) + `Usage` + fmt.Sprintf(color.RESET) + `
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} command{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}

` + fmt.Sprintf(color.H2) + `Commands` + fmt.Sprintf(color.RESET) + `{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{end}}{{range .VisibleCommands}}{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{if .VisibleFlags}}
` + fmt.Sprintf(color.H2) + `Options` + fmt.Sprintf(color.RESET) + `
   {{range .VisibleFlags}}{{.}}{{end}}{{end}}
`

type helpPrinter func(w io.Writer, templ string, data interface{})

var HelpPrinter helpPrinter = printHelp

func ShowCLIHelpAndExit(c *Context, exitCode int) {
	ShowCLIHelp(c)
	os.Exit(exitCode)
}

func ShowCLIHelp(c *Context) {
	c.CLI.PrintBanner()
	c.CLI.PrintHelp()
	HelpPrinter(c.CLI.Writer, CLIHelpTemplate, c.CLI)
}

func DefaultCLIComplete(c *Context) {
	for _, command := range c.CLI.Commands {
		if command.Hidden {
			continue
		}
		for _, name := range command.Names() {
			fmt.Fprintln(c.CLI.Writer, name)
		}
	}
}

func ShowCommandHelpAndExit(c *Context, command string, code int) {
	ShowCommandHelp(c, command)
	os.Exit(code)
}

func ShowCommandHelp(ctx *Context, command string) error {
	if command == "" {
		HelpPrinter(ctx.CLI.Writer, SubcommandHelpTemplate, ctx.CLI)
		return nil
	}

	for _, c := range ctx.CLI.Commands {
		if c.HasName(command) {
			HelpPrinter(ctx.CLI.Writer, CommandHelpTemplate, c)
			return nil
		}
	}

	if ctx.CLI.CommandNotFound == nil {
		return NewExitError(fmt.Sprintf("No help topic for '%v'", command), 3)
	}

	ctx.CLI.CommandNotFound(ctx, command)
	return nil
}

// ShowSubcommandHelp prints help for the given subcommand
func ShowSubcommandHelp(c *Context) error {
	return ShowCommandHelp(c, c.Command.Name)
}

func PrintVersion(c *Context) {
	fmt.Fprintf(c.CLI.Writer, "%v version %v\n", c.CLI.Name, c.CLI.Version.String())
}

// ShowCompletions prints the lists of commands within a given context
func ShowCompletions(c *Context) {
	a := c.CLI
	if a != nil && a.BashComplete != nil {
		a.BashComplete(c)
	}
}

// ShowCommandCompletions prints the custom completions for a given command
func ShowCommandCompletions(ctx *Context, command string) {
	c := ctx.CLI.Command(command)
	if c != nil && c.BashComplete != nil {
		c.BashComplete(ctx)
	}
}

func printHelp(out io.Writer, templ string, data interface{}) {
	funcMap := template.FuncMap{
		"join": strings.Join,
	}
	w := tabwriter.NewWriter(out, 1, 8, 2, ' ', 0)
	t := template.Must(template.New("help").Funcs(funcMap).Parse(templ))
	err := t.Execute(w, data)
	if err != nil {
		return
	}
	w.Flush()
}

func checkVersion(c *Context) bool {
	found := false
	if VersionFlag.GetName() != "" {
		eachName(VersionFlag.GetName(), func(name string) {
			if c.GlobalBool(name) || c.Bool(name) {
				found = true
			}
		})
	}
	return found
}

func checkHelp(c *Context) bool {
	found := false
	if HelpFlag.GetName() != "" {
		eachName(HelpFlag.GetName(), func(name string) {
			if c.GlobalBool(name) || c.Bool(name) {
				found = true
			}
		})
	}
	return found
}

func checkCommandHelp(c *Context, name string) bool {
	if c.Bool("h") || c.Bool("help") {
		ShowCommandHelp(c, name)
		return true
	}
	return false
}

func checkSubcommandHelp(c *Context) bool {
	if c.Bool("h") || c.Bool("help") {
		ShowSubcommandHelp(c)
		return true
	}
	return false
}

func checkShellCompleteFlag(a *CLI, arguments []string) (bool, []string) {
	if !a.BashCompletion {
		return false, arguments
	}
	pos := len(arguments) - 1
	lastArg := arguments[pos]
	if lastArg != "--"+BashCompletionFlag.GetName() {
		return false, arguments
	}
	return true, arguments[:pos]
}

func checkCompletions(c *Context) bool {
	if !c.shellComplete {
		return false
	}
	if args := c.Args(); args.Present() {
		name := args.First()
		if cmd := c.CLI.Command(name); cmd != nil {
			// let the command handle the completion
			return false
		}
	}
	ShowCompletions(c)
	return true
}

func checkCommandCompletions(c *Context, name string) bool {
	if !c.shellComplete {
		return false
	}

	ShowCommandCompletions(c, name)
	return true
}

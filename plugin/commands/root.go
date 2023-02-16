/**
 * (C) Copyright IBM Corp. 2023.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
	"errors"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	translation "github.com/IBM/platform-services-go-sdk/i18n"
	"github.com/IBM/platform-services-go-sdk/plugin/commands/projectsv1"
	"github.com/IBM/platform-services-go-sdk/plugin/version"
	"github.com/IBM/platform-services-go-sdk/utils"
	"github.com/spf13/cobra"
	"strings"
)

var u *utils.Utils

var rootCommand *cobra.Command

func Init() {
	ui := terminal.NewStdUI()
	u = utils.NewUtils(ui)
	rootCommand = projectsv1.GetProjectsV1Command(u)
	rootCommand.SilenceErrors = true
	rootCommand.Version = version.GetPluginVersion().String()

	// cobra doesnt have a good way of printing the "name and alias" of commands
	// in a list without messing up the padding. this custom template function
	// will allow for dynamically generated uniform padding for listing commands
	// along with their aliases
	cobra.AddTemplateFunc("getNameAndAliasPadding", func(cmds []*cobra.Command) int {
		result := 0
		for _, cmd := range cmds {
			minPadLength := len(cmd.NameAndAliases())
			if minPadLength > result {
				result = minPadLength
			}
		}

		// add an extra two characters of spacing, otherwise the strings
		// look like they run together
		return result + 2
	})

	// make the help menu template headers translatable
	cobra.AddTemplateFunc("printHeader", func(translationId string) string {
		return translation.T(translationId)
	})

	// make the usage line at the end of the help meny translatable
	cobra.AddTemplateFunc("printFinalUsageLine", func(commandPath string) string {
		return translation.T("final-usage-line-help-menu", map[string]interface{}{
			"COMMAND_PATH": commandPath,
		})
	})

	// mark the "help" flag as global to be consistent with the ibm core cli
	rootCommand.PersistentFlags().BoolP("help", "h", false, translation.T("show-help"))

	// Add the Cobra help command
	rootCommand.InitDefaultHelpCmd()

	// Disable the default Cobra completion command - we use ibmcloud for autocompletion
	rootCommand.CompletionOptions.DisableDefaultCmd = true

	// Hide the help command (this just prevents it from appearing in the list)
	for _, cmd := range rootCommand.Commands() {
		if cmd.Name() == "help" {
			cmd.Hidden = true
			// force a translatable description, this isn't "hidden" in the `plugin show` command
			cmd.Short = translation.T("show-help")
		}
	}

	rootCommand.SetHelpTemplate(`{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{else}}{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}
{{end}}{{end}}`)

	rootCommand.SetVersionTemplate(`ibmcloud {{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}
`)

	rootCommand.SetUsageTemplate(`{{printHeader "name-header-help-menu"}}:
  {{.Name}} - {{.Long}}

{{printHeader "usage-header-help-menu"}}:
  ibmcloud {{.UseLine}}{{if gt (len .Aliases) 0}}

{{printHeader "aliases-header-help-menu"}}:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

{{printHeader "examples-header-help-menu"}}:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

{{printHeader "commands-header-help-menu"}}:{{$nameAndAliasPadding := (getNameAndAliasPadding .Commands)}}{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .NameAndAliases $nameAndAliasPadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

{{printHeader "options-header-help-menu"}}:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

{{printHeader "global-options-header-help-menu"}}:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

{{printHeader "additional-help-topics-header-help-menu"}}:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

{{printFinalUsageLine .CommandPath}}{{end}}
`)

	// cobra handles "unknown commands" differently than the ibm sdk. it wants to print
	// the help menu unconditionally. to be more consistent with other ibm plugins,
	// we should alert the user when a command is not supported. we can also take this a step
	// further and print suggestions for commands close in spelling to what the user entered.
	// to do this, we override the "HelpFunc" here:
	defaultHelpFunc := rootCommand.HelpFunc()
	rootCommand.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		// args are any strings that come after the plugin root command
		// for example, 'watson' would not be an arg, but 'compare-comply'
		// would be, even if the matched command is the compare-comply service command

		// root level command, print the help menu
		if len(args) == 0 {
			defaultHelpFunc(cmd, args)
			return
		}

		// we are in this function because the command is not "runnable" (has no Run function)
		// the command is either:
		// 1) a namespace, i.e. a parent command with subcommands but no "Run" function of its own
		// 2) an unknown/mistyped command. is this case the "cmd" will be for the lowest level matching command
		userSuppliedCommand := args[len(args)-1]

		// if the user entered a valid namespace without a Run command
		// or if the user explicitly invoked the "help" command or help flag (--help, -h),
		// print the help menu
		if isNamespaceCommand(userSuppliedCommand, cmd) || isHelpCommandOrFlag(userSuppliedCommand) {
			defaultHelpFunc(cmd, args)
			return
		}

		// if we are at this point, the supplied command is not registered
		// fail the execution and print any suggestions, if applicable
		cmdNotRegistered := translation.T("unregistered-command", map[string]interface{}{
			"COMMAND":   userSuppliedCommand,
			"NAMESPACE": getFullNamespace(cmd),
		})
		ui.Failed(cmdNotRegistered)

		suggestions := cmd.SuggestionsFor(userSuppliedCommand)
		if len(suggestions) > 0 {
			u.Say(translation.T("suggestions", map[string]interface{}{
				"SUGGESTIONS": strings.Join(suggestions, ", "),
			}))
		}
	})
}

func isNamespaceCommand(userSupplied string, cmd *cobra.Command) bool {
	closestMatchingParent := cmd.Name()

	// return true if the user specified the name, or an alias of, the parent
	return userSupplied == closestMatchingParent || cmd.HasAlias(userSupplied)
}

func isHelpCommandOrFlag(arg string) bool {
	return arg == "help" || arg == "-h" || arg == "--help"
}

func getFullNamespace(cmd *cobra.Command) string {
	if cmd.HasParent() {
		return getFullNamespace(cmd.Parent()) + " " + cmd.Name()
	} else {
		return cmd.Name()
	}
}

func Execute(context plugin.PluginContext, args []string) {
	u.SetContext(context)
	u.SetCommandName(args)

	// try and get cobra to use the full list of args
	rootCommand.SetArgs(replicateOriginalArgs(context.CommandNamespace(), args))

	// pass the error returned from cobra execution to the translator function
	// to allow for user customization of cobra errors
	err := translateCobraError(rootCommand.Execute())
	u.HandleError(err, translation.T("root-command-error"))
}

// the core cli recognizes all of the parent commands due to their existence in the plugin metadata.
// it extracts only the final word in the command string to pass to this plugin. cobra needs the
// fully resolved command from the user to operate properly, so we add back in the masked arguments here
func replicateOriginalArgs(maskedArgs string, args []string) []string {
	// plugins can run outside of the ibmcloud context, in which case the masked args will be empty
	if maskedArgs == "" {
		return args
	}
	masked := strings.Fields(maskedArgs) // get the args as an array
	return append(masked[1:], args...)   // don't include the first masked arg, it's the plugin name
}

func GetNamespaceAndCommandMetadata(rootNamespace string) (n []plugin.Namespace, c []plugin.Command) {
	// define root namespace metadata
	n = []plugin.Namespace{
		{
			Name: rootNamespace,
			Description: translation.T("project-short-description"),
		},
	}

	// gather metadata for all nested commands
	namespaces, commands := getMetadataForCommands(rootCommand.Commands(), rootNamespace)

	// add nested namespaces (commands with subcommands) to the list
	n = append(n, namespaces...)

	// hack to convince the ibm cli sdk to allow a "version" flag
	// on the root command - the code wont reach the cobra handler otherwise
	c = append(commands, plugin.Command{
		Namespace: rootCommand.Name(),
		Name:      "--version",
		Aliases:   []string{"-v"},
		Hidden:    true,
	})

	return
}

func getMetadataForCommands(cmds []*cobra.Command, namespace string) ([]plugin.Namespace, []plugin.Command) {
	namespaces, commands := []plugin.Namespace{}, []plugin.Command{}

	for _, cmd := range cmds {
		n, c := getCommandMetadata(cmd, namespace)
		namespaces = append(namespaces, n...)
		commands = append(commands, c...)
	}

	return namespaces, commands
}

func getCommandMetadata(command *cobra.Command, namespace string) ([]plugin.Namespace, []plugin.Command) {
	if command.HasSubCommands() {
		namespaces, cmds := getMetadataForCommands(command.Commands(), namespace+" "+command.Name())
		namespaceMetadata := plugin.Namespace{
			ParentName:  namespace,
			Name:        command.Name(),
			Description: command.Short,
			Aliases:     command.Aliases,
		}

		return append(namespaces, namespaceMetadata), cmds
	}

	cmd := getPluginCommand(command, namespace)

	// Return blank plugin namespace and converted command
	return nil, []plugin.Command{cmd}
}

func getPluginCommand(cmd *cobra.Command, namespace string) plugin.Command {
	return plugin.Command{
		Namespace:   namespace,
		Name:        cmd.Name(),
		Description: cmd.Short,
		Usage:       cmd.Use,
		Aliases:     cmd.Aliases,
		Flags:       plugin.ConvertCobraFlagsToPluginFlags(cmd),
		Hidden:      cmd.Hidden || cmd.Name() == "help", // hide the help command
	}
}

// the purpose of this method is to take error messages hardcoded in cobra
// and make them translatable, to support cli plugins in languages other
// than english. common errors are recognized and the variable values
// are parsed from the message and passed to the translation function.
// "templates" of the expected formats are used to identify where in the
// message to extract the variables.
func translateCobraError(err error) error {
	if err == nil {
		return err
	}

	// if the error isn't recognized by this function, we want to return
	// the original error
	msg := err.Error()

	if strings.HasPrefix(msg, "required flag(s)") {
		// e.g. required flag(s) "resource-id" not set
		template := "required flag(s) {} not set"
		msg = translation.T("missing-required-flags-error", map[string]interface{}{
			"FLAGS": extract(template, msg)[0],
		})
	} else if strings.HasPrefix(msg, "flag needs an argument:") {
		// e.g. flag needs an argument: --resource-id
		template := "flag needs an argument: {}"
		msg = translation.T("missing-flags-argument-error", map[string]interface{}{
			"FLAG": extract(template, msg)[0],
		})
	} else if strings.HasPrefix(msg, "unknown shorthand flag:") {
		// e.g. unknown shorthand flag: 'r' in -r
		template := "unknown shorthand flag: {} in {}"
		values := extract(template, msg)
		msg = translation.T("unknown-shorthand-flag-error", map[string]interface{}{
			"UNKNOWN_SHORTHAND": values[0],
			"GIVEN_SHORTHANDS":  values[1],
		})
	} else if strings.HasPrefix(msg, "unknown flag:") {
		// e.g. unknown flag: --r
		template := "unknown flag: {}"
		msg = translation.T("unknown-flag-error", map[string]interface{}{
			"FLAG": extract(template, msg)[0],
		})
	}

	return errors.New(msg)
}

// uses the provided template string to strip variables values
// from a string and return them in a slice
func extract(template string, message string) []string {
	chunks := removeEmptyElements(strings.Split(template, "{}"))
	marker := "***"
	for _, chunk := range chunks {
		message = strings.ReplaceAll(message, chunk, marker)
	}

	return removeEmptyElements(strings.Split(message, marker))
}

// the "strings.Split" method in Go will return empty strings
// if the separator is at the beginning or end of a string.
// these empty string elements in the resulting slice cause
// problems, so this function strips them from the slice.
func removeEmptyElements(arr []string) []string {
	result := make([]string, 0)

	for _, e := range arr {
		if len(e) > 0 {
			result = append(result, e)
		}
	}

	return result
}

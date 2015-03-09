package commands

import (

	"github.com/cloudfoundry/cli/cf/command_metadata"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	. "github.com/cloudfoundry/cli/cf/i18n"
	"github.com/cloudfoundry/cli/cf/requirements"
	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/codegangsta/cli"
)

type AliasCommands struct {
	ui     terminal.UI
	config core_config.ReadWriter
}

func NewAlias(ui terminal.UI, config core_config.ReadWriter) AliasCommands {
	return AliasCommands{ui: ui, config: config}
}

func (cmd AliasCommands) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "alias",
		Description: T("create alias for api endpoint"),
		Usage:       T("CF_NAME alias API_URL ALIAS"),
	}
}

func (cmd AliasCommands) GetRequirements(_ requirements.Factory, _ *cli.Context) ([]requirements.Requirement, error) {
	return nil, nil
}

func (cmd AliasCommands) Run(context *cli.Context) {
        cmd.ui.Say("Alias")
}

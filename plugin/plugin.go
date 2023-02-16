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

package plugin

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/IBM/platform-services-go-sdk/plugin/commands"
	"github.com/IBM/platform-services-go-sdk/plugin/version"
)

const PluginName = "project"

type Plugin struct {
	ui terminal.UI
}

func (p *Plugin) GetMetadata() plugin.PluginMetadata {
	namespaces, cmds := commands.GetNamespaceAndCommandMetadata(PluginName)

	return plugin.PluginMetadata{
		Name:       PluginName,
		Version:    version.GetPluginVersion(),
		Namespaces: namespaces,
		Aliases: namespaces[0].Aliases,
		Commands:   cmds,
		IsCobraPlugin: true,
	}
}

func (p *Plugin) Run(context plugin.PluginContext, args []string) {
	p.ui = terminal.NewStdUI()
	commands.Execute(context, args)
}

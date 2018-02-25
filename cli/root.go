//
// Copyright 2018 Bryan T. Meyers <bmeyers@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cli

import (
	"github.com/DataDrake/cli-ng/cmd"
)

// Root is the main command for this application
var Root *cmd.RootCMD

func init() {
	// Build Application
	Root = &cmd.RootCMD{
		Name:  "proc-maps",
		Short: "Analyze the contents of /proc/[pid]/maps",
	}
	// Setup the Sub-Commands
	Root.RegisterCMD(&cmd.Help)
	Root.RegisterCMD(&FileStats)
	Root.RegisterCMD(&RankAllSizes)
	Root.RegisterCMD(&RankSizes)
}

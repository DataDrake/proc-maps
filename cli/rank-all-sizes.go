//
// Copyright 2017-202q Bryan T. Meyers <root@datadrake.com>
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
	"github.com/DataDrake/proc-maps/data"
	"github.com/DataDrake/proc-maps/sorts"
	"os"
)

// RankAllSizes reads all maps and sorts files by size
var RankAllSizes = cmd.CMD{
	Name:  "rank-all-sizes",
	Alias: "ras",
	Short: "Read all maps and sort files by size",
	Args:  &RankAllSizesArgs{},
	Run:   RankAllSizesRun,
}

// RankAllSizesArgs contains the arguments for the "RankAllSizes" subcommand
type RankAllSizesArgs struct{}

// RankAllSizesRun carries out reading parsing, sorting, and listing
func RankAllSizesRun(r *cmd.RootCMD, c *cmd.CMD) {
	//args := c.Args.(*RankAllSizesArgs)
	s := sorts.NewBySize(data.ParseMaps())
	s.Print()
	os.Exit(0)
}

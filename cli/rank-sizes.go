//
// Copyright 2017 Bryan T. Meyers <bmeyers@datadrake.com>
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

// RankSizes reads a single map and sorts files by size
var RankSizes = cmd.CMD{
	Name:  "rank-sizes",
	Alias: "rs",
	Short: "Read a single map and sort files by size",
	Args:  &RankSizesArgs{},
	Run:   RankSizesRun,
}

// RankSizesArgs contains the arguments for the "RankSizes" subcommand
type RankSizesArgs struct {
	PID string `desc:"The Process ID of the map to scan"`
}

// RankSizesRun carries out reading parsing, sorting, and listing
func RankSizesRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*RankSizesArgs)
	entries := make(data.EntryMap)
	data.ParseMap(args.PID, entries)
	s := sorts.NewBySize(entries)
	s.Print()
	os.Exit(0)
}

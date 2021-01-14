//
// Copyright 2017-2021 Bryan T. Meyers <root@datadrake.com>
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
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"github.com/DataDrake/proc-maps/data"
	"os"
)

// FileStats reads all maps and gathers the stats for a single file
var FileStats = cmd.CMD{
	Name:  "file-stats",
	Alias: "fs",
	Short: "Reads all maps and gathers the stats for a single file",
	Args:  &FileStatsArgs{},
	Run:   FileStatsRun,
}

// FileStatsArgs contains the arguments for the "FileStats" subcommand
type FileStatsArgs struct {
	Path string `desc:"The full path to the file"`
}

// FileStatsRun carries out parsing and then prints out a single file
func FileStatsRun(r *cmd.RootCMD, c *cmd.CMD) {
	args := c.Args.(*FileStatsArgs)
	for _, entry := range data.ParseMaps() {
		if entry.Name == args.Path {
			entry.Print()
			os.Exit(0)
		}
	}
	fmt.Fprintf(os.Stderr, "File at '%s' not found\n", args.Path)
	os.Exit(1)
}

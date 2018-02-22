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

package main

import (
	"fmt"
    "github.com/DataDrake/proc-maps/data"
    "github.com/DataDrake/proc-maps/sorts"
	"os"
    "sort"
    "text/tabwriter"
)

func main() {
    entries := data.ParseMaps()
    s := make(sorts.BySize,0)
    var total uint64
	for _, entry := range entries {
        s = append(s, entry)
        total += entry.Total
	}
    sort.Sort(s)
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
    fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", "Rank", "Weight (B)", "Cumulative (B)", "Size (B)", "Filename")
    for i, entry := range s {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", i+1, data.CanonicalSize(entry.Weight), data.CanonicalSize(total), data.CanonicalSize(entry.Total), entry.Name)
        total -= entry.Total
    }
    w.Flush()
}

//
// Copyright 2018-2021 Bryan T. Meyers <root@datadrake.com>
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

package data

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

// EntryMap maps unique IDs (DeviceID:Inode) to their respective FileEntry
type EntryMap map[string]*FileEntry

// FileEntry is a struct representing a single mapped file
type FileEntry struct {
	// Name is the complete filepath of this file
	Name string
	// Weight is a relative score for this File's usage and size
	Weight uint64
	// Total memory used by mapping this file
	Total uint64
	// Sizes is a mapping of memory permissions to SizeEntries
	Sizes map[string]*SizeEntry
}

// Increment entry by new address range
func (f *FileEntry) Increment(startAddress, endAddress, permissions string) {
	start, err := strconv.ParseUint(startAddress, 16, 64)
	if err != nil {
		panic(err)
	}
	end, err := strconv.ParseUint(endAddress, 16, 64)
	if err != nil {
		panic(err)
	}
	f.Sizes[permissions] = &SizeEntry{
		Size: end - start,
		Refs: 1,
	}
	f.Total += end - start
	f.Weight += end - start
}

// Print summarizes the stats related to a FileEntry
func (f *FileEntry) Print() {
	fmt.Printf("File Information:\n\n")
	fmt.Printf("Name   : %s\n", f.Name)
	fmt.Printf("Size   : %sB\n", CanonicalSize(f.Total))
	fmt.Printf("Weight : %sB\n", CanonicalSize(f.Weight))
	fmt.Println("")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\n", "Permissions", "Size (B)", "References")
	for permission, entry := range f.Sizes {
		fmt.Fprintf(w, "%s\t%s\t%d\n", permission, CanonicalSize(entry.Size), entry.Refs)
	}
	w.Flush()
}

// SizeEntry is a struct representing the size and frequency of use for a mapped region corresponding to FileEntry
type SizeEntry struct {
	// Size is the memory allocated for this permission set
	Size uint64
	// Refs is a count of the number of processes with this range mapped
	Refs uint64
}

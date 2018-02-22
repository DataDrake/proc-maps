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

package data

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

// SizeEntry is a struct representing the size and frequency of use for a mapped region corresponding to FileEntry
type SizeEntry struct {
	// Size is the memory allocated for this permission set
	Size uint64
	// Refs is a count of the number of processes with this range mapped
	Refs uint64
}

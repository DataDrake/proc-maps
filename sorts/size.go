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

package sorts

import (
	"github.com/DataDrake/proc-maps/data"
)

// BySize is a list of FileEntries that can be sorted by size
type BySize []*data.FileEntry

// Len returns the length of the list
func (b BySize) Len() int {
	return len(b)
}

// Swap exchanges values at indices
func (b BySize) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less decides if the entry at 'i' should sort lower than the entry at 'j'
func (b BySize) Less(i, j int) bool {
	return b[i].Total > b[j].Total
}

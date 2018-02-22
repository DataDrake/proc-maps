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

// ByWeight is a list of FileEntries that can be sorted by weight
type ByWeight []*data.FileEntry

// Len returns the length of the list
func (b ByWeight) Len() int {
	return len(b)
}

// Swap exchanges values at indices
func (b ByWeight) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less decides if the entry at 'i' should sort lower than the entry at 'j'
func (b ByWeight) Less(i, j int) bool {
	// Sort order:
	// 1. Higher weight always wins
	// 2. Higher total wins in weights match
	return b[i].Weight > b[j].Weight || ((b[i].Weight == b[j].Weight) && (b[i].Total > b[j].Total))
}

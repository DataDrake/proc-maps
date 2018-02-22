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

import "strconv"

// CanonicalSize converts a number of Bytes into a human readable quantity (ie. 10M)
func CanonicalSize(size uint64) string {
	var suffix string
	switch {
	case size < 1000:
		suffix = ""
	case size < 1000000:
		suffix = "K"
		size /= 1000
	case size < 1000000000:
		suffix = "M"
		size /= 1000000
	case size < 1000000000000:
		suffix = "G"
		size /= 1000000000
	case size < 1000000000000000:
		suffix = "T"
		size /= 1000000000000
	default:
		suffix = "P"
		size /= 1000000000000000000
	}
	return strconv.FormatUint(size, 10) + suffix
}

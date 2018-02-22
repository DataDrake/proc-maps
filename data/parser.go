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

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

var pidMatch *regexp.Regexp

// Format:
//
// [1]          [2]          [3]  [4]      [5]  [6][7]
// 7ffda1448000-7ffda1469000 rw-p 00000000 00:00 0 [stack]
//
// Definitions:
//
// [1] Start Address
// [2] End Address
// [3] Permissions
// [4] Offset
// [5] Device ID
// [6] Inode on Device
// [7] Filename
var rowMatch *regexp.Regexp

func init() {
	pidMatch = regexp.MustCompile("\\d+")
	rowMatch = regexp.MustCompile("^(\\w+)-(\\w+)\\s+(\\S+)\\s+(\\d+)\\s+(\\S+)\\s+(\\d+)\\s*(\\S+)?$")
}

// ParseMap gets all of the entries for a specific map
func ParseMap(pid string, entries EntryMap) {
	maps, err := os.Open("/proc/" + pid + "/maps")
	if err != nil {
		log.Print(err)
	}
	defer maps.Close()
	buff := bufio.NewReader(maps)
	line, _, err := buff.ReadLine()
	for err == nil {
		raw := rowMatch.FindSubmatch(line)
		if raw != nil && !bytes.Equal(raw[5], []byte("00:00")) {
			key := string(raw[5]) + ":" + string(raw[6])
			if entry := entries[key]; entry != nil {
				if size := entry.Sizes[string(raw[3])]; size != nil {
					size.Refs++
					entry.Weight += size.Size
				} else {
					start, err := strconv.ParseUint(string(raw[1]), 16, 64)
					end, err := strconv.ParseUint(string(raw[2]), 16, 64)
					if err == nil {
						entry.Sizes[string(raw[3])] = &SizeEntry{
							Size: end - start,
							Refs: 1,
						}
						entry.Total += end - start
						entry.Weight += end - start
					}
				}
			} else {
				entry := &FileEntry{
					Name:  string(raw[7]),
					Sizes: make(map[string]*SizeEntry),
				}
				start, err := strconv.ParseUint(string(raw[1]), 16, 64)
				end, err := strconv.ParseUint(string(raw[2]), 16, 64)
				if err == nil {
					entry.Sizes[string(raw[3])] = &SizeEntry{
						Size: end - start,
						Refs: 1,
					}
					entry.Total += end - start
					entry.Weight += end - start
				}
				entries[key] = entry
			}
		}
		line, _, err = buff.ReadLine()
	}
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

// ParseMaps gathers all of the file entries from every process
func ParseMaps() EntryMap {

	entries := make(EntryMap)

	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() && pidMatch.MatchString(file.Name()) {
			ParseMap(file.Name(), entries)
		}
	}
	return entries
}

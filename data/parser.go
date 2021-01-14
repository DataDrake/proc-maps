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
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
	pidMatch = regexp.MustCompile(`\d+`)
	rowMatch = regexp.MustCompile(`^(\w+)-(\w+)\s+(\S+)\s+(\d+)\s+(\S+)\s+(\d+)\s*(\S+)?$`)
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

// ParseMap gets all of the entries for a specific map
func ParseMap(pid string, entries EntryMap) {
	maps, err := os.Open("/proc/" + pid + "/maps")
	if err != nil {
		log.Print(err)
	}
	defer maps.Close()
	buff := bufio.NewReader(maps)
	line, _, err := buff.ReadLine()
	// For each line
	for err == nil {
		parseRow(line, entries)
		line, _, err = buff.ReadLine()
	}
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}

func parseRow(line []byte, entries EntryMap) {
	raw := rowMatch.FindStringSubmatch(string(line))
	// IF the line matches and isn't the root Device ID
	if raw != nil && (raw[5] != "00:00") {
		// Generate Key <DeviceID>:<Inode>
		key := raw[5] + ":" + raw[6]
		// If key already exists
		if entry := entries[key]; entry != nil {
			if size := entry.Sizes[raw[3]]; size != nil {
				size.Refs++
				entry.Weight += size.Size
			} else {
				entry.Increment(raw[1], raw[2], raw[3])
			}
		} else {
			// Create new entry
			entry := &FileEntry{
				Name:  raw[7],
				Sizes: make(map[string]*SizeEntry),
			}
			entry.Increment(raw[1], raw[2], raw[3])
			// Store new Entry
			entries[key] = entry
		}
	}
}

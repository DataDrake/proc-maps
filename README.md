# proc-maps
Tool for analyzing the information provided by `/proc/[PID]/maps`

[![Go Report Card](https://goreportcard.com/badge/github.com/DataDrake/proc-maps)](https://goreportcard.com/report/github.com/DataDrake/proc-maps) [![license](https://img.shields.io/github/license/DataDrake/proc-maps.svg)]() 

## Motivation

The `proc` filesystem provides useful insight into the internals of Linux processes. One area of interest is the memory usage of files (e.g. databases, libraries, caches) at run-time. Each running process has a file located at `/proc/[PID]/maps` that catalogs the mapped memory regions in use by the process. A significant portion of these mapped regions will be the files in use by the process. This tool seeks to:

1. Provide easily readable descriptions of file-related memory usage
2. Assist developers in identifying the sources of significant memory usage
3. Allow for different views of the same memory usage statistics

## Building

### Requirements

* Linux-based OS 
* Go (Tested on 1.8.3, but should work for older versions)
* Make
* sudo (cannot be run as a normal user)

### Compile and Install

```bash
make
sudo make install
```

Install optionally supports two arguments:

* `DESTDIR` is an alternative path for packaging
* `PREFIX` defaults to `/usr` but may safely be set to something like `/usr/local`

## Usage

```
USAGE: proc-maps CMD [OPTIONS]

DESCRIPTION: Analyze the contents of /proc/[pid]/maps

COMMANDS:

        file-stats (fs)  : Reads all maps and gathers the stats for a single file
              help (?)   : Get help with a specific subcommand
    rank-all-sizes (ras) : Read all maps and sort files by size
        rank-sizes (rs)  : Read a single map and sort files by size
```

### Examine the whole system

The `rank-all-sizes` subcommand reads through the maps of every running process and consolidates every mapped file into a single table sorted in descending order by size. `Cumulative Size` represents all of the memory used up to and including that file. These sizes are only for the mapped file and do not include allocations made by binary or library execution.

**Command**
```bash
sudo proc-maps ras
sudo proc-maps rank-all-sizes
```

**Example**
```
Rank   Size (B)   Cumulative Size (B)   Filename
1      113M       640M                  /usr/lib64/locale/locale-archive
2      38M        527M                  /usr/lib64/libQt5WebKit.so.5.9.1
3      35M        488M                  /usr/lib64/firefox/browser/omni.ja
4      26M        453M                  /usr/lib64/libicudata.so.59.1
5      20M        426M                  /usr/lib64/libmozjs-38.so
...
```

### Examine a single process

The `rank-sizes` subcommand reads through the map of a single running process and consolidates every mapped file into a single table sorted in descending order by size. `Cumulative Size` represents all of the memory used up to and including that file. These sizes are only for the mapped file and do not include allocations made by binary or library execution.

**Command**
```bash
sudo proc-maps rs 856
sudo proc-maps rank-sizes 856
```

**Example**
```
Rank   Size (B)   Cumulative Size (B)   Filename
1      113M       125M                  /usr/lib64/locale/locale-archive
2      2M         12M                   /usr/lib64/libffi.so.6.0.4
3      1M         9M                    /usr/lib64/libc-2.25.so
4      1M         8M                    /usr/lib64/libgio-2.0.so.0.5400.3
5      1M         6M                    /usr/lib64/libglib-2.0.so.0.5400.3
...
```

### Examine a single file

The `file-stats` subcommand reads through the maps of every running process and consolidates the statistics of a single mapped file into a human-readable format. This view also lists the mapped section of memory according to permissions, providing both a size and number of references (ie. how many processes have a file mapped). `Weight` is an additional characteristic of a file that is calculated as the sum of the products of size and references for each permission set. It may be used as a loose indicator of how widespread the impact of this file is on a system.

**Command**
```bash
sudo proc-maps fs /usr/lib64/libffi.so.6.0.4
sudo proc-maps rile-stats /usr/lib64/libffi.so.6.0.4
```

**Example**
```
File Information:

Name   : /usr/lib64/libffi.so.6.0.4
Size   : 2MB
Weight : 170MB

Permissions   Size (B)   References
r-xp          32K        80
---p          2M         80
r--p          4K         80
rw-p          4K         80
```

### License

Copyright 2018 Bryan T. Meyers <bmeyers@datadrake.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

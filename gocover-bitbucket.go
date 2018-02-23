package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"bytes"
	"strconv"
	"strings"
	"flag"
)

type Coverage struct {
	Files []*FileInfo `json:"files"`
}

type FileInfo struct {
	Path     string	`json:"path"`
	Coverage string `json:"coverage"`
}

type Line struct {
	number	int
	blocks	[]*LineBlock
}

type LineBlock struct {
	startCol	int
	endCol		int
	count		int64
}

func main() {
	prefixPtr := flag.String("prefix", "", "prefix string")
	flag.Parse()
	convert(os.Stdin, os.Stdout, *prefixPtr)
}

func convert(in io.Reader, out io.Writer, prefix string) {
	profiles, err := ParseProfiles(in)
	if err != nil {
		panic("Can't parse profiles")
	}

	coverage := &Coverage{Files: []*FileInfo{}}
	coverage.parseProfiles(profiles, prefix)

	b, err := json.MarshalIndent(coverage, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(out,"%s", string(b))
}

func (cov *Coverage) parseProfiles(profiles []*Profile, prefix string) error {
	for _, profile := range profiles {
		cov.parseProfile(profile, prefix)
	}
	return nil
}

func (cov *Coverage) parseProfile(profile *Profile, prefix string) error {
	fileName := profile.FileName
	var file *FileInfo
	for _, f := range cov.Files {
		if f.Path == fileName {
			file = f
		}
	}
	if file == nil {
		file = &FileInfo{Path: strings.Replace(fileName, prefix, "", -1), Coverage: ""}
		cov.Files = append(cov.Files, file)
	}

	lines := []*Line{}
	for _, b := range profile.Blocks {
		for i := b.StartLine; i <= b.EndLine; i++ {
			var line *Line
			for _, l := range lines {
				if l.number == i {
					line = l
				}
			}
			if line == nil {
				line = &Line{number: i, blocks: []*LineBlock{}}
				lines = append(lines, line)
			}
			newblock := &LineBlock{startCol: 2, endCol: -1, count: int64(b.Count)}
			if line.number == b.StartLine {
				newblock.startCol = b.StartCol
			}
			if line.number == b.EndLine {
				newblock.endCol = b.EndCol
			}
			line.blocks = append(line.blocks, newblock)
		}
	}

	outCovered := bytes.NewBufferString("C:")
	outPartial := bytes.NewBufferString("P:")
	outUncovered := bytes.NewBufferString("U:")
	for _, l := range lines {
		partial := false
		var sum int64 = 0
		mergedBlocks := []*LineBlock{}
		for _, b := range l.blocks {
			var block *LineBlock
			for _, mb := range mergedBlocks {
				if mb.startCol == b.startCol && mb.endCol == b.endCol {
					block = mb
					block.count += b.count
				}
			}
			if block == nil {
				block = &LineBlock{startCol: b.startCol, endCol: b.endCol, count: b.count}
				mergedBlocks = append(mergedBlocks, block)
			}
		}
		for _, b := range mergedBlocks {
			if b.count == 0 {
				partial = true
			}
			sum += b.count
		}
		if sum > 0 {
			if partial {
				outPartial.WriteString(strconv.Itoa(l.number))
				outPartial.WriteString(",")
			} else {
				outCovered.WriteString(strconv.Itoa(l.number))
				outCovered.WriteString(",")

			}
		} else {
			outUncovered.WriteString(strconv.Itoa(l.number))
			outUncovered.WriteString(",")
		}
	}
	file.Coverage = strings.TrimRight(outCovered.String(), ",") + ";" +
						strings.TrimRight(outPartial.String(), ",") + ";" +
						strings.TrimRight(outUncovered.String(), ",")

	return nil
}

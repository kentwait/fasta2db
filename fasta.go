package main

import (
	"strings"
	"bytes"
	"io/ioutil"
)

type sequence struct {
	id string
	description string
	seq string
} 

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadFasta(path string) []sequence {
	dat, err := ioutil.ReadFile(path)
	check(err)

	var (
		ids []string
		descriptions []string
		seqs []string
	)

	var b bytes.Buffer
	for _, line := range bytes.Split(dat, '\n') {
		strLine := string(line)
		if strings.HasPrefix(strLine, ">") {
			splittedStrLine := strings.SplitN(strLine[1:len(strLine) - 1], " ", 2)
			if len(splittedStrLine) > 1 {
				descriptions = append(descriptions, splittedStrLine[1])
			} else {
				descriptions = append(descriptions, "")
			}
			ids = append(ids, splittedStrLine[0])

			if b.Len() > 0 {
				seqs = append(seqs, b.String())
				b.Reset()
			}
		} else {
			b.Write(line)
		}
	}

	var sequences []sequence
	for i := range seqs {
		sequences = append(sequences, 
			sequence{
				id: ids[i],
				description: descriptions[i],
				seq: seqs[i]
			}
		)
	}
	return []sequence
}
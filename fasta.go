package main

import (
	"bytes"
	"io/ioutil"
	"strings"
)

type sequence struct {
	id            string
	description   string
	seq           string
	alignmentName string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFasta(fp, alnID string) []sequence {
	dat, err := ioutil.ReadFile(fp)
	check(err)

	var (
		ids          []string
		descriptions []string
		seqs         []string
	)

	var b bytes.Buffer
	for _, strLine := range strings.Split(string(dat), "\n") {
		if strings.HasPrefix(strLine, ">") {
			splittedStrLine := strings.SplitN(strLine[1:len(strLine)-1], " ", 2)
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
			b.WriteString(strLine)
		}
	}

	var sequences []sequence
	for i := range seqs {
		sequences = append(sequences,
			sequence{
				id:            ids[i],
				description:   descriptions[i],
				seq:           seqs[i],
				alignmentName: alnID,
			},
		)
	}
	return sequences
}

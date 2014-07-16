package main

import (
	"fmt"
	"os"
	"flag"
	"bufio"
	"strings"
	"errors"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	data_dir := "/home/pvh/PycharmProjects/emad-graph-tools/gene_db/data"
 	gene_syn_fn_Ptr := flag.String("synonyms", data_dir + "/gene_synonyms", "file to write synonyms to")
	flag.Parse()
	var input *bufio.Scanner
	var output *bufio.Writer
	if flag.NArg() == 0 {
		input = bufio.NewScanner(os.Stdin)
	} else {
		infile, err := os.Open(flag.Arg(0))
		check(err)
		input = bufio.NewScanner(infile)
		if flag.NArg() == 2 {
			outfile, err := os.Create(flag.Arg(1))
			check(err)
			defer outfile.Close()
			output = bufio.NewWriter(outfile)
		} else {
			output = bufio.NewWriter(os.Stdout)
		}
		defer output.Flush()
	}
	out_syn_file, err := os.Create(*gene_syn_fn_Ptr)
	check(err)
	defer out_syn_file.Close()
	out_syn := bufio.NewWriter(out_syn_file)
	defer out_syn.Flush()
	for tokens := input.Scan(); tokens == true; tokens = input.Scan() {
		line := input.Text()
		if line[0] == '#' {
			continue
		}
		fields := strings.Split(line, "\t")
		if len(fields) != 15 {
			panic(errors.New("invalid line (didn't have 15 fields):\n" + line + "\n"))
		}
		if fields[2] == "NEWENTRY" {
			continue
		}
		if fields[4] != "-" {
			gene_id := fields[1]
			taxon_id := fields[0]
			synonyms := strings.Split(fields[4], "|")
			for _, synonym := range synonyms {
				synonym := strings.TrimSpace(synonym)
				out_line := strings.Join([]string{gene_id, taxon_id,synonym}, "\t") + "\n"
				numbytes, err := out_syn.WriteString(out_line)
				if numbytes != len(out_line) {
					fmt.Fprintf(os.Stderr, "write error")
				}
				check(err)
			}
		}
		trimmed_fields := append(fields[0:4], fields[5:]...)
		output.WriteString(strings.Join(trimmed_fields, "\t") + "\n")
	}

}

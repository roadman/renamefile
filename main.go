package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	//fmt.Println(os.Args, len(os.Args))
	numLen := len(fmt.Sprintf("%d", len(os.Args)))
	var newfile string
	var suffix string

	for k, v := range os.Args[1:] {
		newfile = fmt.Sprintf("%0*d", numLen, k+1)
		suffix = path.Ext(v)
		if suffix != "" {
			newfile += suffix
		}

		fmt.Printf("%s -> %s\n", v, newfile)
		os.Rename(v, newfile)
	}

}

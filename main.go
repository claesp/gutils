package main

import (
	"fmt"
	"os"
)

func count(dir string) int64 {
	var c int64
	c = 0

	ent, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return c
	}

	for _, e := range ent {
		if e.IsDir() {
			c = c + count(fmt.Sprintf("%s/%s", dir, e.Name()))
			continue
		}
		fs, fserr := e.Info()
		if fserr != nil {
			fmt.Fprintf(os.Stderr, "%s\n", fserr)
			continue
		}
		c = c + fs.Size()
	}

	return c
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: missing argment\n", os.Args[0])
		os.Exit(1)
	}

	c := count(os.Args[1])
	fmt.Printf("%d\n", c)
}

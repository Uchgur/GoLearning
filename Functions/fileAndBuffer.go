package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func FileAndBuffer() {
	p := person{
		first: "Jenny",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b) //That's a pointer. Explanation will be soon in this course
	fmt.Println(b.String())
}

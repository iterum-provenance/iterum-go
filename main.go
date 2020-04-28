package main

import (
	"fmt"

	desc "github.com/iterum-provenance/iterum-go/descriptors"
	"github.com/iterum-provenance/iterum-go/util"
)

func main() {
	frag := desc.LocalFragmentDesc{}
	err := util.ReadJSONFile("./build/test.json", &frag)
	fmt.Println(frag)
	fmt.Println(frag.Metadata)
	fmt.Println(err)
}

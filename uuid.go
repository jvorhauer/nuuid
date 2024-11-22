package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/gofrs/uuid/v5"
)

func main() {
	var ver int

	flag.IntVar(&ver, "version", 0, "UUID Version")
	flag.Parse()

	if ver == 0 {
		log.Fatal("--version is manadatory with an integer value between 1 and 7 inclusive")
	}

	res, err := generate(ver)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%v\n", res)
}

func generate(i int) (uuid.UUID, error) {
	switch i {
	case 1:
		return uuid.NewV1()
	case 3:
		return uuid.NewV3(uuid.NamespaceOID, "nuuid"), nil
	case 4:
		return uuid.NewV4()
	case 5:
		return uuid.NewV5(uuid.NamespaceOID, "nuuid"), nil
	case 6:
		return uuid.NewV6()
	case 7:
		return uuid.NewV7()
	}
	return uuid.Nil, errors.New(fmt.Sprintf("ERROR: '%d' is not a valid UUID version for this tool", i))
}

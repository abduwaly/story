package story

import "github.com/teris-io/shortid"

func GenerateId() string {
	str, _ := shortid.Generate()
	return str
}

func GetIdLength() int {
	str, _ := shortid.Generate()
	return len(str)
}


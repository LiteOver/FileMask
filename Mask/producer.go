package Mask

import (
	"bufio"
	"bytes"
	"os"
)

type Producer interface {
	Produce() ([]string, error)
}

type Prod struct {
	Adress string
}

func NewProducer(adress string) *Prod {
	return &Prod{Adress: adress}
}

func (p Prod) Produce() ([]string, error) {
	file, err := os.Open(p.Adress)
	res := make([]string, 0)
	if err != nil {
		return res, err
		os.Exit(1)

	}
	defer file.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	res = append(res, wr.String())
	return res, nil
}

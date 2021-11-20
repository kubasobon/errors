package errors

import (
	"log"
)

func ExampleMaskf() {
	c := func() error {
		return New("If you see this - it's working")
	}

	b := func() error {
		if err := c(); err != nil {
			return Maskf(err, "b error")
		}
		return nil
	}

	a := func() error {
		if err := b(); err != nil {
			return Maskf(err, "a error")
		}
		return nil
	}

	err := a()
	if err != nil {
		log.Fatal(err)
	}
}

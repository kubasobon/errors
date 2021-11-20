package errors

import (
	"log"

	"github.com/kubasobon/errors"
)

func ExampleMaskf() {
	c := func() error {
		return errors.New("If you see this - it's working")
	}

	b := func() error {
		if err := c(); err != nil {
			return errors.Maskf(err, "b error")
		}
		return nil
	}

	a := func() error {
		if err := b(); err != nil {
			return errors.Maskf(err, "a error")
		}
		return nil
	}

	err := a()
	if err != nil {
		log.Fatal(err)
	}
}

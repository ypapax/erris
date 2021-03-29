package main

import (
	"github.com/pkg/errors"
	"log"
)

func main(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var se1 = someError{""}
	var se2 = someError{""}
	log.Println(errors.Is(se1, se2))
	log.Println(errors.Is(se2, se1))
	se1 = someError{"1"}
	se2 = someError{"2"}
	log.Println(errors.Is(se1, se2))
	log.Println(errors.Is(se2, se1))

	se3 := someError{"1"}
	se4 := errors.WithStack(someError{"2"})
	log.Println(errors.Is(se3, se4))
	log.Println(errors.Is(se4, se3))

	se5 := someError{"1"}
	se6 := errors.WithStack(someError{"1"})
	log.Println(errors.Is(se5, se6))
	log.Println(errors.Is(se6, se5))

	se7 := someError{}
	se8 := errors.WithStack(someError{})
	log.Println(errors.Is(se7, se8))
	log.Println(errors.Is(se8, se7))
}

type someError struct {
	field string
}

func (se someError) Error() string {
	return "some error"
}
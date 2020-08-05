package error

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

func getErr() error {
	return xerrors.Errorf("%+v", errors.New("error happens"))
}

func wrap0() error {
	if err := getErr(); err != nil {
		return xerrors.Errorf("wrap_0 : %+v", err)
	}
	return nil
}

func wrap1() error {
	if err := wrap0(); err != nil {
		return xerrors.Errorf("wrap_1 : %+v", err)
	}
	return nil
}

func wrap2() error {
	if err := wrap1(); err != nil {
		return xerrors.Errorf("wrap_2 : %+v", err)
	}
	return nil
}

func print() {
	fmt.Println(wrap2())
	/*
		Ref: https://play.golang.org/p/kSuW1sdRGg1
			wrap_2 : wrap_1 : wrap_0 : error happens:
		            main.getErr
		                /tmp/sandbox485412495/prog.go:11:
		        main.wrap_0
		            /tmp/sandbox485412495/prog.go:16:
		    main.wrap_1
		        /tmp/sandbox485412495/prog.go:23
	*/
}

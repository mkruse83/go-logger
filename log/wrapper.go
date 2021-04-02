package log

import (
	"log"
	"os"
)

type Wrapper interface {
	Print(v ...interface{})
	Fatal(v ...interface{})
}

type WrapperImpl struct {
	Log *log.Logger
}

func newWrapper() *WrapperImpl {
	sl := log.New(os.Stdout, "", 0)
	w := WrapperImpl{Log: sl}
	return &w
}

func (w WrapperImpl) Print(v ...interface{}) {
	w.Log.Print(v...)
}

func (w WrapperImpl) Fatal(v ...interface{}) {
	w.Log.Fatal(v...)
}

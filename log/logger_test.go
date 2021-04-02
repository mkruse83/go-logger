package log

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestStockSignalsLogger_Info(t *testing.T) {
	w := new(FakeWrapper)
	w.On("Print", mock.Anything)
	l := StockSignalsLogger{w, INFO}
	l.Info("Test")
	w.AssertExpectations(t)
}

func TestStockSignalsLogger_Debug(t *testing.T) {
	w := new(FakeWrapper)
	l := StockSignalsLogger{w, INFO}
	l.Debug("Test")
	w.AssertExpectations(t)
}

type FakeWrapper struct {
	mock.Mock
}

func (w *FakeWrapper) Fatal(v ...interface{}) {
	w.Called(v)
}

func (w *FakeWrapper) Print(v ...interface{}) {
	w.Called(v)
}

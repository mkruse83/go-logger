package log

type Level int

const DEBUG = 3
const INFO = 2
const ERROR = 1

type StockSignalsLogger struct {
	log Wrapper
	lvl Level
}


func New(lvl Level) *StockSignalsLogger {
	w := newWrapper()
	ssl := StockSignalsLogger{
		log: w,
		lvl: lvl,
	}
	return &ssl
}

func (l StockSignalsLogger) Info(msg... interface{}) {
	if l.lvl >= INFO {
		l.log.Print(msg)
	}
}

func (l StockSignalsLogger) Debug(msg string) {
	if l.lvl >= DEBUG {
		l.log.Print(msg)
	}
}

func (l StockSignalsLogger) Error(err error) {
	if err != nil {
		l.log.Fatal(err)
	}
}

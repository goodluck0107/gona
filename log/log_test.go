package log

import (
	"log"
	"testing"
)

func TestLog_WithNoSet(t *testing.T) {
	Println("TestLog_WithNoSet Println")
	Debugln("TestLog_WithNoSet Debugln")
	Infoln("TestLog_WithNoSet Infoln")
	Warningln("TestLog_WithNoSet Warningln")
	Errorln("TestLog_WithNoSet Errorln")
}

type CustomLogger struct{}

func (logger *CustomLogger) Printf(format string, v ...interface{}) {

}
func (logger *CustomLogger) Print(v ...interface{}) {

}
func (logger *CustomLogger) Println(v ...interface{}) {
	log.Println("TestLog_WithCustomLogger")

}

func (logger *CustomLogger) Fatalf(format string, v ...interface{}) {

}
func (logger *CustomLogger) Fatal(v ...interface{}) {

}
func (logger *CustomLogger) Fatalln(v ...interface{}) {

}

func (logger *CustomLogger) Panicf(format string, v ...interface{}) {

}
func (logger *CustomLogger) Panic(v ...interface{}) {

}
func (logger *CustomLogger) Panicln(v ...interface{}) {

}
func TestLog_WithCustomLogger(t *testing.T) {
	setLogger(&CustomLogger{})
	Println("TestLog_WithCustomLogger Println")
	Debugln("TestLog_WithCustomLogger Debugln")
	Infoln("TestLog_WithCustomLogger Infoln")
	Warningln("TestLog_WithCustomLogger Warningln")
	Errorln("TestLog_WithCustomLogger Errorln")
}

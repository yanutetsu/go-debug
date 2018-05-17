package debug

import "testing"

func TestPrint(t *testing.T) {
	Print("Print")
}

func TestPrintln(t *testing.T) {
	Println("Println")
}

func TestPrintf(t *testing.T) {
	Printf("%s%d", "hoge", 1)
}

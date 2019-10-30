package duotai_test

import "testing"

type Code string
type Programmer interface {
	WriteCode() Code
}

type GoProgrammer struct {

}

func (g *GoProgrammer)WriteCode() Code  {
	return "fmt.Println(\"hello go\")"
}

type PhpProgrammer struct {

}

func (p PhpProgrammer)WriteCode() Code{
	return "fmt.Println(\"hello php\")"
}

func TestProgrammer(t *testing.T)  {
	gop := new(GoProgrammer)
	php := new(PhpProgrammer)

	t.Log(gop.WriteCode())
	t.Log(php.WriteCode())
}

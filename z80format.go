package z80format

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type instruction struct {
	op         string
	twoOp bool
	arg1       opInterface
	arg2       opInterface
}

type opInterface interface {
	Offset() bool
	Op() []string
}

type op struct {
	offset bool
	op     []string
}

func (o op) Offset() bool {
	return o.offset
}
func (o op) Op() []string {
	return o.op
}

var (
	noOp                = op{}
	op8Instructions     = []string{"A", "H", "L", "D", "E", "B", "C", "I", "R", "IXH", "IYH", "IXL", "IYL", "(HL)", "(DE)", "(A)", "(H)", "(L)", "(D)", "(E)", "(B)", "(C)"}
	op8FullInstructions = []string{"A", "H", "L", "D", "E", "B", "C", "I", "R", "IXH", "IYH", "IXL", "IYL", "(HL)", "(DE)", "(A)", "(H)", "(L)", "(D)", "(E)", "(B)", "(C)", "(IX+n)", "(IY+n)", "(nn)"}
	op16Instructions    = []string{"HL", "BC", "DE", "(SP)", "AF", "AF'", "HL'", "BC'", "DE'", "IX", "IY", "IX'", "IY'"}
	instructions        = map[string]instruction{
		instruction{op:"ADC", twoOp: true, }
	}
)

/*

var op8 = []string{"A", "H", "L", "D", "E", "B", "C", "I", "R", "IXH", "IYH", "IXL", "IYL", "(HL)", "(DE)", "(A)", "(H)", "(L)", "(D)", "(E)", "(B)", "(C)"}
var op8Iter = []string{"(IX+n)", "(IY+n)", "(nn)"}
var op16 = []string{"HL", "BC", "DE", "(SP)", "AF", "AF'", "HL'", "BC'", "DE'", "IX", "IY", "IX'", "IY'"}
var op16Full = [][]string{op16}
var op8Full = [][]string{op8, op8Add}

var instructions = map[string][][]string{
	"ADC":  noOp,
	"ADD":  noOp,
	"AND":  noOp,
	"BIT":  noOp,
	"CALL": noOp,
	"CCF":  noOp,
	"CP":   noOp,
	"CPD":  noOp,
	"CPDR": [][]string{},
	"CPI":  [][]string{},
	"CPIR": [][]string{},
	"CPL":  [][]string{},
	"DAA":  [][]string{},
	"DEC":  [][]string{},
	"DI":   [][]string{},
	"DJNZ": [][]string{},
	"EI":   [][]string{},
	"EX":   [][]string{},
	"EXX":  noOp,
	"HALT": noOp,
	"IM":   noOp,
	"IN":   noOp,
	"INC":  noOp,
	"IND":  noOp,
	"INDR",
	"INI",
	"INIR",
	"JP",
	"JR",
	"LD",
	"LDD",
	"LDDR",
	"LDI",
	"LDIR",
	"NEG",
	"NOP",
	"OR",
	"OTDR",
	"OTIR",
	"OUT",
	"OUTD",
	"OUTI",
	"POP",
	"PUSH",
	"RES",
	"RET",
	"RETI",
	"RETN",
	"RL",
	"RLA",
	"RLC",
	"RLCA",
	"RLD",
	"RR",
	"RRA",
	"RRC",
	"RRCA",
	"RRD",
	"RST",
	"SBC",
	"SCF",
	"SET",
	"SLA",
	"SLL/SL1",
	"SRA",
	"SRL",
	"SUB",
	"XOR",
}
*/

func Format(in string) (string, error) {
	out := new(bytes.Buffer)
	buf := bytes.NewBufferString(in)
	r := bufio.NewReader(buf)
	scanner := bufio.NewScanner(r)
	line := 1
	for scanner.Scan() {
		t := scanner.Text()
		insts := strings.Split(t, ":")
		for _, v0 := range insts {
			cleaned := strings.Trim(v0, " \t")
			instr := strings.FieldsFunc(cleaned, split)
			if len(instr) == 1 || len(instr) == 0 {
				out.WriteString(cleaned)
			} else {
				v1 := strings.ToUpper(instr[0])
				op := strings.Join(instr[1:], " ")
				for _, v2 := range instructions {
					if v1 == v2 {
						out.WriteString(fmt.Sprintf("\t%s %s", v1, op))
						break
					}
				}
			}
			if len(insts) > 1 {
				out.WriteString(":")
			}
		}
		out.WriteString("\n")
		line++
	}
	return out.String(), nil
}

func split(r rune) bool {
	if r == '\t' || r == ' ' {
		return true
	}
	return false
}

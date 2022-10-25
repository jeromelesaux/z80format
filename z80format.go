package z80format

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

var instructions = []string{
	"ADC",
	"ADD",
	"AND",
	"BIT",
	"CALL",
	"CCF",
	"CP",
	"CPD",
	"CPDR",
	"CPI",
	"CPIR",
	"CPL",
	"DAA",
	"DEC",
	"DI",
	"DJNZ",
	"EI",
	"EX",
	"EXX",
	"HALT",
	"IM",
	"IN",
	"INC",
	"IND",
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

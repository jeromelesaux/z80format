package z80format

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type instruction struct {
	op       string
	operands []operand
}

func (i *instruction) hasOperands() bool {
	return len(i.operands) > 0
}

type operand struct {
	OperandLeft  []string
	OperandRight []string
}

func (o *operand) hasTwoArguments() bool {
	return !reflect.DeepEqual(o.OperandLeft, noOp)
}

func (o *operand) isCondition() bool {
	return reflect.DeepEqual(o.OperandLeft, conditions)
}

var (
	noOp                = []string{}
	op8Instructions     = []string{"A", "H", "L", "D", "E", "B", "C", "I", "R", "IXH", "IYH", "IXL", "IYL", "(HL)", "(DE)", "(A)", "(H)", "(L)", "(D)", "(E)", "(B)", "(C)"}
	op8FullInstructions = []string{"A", "H", "L", "D", "E", "B", "C", "I", "R", "IXH", "IYH", "IXL", "IYL", "(HL)", "(DE)", "(A)", "(H)", "(L)", "(D)", "(E)", "(B)", "(C)", "(IX+n)", "(IY+n)", "(nn)"}
	op16Instructions    = []string{"HL", "BC", "DE", "(SP)", "AF", "AF'", "HL'", "BC'", "DE'", "IX", "IY", "IX'", "IY'"}
	conditions          = []string{"", "NZ", "Z", "NC", "C", "PO", "PE", "P", "M"}
	instructions        = map[string]instruction{
		"ADC": {
			op: "ADC",
			operands: []operand{
				{OperandLeft: []string{"A"}, OperandRight: op8Instructions},
				{OperandLeft: []string{"HL"}, OperandRight: op16Instructions},
			}},
		"ADD": {
			op: "ADD",
			operands: []operand{
				{OperandLeft: []string{"A"}, OperandRight: op8Instructions},
				{OperandLeft: op16Instructions, OperandRight: op16Instructions},
			}},
		"AND": {
			op: "AND",
			operands: []operand{
				{OperandLeft: op8Instructions, OperandRight: op8Instructions},
			}},
		"BIT": {
			op: "BIT",
			operands: []operand{
				{OperandLeft: noOp, OperandRight: op8Instructions},
			}},
		"CALL": {
			op: "CALL",
			operands: []operand{
				{OperandLeft: conditions, OperandRight: noOp},
			}},
		"CCF": {
			op:       "CCF",
			operands: []operand{}},
		"CP": {
			op: "CP",
			operands: []operand{
				{OperandLeft: op8Instructions, OperandRight: noOp},
			}},
		"CPD": {
			op:       "CPD",
			operands: []operand{}},
		"CPDR": {
			op:       "CPDR",
			operands: []operand{}},
		"CPI": {
			op:       "CPI",
			operands: []operand{}},
		"CPIR": {
			op:       "CPIR",
			operands: []operand{}},
		"CPL": {
			op:       "CPL",
			operands: []operand{}},
		"DAA": {
			op:       "DAA",
			operands: []operand{}},
		"DEC": {
			op: "DEC",
			operands: []operand{
				{OperandLeft: op8Instructions, OperandRight: noOp},
				{OperandLeft: op16Instructions, OperandRight: noOp},
			}},
		"DI": {
			op:       "DI",
			operands: []operand{}},
		"DJNZ": {
			op:       "DJNZ",
			operands: []operand{}},
		"EI": {
			op:       "EI",
			operands: []operand{}},
		"EX": {
			op: "EX",
			operands: []operand{
				{OperandLeft: op16Instructions, OperandRight: op16Instructions},
			}},
		"EXX": {
			op:       "EXX",
			operands: []operand{}},
		"HALT": {
			op:       "HALT",
			operands: []operand{}},
		"IM": {
			op:       "IM",
			operands: []operand{}},
		"IN": {
			op:       "IN",
			operands: []operand{{OperandLeft: []string{"A"}, OperandRight: []string{"(B)", "(C)", "(D)", "(E)", "(H)", "(L)"}}}},
		"INC": {
			op: "INC",
			operands: []operand{
				{OperandLeft: op8Instructions, OperandRight: noOp},
				{OperandLeft: op16Instructions, OperandRight: noOp},
			}},
		"IND": {
			op:       "IND",
			operands: []operand{}},
		"INDR": {
			op:       "INDR",
			operands: []operand{}},
		"INI": {
			op:       "INI",
			operands: []operand{}},
		"INIR": {
			op:       "INIR",
			operands: []operand{}},
		"JP": {
			op: "JP",
			operands: []operand{
				{OperandLeft: conditions, OperandRight: noOp},
				{OperandLeft: op16Instructions, OperandRight: noOp},
				{OperandLeft: noOp, OperandRight: noOp},
			}},
		"JR": {
			op: "JR",
			operands: []operand{
				{OperandLeft: conditions, OperandRight: noOp},
			}},
		"LD": {
			op: "LD",
			operands: []operand{
				{OperandLeft: op8FullInstructions, OperandRight: noOp},
				{OperandLeft: op16Instructions, OperandRight: noOp},
			}},
		"LDD": {
			op:       "LDD",
			operands: []operand{}},
		"LDDR": {
			op:       "LDDR",
			operands: []operand{}},
		"LDI": {
			op:       "LDI",
			operands: []operand{}},
		"LDIR": {
			op:       "LDIR",
			operands: []operand{}},
		"NEG": {
			op:       "NEG",
			operands: []operand{}},
		"NOP": {
			op:       "NOP",
			operands: []operand{}},
		"OR": {
			op: "OR",
			operands: []operand{
				{OperandLeft: op8Instructions, OperandRight: noOp}}},
		"OTDR": {
			op:       "OTDR",
			operands: []operand{}},
		"OTIR": {
			op:       "OTIR",
			operands: []operand{}},
		"OUT": {
			op:       "OUT",
			operands: []operand{{OperandLeft: []string{"(B)", "(C)", "(D)", "(E)", "(H)", "(L)"}, OperandRight: op8FullInstructions}}},
		"OUTD": {
			op:       "OUTD",
			operands: []operand{}},
		"OUTI": {
			op:       "OUTI",
			operands: []operand{}},
		"POP": {
			op:       "POP",
			operands: []operand{{OperandLeft: op16Instructions, OperandRight: noOp}}},
		"PUSH": {
			op:       "PUSH",
			operands: []operand{{OperandLeft: op16Instructions, OperandRight: noOp}}},
		"RES": {
			op:       "RES",
			operands: []operand{{OperandLeft: noOp, OperandRight: op8Instructions}}},
		"RET": {
			op:       "RET",
			operands: []operand{{OperandLeft: conditions, OperandRight: noOp}}},
		"RETI": {
			op:       "RETI",
			operands: []operand{}},
		"RETN": {
			op:       "RETN",
			operands: []operand{}},
		"RL": {
			op:       "REL",
			operands: []operand{{OperandLeft: op8Instructions, OperandRight: noOp}}},
		"RLA": {
			op:       "RLA",
			operands: []operand{}},
		"RLC": {
			op:       "RLC",
			operands: []operand{{OperandLeft: op8Instructions, OperandRight: noOp}}},
		"RLCA": {
			op:       "RLCA",
			operands: []operand{}},
		"RLD": {
			op:       "RLD",
			operands: []operand{}},
		"RR": {
			op:       "RR",
			operands: []operand{{OperandLeft: op8Instructions, OperandRight: noOp}}},
		"RRA": {
			op:       "RRA",
			operands: []operand{}},
		"RRC": {
			op:       "RRC",
			operands: []operand{{OperandLeft: op8Instructions, OperandRight: noOp}}},
		"RRCA": {
			op:       "RRCA",
			operands: []operand{}},
		"RRD": {
			op:       "RRD",
			operands: []operand{}},
		"RST": {
			op:       "RRCA",
			operands: []operand{{OperandLeft: noOp, OperandRight: noOp}}},
		"SDC": {
			op: "SDC",
			operands: []operand{
				{OperandLeft: []string{"A"}, OperandRight: op8Instructions},
				{OperandLeft: []string{"HL"}, OperandRight: op16Instructions},
			}},
		"SCF": {
			op:       "SCF",
			operands: []operand{}},
		"SET": {
			op:       "SET",
			operands: []operand{{OperandLeft: noOp, OperandRight: op8Instructions}}},
		"SLA": {
			op:       "SLA",
			operands: []operand{{OperandLeft: op8Instructions, OperandRight: noOp}}},
		"SRA": {
			op:       "SRA",
			operands: []operand{{OperandLeft: op8Instructions, OperandRight: noOp}}},
		"SRL": {
			op:       "SRL",
			operands: []operand{{OperandLeft: op8Instructions, OperandRight: noOp}}},
		"SUB": {
			op:       "SUB",
			operands: []operand{{OperandLeft: op8Instructions, OperandRight: noOp}}},
		"XOR": {
			op:       "XOR",
			operands: []operand{{OperandLeft: op8Instructions, OperandRight: noOp}}},
	}
)

func Format(r io.Reader) (string, error) {
	out := new(bytes.Buffer)
	scanner := bufio.NewScanner(r)
	line := 1
	for scanner.Scan() {
		t := scanner.Text()
		insts := strings.Split(t, ":")
		for indice, v0 := range insts {
			cleaned := strings.Trim(v0, " \t")
			instr := strings.FieldsFunc(cleaned, split)
			if len(instr) == 0 {
				out.WriteString(cleaned)
			} else {
				v1 := strings.ToUpper(instr[0])
				i, ok := instructions[v1]
				if ok {
					if i.hasOperands() {
						for _, op := range i.operands {
							if !op.hasTwoArguments() {
								out.WriteString(fmt.Sprintf("\t%s %s", v1, instr[1]))
							} else {
								if op.isCondition() {
									out.WriteString(fmt.Sprintf("\t%s", v1))
								} else {
									mOp := strings.Split(instr[1], ",")
									if len(mOp) == 2 {
										opLeft := mOp[0]
										opRight := mOp[1]
										ok, val0 := contains(opLeft, op.OperandLeft)
										if !ok {
											continue
										}
										ok, val1 := contains(opRight, op.OperandRight)
										if !ok {
											continue
										}
										out.WriteString(fmt.Sprintf("\t%s %s,%s %s", v1, val0, val1, strings.Join(instr[2:], " ")))
									}
								}
							}
							// verifier les syntaxes a la suite de l'iteration
						}
					} else {
						// verifier la syntaxe
						out.WriteString(fmt.Sprintf("\t%s", v1))
					}
				} else {
					// label
					out.WriteString(strings.Join(instr, " "))
				}
			}
			if len(insts) > 1 && indice < (len(insts)-1) {
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

func contains(s string, collection []string) (bool, string) {
	if reflect.DeepEqual(collection, noOp) {
		return true, s
	}
	for _, v := range collection {
		if v == strings.ToUpper(s) {
			return true, strings.ToUpper(s)
		}
	}
	return false, ""
}

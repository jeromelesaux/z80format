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
	return !reflect.DeepEqual(o.OperandRight, noOp)
}

func (o *operand) isCondition() bool {
	return reflect.DeepEqual(o.OperandLeft, conditions)
}

func IsCondition(value string) bool {
	ok, _ := contains(strings.ToUpper(value), conditions)
	return ok
}

func IsRegister8Bits(value string) bool {
	ok, _ := contains(strings.ToUpper(value), op8FullInstructions)
	return ok
}

func IsRegister16Bits(value string) bool {
	ok, _ := contains(strings.ToUpper(value), op16Instructions)
	return ok
}

func IsLabel(value string) bool {
	_, ok := instructions[value]
	if ok {
		return false
	}
	if IsCondition(value) {
		return false
	}
	if IsRegister16Bits(value) {
		return false
	}
	if IsRegister8Bits(value) {
		return false
	}
	if IsHexadecimal(value) {
		return false
	}

	return true
}

func IsHexadecimal(value string) bool {
	var i int
	res, err := fmt.Sscanf(strings.ToUpper(value), "&%X", &i)
	if err == nil && res == 1 {
		return true
	}
	res, err = fmt.Sscanf(strings.ToUpper(value), "#%X", &i)
	if err == nil && res == 1 {
		return true
	}
	return false
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
				{OperandLeft: noOp, OperandRight: noOp},
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
				var label, comment, instruction, arg1, arg2 string
				instruction = strings.ToUpper(instr[0])
				if len(instr) > 1 {
					arg1 = instr[1]
				}
				if len(instr) > 2 {
					arg2 = instr[2]
				}
				if len(instr) > 3 {
					comment = strings.Join(instr[3:], " ")
				}
				if IsLabel(instruction) {
					label = instr[0]
					instruction = strings.ToUpper(arg1)
					arg1 = arg2
					if len(instr) > 3 {
						arg2 = instr[3]
					}
					if len(instr) > 4 {
						comment = strings.Join(instr[4:], " ")
					} else {
						comment = ""
					}

				}

				i, ok := instructions[instruction]
				if ok {
					if i.hasOperands() {
						for _, op := range i.operands {
							if !op.hasTwoArguments() {
								if (IsLabel(arg1) && reflect.DeepEqual(op.OperandRight, noOp)) ||
									(IsHexadecimal(arg1) && reflect.DeepEqual(op.OperandRight, noOp)) ||
									(IsCondition(arg1) && reflect.DeepEqual(op.OperandLeft, conditions)) ||
									(IsRegister8Bits(arg1) && reflect.DeepEqual(op.OperandLeft, op8Instructions)) ||
									(IsRegister16Bits(arg1) && reflect.DeepEqual(op.OperandLeft, op16Instructions)) {
									if label != "" {
										out.WriteString(label)
									}
									out.WriteString(fmt.Sprintf("\t%s ", instruction))
									if IsHexadecimal(arg1) || IsLabel(arg1) {
										out.WriteString(arg1)
									} else {
										out.WriteString(strings.ToUpper(arg1))
									}
									if arg2 != "" {
										if IsHexadecimal(arg2) || IsLabel(arg2) {
											out.WriteString(fmt.Sprintf(",%s", arg2))
										} else {
											out.WriteString(fmt.Sprintf(",%s", strings.ToUpper(arg2)))
										}
									}
									if comment != "" {
										out.WriteString(comment)
									}
									break
								}

							} else {
								var conditionValue string
								if IsCondition(arg1) {
									conditionValue = arg1
								}
								ok, condition := contains(conditionValue, op.OperandLeft)
								if op.isCondition() && ok {
									out.WriteString(fmt.Sprintf("\t%s %s", instruction, condition))
									if label != "" {
										out.WriteString(label)
									}
									out.WriteString(fmt.Sprintf("\t%s", instruction))
									out.WriteString(condition)
									if arg2 != "" {
										out.WriteString(fmt.Sprintf(",%s", arg2))
									}
									if comment != "" {
										out.WriteString(comment)
									}
									break
								} else {

									ok, val0 := contains(arg1, op.OperandLeft)
									if !ok {
										continue
									}
									ok, val1 := contains(arg2, op.OperandRight)
									if !ok {
										continue
									}
									if label != "" {
										out.WriteString(label)
									}
									out.WriteString(fmt.Sprintf("\t%s", instruction))
									if val0 != "" {
										out.WriteString(fmt.Sprintf(" %s", val0))
									}
									if val1 != "" {
										out.WriteString(fmt.Sprintf(",%s", val1))
									}
									if comment != "" {
										out.WriteString(comment)
									}
									break

								}
							}
							// verifier les syntaxes a la suite de l'iteration
						}
					} else {
						// verifier la syntaxe
						out.WriteString(fmt.Sprintf("\t%s", instruction))
						if arg1 != "" {
							out.WriteString(" " + arg1)
						}
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
	if r == '\t' || r == ' ' || r == ',' {
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

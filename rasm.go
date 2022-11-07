package z80format

import "strings"

var (
	rasmWords = []string{
		"INCBIN",
		"INCLUDE",
		"DW",
		"DB",
		"RUN",
		"ORG",
		"SAVE",
		"ALIGN",
		"LIMIT",
		"PROTECT",
		"CONFINE",
		"DEFB",
		"DM",
		"DEFM",
		"DEFI",
		"DEFS",
		"STR",
		"CHARSET",
		"ASSERT",
		"IF",
		"IFNOT",
		"UNDEF",
		"IFUSED",
		"IFNUSED",
		"REPEAT",
		"WHILE",
		"WHEND",
		"STRUCT",
		"SIZEOF",
		"BANK",
		"PAGE",
		"PAGESET",
		"BUILDSNA",
		"SETCPC",
		"SETCRTC",
		"SETSNA",
		"BANK",
		"BANKSET",
		"BRK",
		"GET_R",
		"GET_G",
		"GET_B",
		"SET_R",
		"SET_G",
		"SET_B",
	}

	instructionsToReplace = map[string]string{
		"BYTE": "DB",
		"WORD": "DW",
	}
)

func RasmSyntaxe(in string) string {
	in = strings.Replace(in, "&", "#", -1)
	for _, v := range rasmWords {
		var previous int
		for {
			i := strings.Index(strings.ToUpper(in[previous:]), v)
			if i >= 0 {
				old := in[(i + previous):(i + previous + len(v))]
				in = strings.Replace(in, old, v, 1)
				previous += i + len(v)
			} else {
				break
			}

		}
	}

	for k, v := range instructionsToReplace {
		var previous int
		for {
			i := strings.Index(strings.ToUpper(in[previous:]), k)
			if i >= 0 {
				old := in[(i + previous):(i + previous + len(k))]
				in = strings.Replace(in, old, v, 1)
				previous += i + len(v)
			} else {
				break
			}
		}
	}
	return in
}

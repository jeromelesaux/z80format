package z80format_test

import (
	"testing"

	"github.com/jeromelesaux/z80format"
	"github.com/stretchr/testify/assert"
)

func TestReplaceByteRasmSyntaxe(t *testing.T) {
	code := `BYTE &12`
	res := z80format.RasmSyntaxe(code)
	assert.Equal(t, res, "DB #12")
}

func TestFormatInludeWordRasmSyntaxe(t *testing.T) {
	code := `Include 'monfichier.asm'`
	res := z80format.RasmSyntaxe(code)
	assert.Equal(t, res, "INCLUDE 'monfichier.asm'")
}

func TestFormatInbinWordsRasmSyntaxe(t *testing.T) {
	code := `Incbin '1.bin': incbin '2.bin'`
	res := z80format.RasmSyntaxe(code)
	assert.Equal(t, res, "INCBIN '1.bin': INCBIN '2.bin'")
}

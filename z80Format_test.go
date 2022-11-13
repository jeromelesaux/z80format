package z80format_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jeromelesaux/z80format"
	"github.com/stretchr/testify/assert"
)

func TestLdWithLabelComment(t *testing.T) {
	code := `LD	de,my_font ; SpriteHardPtr`
	expected := "\tLD DE,my_font; SpriteHardPtr\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestIncB(t *testing.T) {
	code := `inc b`
	expected := "\t" + strings.ToUpper(code) + "\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestCallWithLabel(t *testing.T) {
	code := `call #BC07`
	expected := "\t" + strings.ToUpper(code) + "\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestOut(t *testing.T) {
	code := `out (c),a`
	expected := "\t" + strings.ToUpper(code) + "\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestLd16Registers(t *testing.T) {
	code := `ld bc,DE`
	expected := "\t" + strings.ToUpper(code) + "\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestLd16RegistersWithLabelComment(t *testing.T) {
	code := `ptr ld bc,DE ; comment`
	expected := "ptr\tLD BC,DE; comment\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestLd16RegistersWithLabelComments(t *testing.T) {
	code := `ptr ld bc,DE ; comment  comment1 comment2`
	expected := "ptr\tLD BC,DE; comment comment1 comment2\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestLdWithLabel(t *testing.T) {
	code := `space     ld bc,#F40E`
	expected := "space\tLD BC,#F40E\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestJumpWithConditionComment(t *testing.T) {
	code := `jr nc, WaitVbl; my comment`
	expected := "\tJR NC,WaitVbl;my comment\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestIm1(t *testing.T) {
	code := `im 1`
	expected := "\tIM 1\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestIsHexadecimal(t *testing.T) {
	res := z80format.IsHexadecimal("#BC77")
	assert.Equal(t, true, res)

	res = z80format.IsHexadecimal("#bc77")
	assert.Equal(t, true, res)

	res = z80format.IsHexadecimal("&77")
	assert.Equal(t, true, res)
}

func TestData(t *testing.T) {
	code := `def 1,2,3,4,5`
	expected := "def 1,2,3,4,5\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestRasmSaveKeyword(t *testing.T) {
	code := `save'disc.bin',#200, end - start,DSK,'martine-animate.dsk'`
	expected := "save'disc.bin',#200, end - start,DSK,'martine-animate.dsk'\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

package z80format_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jeromelesaux/z80format"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	code := `FontOk
	;brk
		LD	de,my_font ; SpriteHardPtr
		ld l,a
		ld h,0
		add hl,hl ;*2
		add hl,hl ;*4
		add hl,hl ;*8
		add hl,hl ;*16
		add hl,hl ;*32
		add hl,hl ;*64
		add hl,hl ;*128
		add hl,hl ;*256 octets taille d'une sprite hard
		add hl,de         ; hl pointe sur la bonne lettre dans la fonte
	
		LD	A,I				; numero du sprite
		ADD	A,#40				; sprites sont en #4000 dans l'asic
		LD	D,A				; adresse du sprite
		LD	E,0
		ld bc,#00FF+1
		ldir  
			 jp asicoff
`
	expected := "FontOk\n;brk\n\tLD de,my_font; SpriteHardPtr\n\tLD l,a\n\tLD h,0\n\tADD HL,HL;*2\n\tADD HL,HL;*4\n\tADD HL,HL;*8\n\tADD HL,HL;*16\n\tADD HL,HL;*32\n\tADD HL,HL;*64\n\tADD HL,HL;*128\n\tADD HL,HL;*256 octets taille d'une sprite hard\n\tADD HL,DE; hl pointe sur la bonne lettre dans la fonte\n\n\tLD A,I; numero du sprite\n\n\tLD D,A; adresse du sprite\n\tLD E,0\n\tLD bc,#00FF+1\n\tLDIR\n\tJP asicoff\n"
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

func TestLdWithLabel(t *testing.T) {
	code := `space     ld bc,#F40E`
	expected := "space\tLD bc,#F40E\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}
func TestSampleSpaceTestRoutine(t *testing.T) {
	code := `
	;
	; Joue Musique
	;
	TestSpace
			 ; call #0244 ; play music
			 LD	B,#F5
	WaitVBL
		IN	A,(C)
		RRA
		JR	NC,WaitVBL			; Attendre VBL
		  call #0244 ; play music
	;
	space     ld bc,#F40E
			  out (c),c
			  ld bc,#F6C0
			  out (c),c
		  DW #71ED        ; out (c),0
			  ld bc,#F792
			  out (c),c
			  ld bc,#F645
			  out (c),c
			  ld b,#F4
			  in a,(c)
			  ld bc,#F782
			  out (c),c
			  bit 7,a
			  jp nz,LoopScroll
	`
	expected := "\n;\n; Joue Musique\n;\nTestSpace\n; call #0244 ; play music\n\tLD B,#F5\nWaitVBL\n\tIN A,(C)\n\tRRA\n\tJR NC,WaitVBL; Attendre VBL\n\tCALL #0244; play music\n;\nspace\tLD bc,#F40E\n\tOUT (C),C\n\tLD bc,#F6C0\n\tOUT (C),C\nDW #71ED ; out (c),0\n\tLD bc,#F792\n\tOUT (C),C\n\tLD bc,#F645\n\tOUT (C),C\n\tLD b,#F4\n\tIN A,(C)\n\tLD bc,#F782\n\tOUT (C),C\n\tBIT 7,A\n\tJP nz,LoopScroll\n\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

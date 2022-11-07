package z80format_test

import (
	"bytes"
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
	expected := "FontOk\n;brk\n\tLD DE,my_font ; SpriteHardPtr\n\tLD L,a \n\tLD H,0 \n\tADD HL,HL ;*2\n\tADD HL,HL ;*4\n\tADD HL,HL ;*8\n\tADD HL,HL ;*16\n\tADD HL,HL ;*32\n\tADD HL,HL ;*64\n\tADD HL,HL ;*128\n\tADD HL,HL ;*256 octets taille d'une sprite hard\n\tADD HL,DE ; hl pointe sur la bonne lettre dans la fonte\n\n\tLD A,I ; numero du sprite\n\n\tLD D,A ; adresse du sprite\n\tLD E,0 \n\tLD BC,#00FF+1 \n\tLDIR\n\tJP asicoff\n"
	res, err := z80format.Format(bytes.NewBufferString(code))
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

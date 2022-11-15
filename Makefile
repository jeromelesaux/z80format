CC=go
RM=rm
MV=mv


build:
	${CC} build -o z80formatter cli/z80formatter/main.go

clean:
	${RM} -f z80formatter

test:
	${CC} test ./... -cover
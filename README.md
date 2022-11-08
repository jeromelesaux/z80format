# z80format


Z80Format is a simple assembly code source formatter.
It may convert the assembly code source to a ready [Rasm](https://github.com/EdouardBERGE/rasm) code source.
It's free to use.

## Installation 
You need a Golang environnement set up and compile it via the makefile command : 
```bash 
make 
```

## Usage 

You can use it like this 
```bash 
    z80formatter -format myfile.asm > myformattedfile.asm
```

or 

```bash 
    z80formatter < myfile.asm > myformatterfile.asm
```


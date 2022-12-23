package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/davidn5013/aoc/utl"
)

// instruction to cpu in a com system
type Aoc struct {
	instrList []instruction // list of instruction from input file
	cycles    int           // com systgem cpu cycles
	// addTocycle   int		// cycle for this instruction
	beforecycles int          // stores cycle before instruction
	register     Reg          // cpu register x ...
	sumcounter   map[int]bool // stop and summarize on index
	scanline     string       // Visual presention for solution 2 screen
}

type instruction struct {
	cpucmd string // Instruction noop and addx
	arg    int    // argument for addx
}

type Reg struct {
	x int // place for register
}

// NewAoc Aoc struct for cpu information
func NewAoc() *Aoc {
	a := Aoc{
		register: Reg{x: 1},
	}

	a.sumcounter = make(map[int]bool)
	a.sumcounter = map[int]bool{
		20:  false,
		60:  false,
		100: false,
		140: false,
		180: false,
		220: false,
	}

	return &a

}

func (a *Aoc) noop() {
	a.DoCycle()
}

func (a *Aoc) addx(i int) {
	a.DoCycle()
	a.DoCycle()

	a.register.x += i
}

func (a *Aoc) DoCycle() {

	if a.cycles > 40*6 {
		return
	}

	if a.cycles%40 == 0 && a.cycles <= 220 {
		a.scanline += "\n"
	}
	if a.cycles%40 >= a.register.x-1 && a.cycles%40 <= a.register.x+1 {
		a.scanline += "##"
	} else {
		a.scanline += ".."
	}

	a.cycles++
}

func (a *Aoc) parsefile(filename string) {
	startmsg(utl.CurrFuncName())

	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	sc := bufio.NewScanner(fp)

	for sc.Scan() {
		var cmd string
		var cmdarg int
		fmt.Sscanf(sc.Text(), "%s %d", &cmd, &cmdarg)
		a.instrList = append(a.instrList, instruction{cpucmd: cmd, arg: cmdarg})
	}

}

func (a *Aoc) instructionRunner(solnr int) (ret int) {
	for nr, instr := range a.instrList {
		a.beforecycles = a.cycles
		switch instr.cpucmd {
		case "noop":
			a.noop()
		case "addx":
			a.addx(instr.arg)
		default:
			log.Fatalf("unknow command in file %s\n", instr.cpucmd)
		}

		// Debug
		if debuglvl >= 10 && nr < 10 {
			utl.Debug(true, "Cycle %d Instr %s Instrarg %d register.x %d\n", a.cycles, instr.cpucmd, instr.arg, a.register.x)
		}

		// Interesting cycles i map sumcounter
		for i := range a.sumcounter {
			if a.beforecycles <= i && a.cycles >= i && a.sumcounter[i] == false {
				a.sumcounter[i] = true
				fmt.Printf("-\tDuring the %dth cycle start on cycle %d register X=%d\n\tStrength is %[1]d * %[3]d = %d\n\n", i, a.cycles, a.register.x, i*a.register.x)
				ret += i * a.register.x
			}
		}
	}

	if solnr == 2 {
		fmt.Println(a.scanline)
	}

	return ret

}

func (a *Aoc) sol1() (ret int) {
	startmsg(utl.CurrFuncName())

	return a.instructionRunner(1)

}

func (a *Aoc) sol2() (ret int) {
	startmsg(utl.CurrFuncName())

	return a.instructionRunner(2)

}

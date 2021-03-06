package vm

import (
	"fmt"
	"os"
	"text/template"
	"xoba/ad/defs"
)

const (
	asm_source = "asm.go"
	vm_source  = "execute.go"
)

func OpForName(name string) VmOp {
	for _, op := range AllOps {
		if name == op.String() {
			return op
		}
	}
	panic("illegal name: " + name)
}

func DerivativeOpForName(base string, index int) VmOp {
	name := fmt.Sprintf("D_%s_D%d", base, index)
	for _, op := range AllOps {
		if name == op.String() {
			return op
		}
	}
	panic("illegal name: " + name)
}

func GenVm(args []string) {

	twos := make(map[VmOp]string)
	threes := make(map[VmOp]string)
	twoArgFuncs := make(map[VmOp]string)

	for _, d := range Defs {
		op := OpForName(d.Name)
		switch d.Type {
		case "twos":
			twos[op] = d.Runtime
			dOp0 := DerivativeOpForName(d.Name, 0)
			twos[dOp0] = dOp0.ToLower()
		case "threes":
			threes[op] = d.Runtime
			dOp0 := DerivativeOpForName(d.Name, 0)
			twoArgFuncs[dOp0] = dOp0.ToLower()
			dOp1 := DerivativeOpForName(d.Name, 1)
			twoArgFuncs[dOp1] = dOp1.ToLower()
		case "funcs2":
			twoArgFuncs[op] = d.Runtime
			dOp0 := DerivativeOpForName(d.Name, 0)
			twoArgFuncs[dOp0] = dOp0.ToLower()
			dOp1 := DerivativeOpForName(d.Name, 1)
			twoArgFuncs[dOp1] = dOp1.ToLower()
		}
	}

	gen := func(name, src string) {
		f, err := os.Create(name)
		check(err)
		t := template.Must(template.New(vm_source).Parse("// autogenerated, do not edit!\npackage vm\n" + src))
		t.Execute(f, map[string]interface{}{
			"twos":   twos,
			"threes": threes,
			"funcs2": twoArgFuncs,
		})
		f.Close()
		defs.Gofmt(name)
	}

	gen(asm_source, `import(
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"strconv"
	"strings"
)

func Compile(f io.Reader) Program {
w:= new(bytes.Buffer)
p:=Program{}
	s := bufio.NewScanner(f)
	var fields []string
	tmp := make([]byte, 20)
	putOp := func(o VmOp) {
		n := binary.PutUvarint(tmp, uint64(o))
		w.Write(tmp[:n])
	}
	putInt := func(i int) uint64 {
		v, err := strconv.ParseUint(fields[i], 10, 64)
		check(err)
		n := binary.PutUvarint(tmp, v)
		w.Write(tmp[:n])
                return v
	}
	putFloat := func(i int) {
		v, err := strconv.ParseFloat(fields[i], 64)
		check(err)
		binary.Write(w, order, v)
	}
	for s.Scan() {
		line := s.Text()
		line = strings.TrimSpace(strings.ToLower(line))
		if len(line) == 0 {
			continue
		}
if p.Name == "" && line[0] == '#' {
p.Name = strings.TrimSpace(line[1:])
continue
} else if line[0] == '#' {
continue
}
		fields = strings.Fields(line)
		switch fields[0] {
		case "halt":
			putOp(Halt)
		case "registers":
			putOp(Registers)
			p.Registers = putInt(1)
		case "literal":
			putOp(Literal)
			putInt(1)
			putFloat(2)
		case "outputs":
			putOp(Outputs)
			p.Outputs = putInt(1)
		case "inputs":
			putOp(Inputs)
			p.Inputs = putInt(1)
		case "models":
			putOp(Models)
			p.Models = putInt(1)
		case "getinput":
			putOp(GetInput)
			putInt(1)
			putInt(2)
		case "getmodel":
			putOp(GetModel)
			putInt(1)
			putInt(2)
		case "setoutput":
			putOp(SetOutput)
			putInt(1)
			putInt(2)


{{range $op,$desc := .threes}}
case "{{$op.ToLower}}":
			putOp({{$op}})
			putInt(1)
			putInt(2)
			putInt(3)
{{end}}

{{range $op,$desc := .funcs2}}
case "{{$op.ToLower}}":
			putOp({{$op}})
			putInt(1)
			putInt(2)
			putInt(3)
{{end}}

{{range $op,$desc := .twos}}
case "{{$op.ToLower}}":
			putOp({{$op}})
			putInt(1)
			putInt(2)
{{end}}

		default:
			log.Fatalf("unknown opcode: %s", fields[0])
		}
}
	check(s.Err())
	p.Code = w.Bytes()
return p
}`)

	gen(vm_source, `import (
"math"
"io"
	"bytes"
	"encoding/binary"
"fmt"
)

func Execute(p Program, model, in, out []float64) (err error) {
        if len(out) == 0 {
            return fmt.Errorf("needs at least one output")
        }
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from: %v", r)
		}
	}()
	fmt.Printf("%d byte program %v\n", len(p.Code), p)
	r := bytes.NewReader(p.Code)
	one := func() uint64 {
		a, err := binary.ReadUvarint(r)
		check(err)
		return a
	}
	two := func() (uint64, uint64) {
		a, err := binary.ReadUvarint(r)
		check(err)
		b, err := binary.ReadUvarint(r)
		check(err)
		return a, b
	}
	three := func() (uint64, uint64, uint64) {
		a, err := binary.ReadUvarint(r)
		check(err)
		b, err := binary.ReadUvarint(r)
		check(err)
		c, err := binary.ReadUvarint(r)
		check(err)
		return a, b, c
	}
var registers []float64

Loop:
	for {
		c, err := binary.ReadUvarint(r)
if err == io.EOF {
			break Loop
}
		check(err)
		// fmt.Printf("op = %s\n", VmOp(c))
		// general rules:
		// locations stored first, values later
		// source first, destination after
		switch VmOp(c) {
		case Literal: // store a literal to register
                        loc := one()
			var lit float64
			binary.Read(r, order, &lit)
			registers[loc] = lit
		case Registers: 
                        registers = make([]float64,one())
		case Inputs:
                        n := one();
                        if len(in) < int(n) {
                            return fmt.Errorf("needs at least %d inputs; has %d",n,len(in));
                        }
		case Models:
                        n := one();
                        if len(model) < int(n) {
                            return fmt.Errorf("needs at least %d model dimensions; has %d",n,len(model));
                        }
		case Outputs:
                        n := one();
                        if len(out) < int(n) {
                            return fmt.Errorf("needs at least %d outputs; has %d",n,len(out));
                        }
                case SetOutput: // set output from register
			src, dest := two()
			out[dest] = registers[src]
                case GetInput: // set register from input
			src, dest := two()
			registers[dest] = in[src]
                case GetModel: // set register from model
			src, dest := two()
			registers[dest] = model[src]
		case Halt:
			break Loop

{{range $name,$func := .twos}}case {{$name}}:
			src, dest := two()
			registers[dest] = {{$func}}(registers[src])
{{end}} 

{{range $name,$op := .threes}}case {{$name}}:
			a, b, dest := three()
			registers[dest] = registers[a] {{$op}} registers[b]
{{end}} 

{{range $name,$op := .funcs2}}case {{$name}}:
			a, b, dest := three()
			registers[dest] = {{$op}}(registers[a], registers[b])
{{end}} 

	default:
			return fmt.Errorf("unhandled op %s", VmOp(c))
		}
	}
	return
}

`)
}

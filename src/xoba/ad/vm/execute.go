// autogenerated, do not edit!
package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
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
			registers = make([]float64, one())
		case Inputs:
			n := one()
			if len(in) < int(n) {
				return fmt.Errorf("needs at least %d inputs; has %d", n, len(in))
			}
		case Models:
			n := one()
			if len(model) < int(n) {
				return fmt.Errorf("needs at least %d model dimensions; has %d", n, len(model))
			}
		case Outputs:
			n := one()
			if len(out) < int(n) {
				return fmt.Errorf("needs at least %d outputs; has %d", n, len(out))
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

		case Abs:
			src, dest := two()
			registers[dest] = math.Abs(registers[src])
		case Acos:
			src, dest := two()
			registers[dest] = math.Acos(registers[src])
		case Asin:
			src, dest := two()
			registers[dest] = math.Asin(registers[src])
		case Atan:
			src, dest := two()
			registers[dest] = math.Atan(registers[src])
		case Cos:
			src, dest := two()
			registers[dest] = math.Cos(registers[src])
		case Cosh:
			src, dest := two()
			registers[dest] = math.Cosh(registers[src])
		case D_Abs_D0:
			src, dest := two()
			registers[dest] = d_abs_d0(registers[src])
		case D_Acos_D0:
			src, dest := two()
			registers[dest] = d_acos_d0(registers[src])
		case D_Asin_D0:
			src, dest := two()
			registers[dest] = d_asin_d0(registers[src])
		case D_Atan_D0:
			src, dest := two()
			registers[dest] = d_atan_d0(registers[src])
		case D_Cos_D0:
			src, dest := two()
			registers[dest] = d_cos_d0(registers[src])
		case D_Cosh_D0:
			src, dest := two()
			registers[dest] = d_cosh_d0(registers[src])
		case D_Exp10_D0:
			src, dest := two()
			registers[dest] = d_exp10_d0(registers[src])
		case D_Exp2_D0:
			src, dest := two()
			registers[dest] = d_exp2_d0(registers[src])
		case D_Exp_D0:
			src, dest := two()
			registers[dest] = d_exp_d0(registers[src])
		case D_Log10_D0:
			src, dest := two()
			registers[dest] = d_log10_d0(registers[src])
		case D_Log2_D0:
			src, dest := two()
			registers[dest] = d_log2_d0(registers[src])
		case D_Log_D0:
			src, dest := two()
			registers[dest] = d_log_d0(registers[src])
		case D_Sin_D0:
			src, dest := two()
			registers[dest] = d_sin_d0(registers[src])
		case D_Sinh_D0:
			src, dest := two()
			registers[dest] = d_sinh_d0(registers[src])
		case D_Sqrt_D0:
			src, dest := two()
			registers[dest] = d_sqrt_d0(registers[src])
		case D_Tan_D0:
			src, dest := two()
			registers[dest] = d_tan_d0(registers[src])
		case D_Tanh_D0:
			src, dest := two()
			registers[dest] = d_tanh_d0(registers[src])
		case Exp:
			src, dest := two()
			registers[dest] = math.Exp(registers[src])
		case Exp10:
			src, dest := two()
			registers[dest] = exp10(registers[src])
		case Exp2:
			src, dest := two()
			registers[dest] = math.Exp2(registers[src])
		case Log:
			src, dest := two()
			registers[dest] = math.Log(registers[src])
		case Log10:
			src, dest := two()
			registers[dest] = math.Log10(registers[src])
		case Log2:
			src, dest := two()
			registers[dest] = math.Log2(registers[src])
		case Sin:
			src, dest := two()
			registers[dest] = math.Sin(registers[src])
		case Sinh:
			src, dest := two()
			registers[dest] = math.Sinh(registers[src])
		case Sqrt:
			src, dest := two()
			registers[dest] = math.Sqrt(registers[src])
		case Tan:
			src, dest := two()
			registers[dest] = math.Tan(registers[src])
		case Tanh:
			src, dest := two()
			registers[dest] = math.Tanh(registers[src])

		case Add:
			a, b, dest := three()
			registers[dest] = registers[a] + registers[b]
		case Divide:
			a, b, dest := three()
			registers[dest] = registers[a] / registers[b]
		case Multiply:
			a, b, dest := three()
			registers[dest] = registers[a] * registers[b]
		case Subtract:
			a, b, dest := three()
			registers[dest] = registers[a] - registers[b]

		case D_Add_D0:
			a, b, dest := three()
			registers[dest] = d_add_d0(registers[a], registers[b])
		case D_Add_D1:
			a, b, dest := three()
			registers[dest] = d_add_d1(registers[a], registers[b])
		case D_Divide_D0:
			a, b, dest := three()
			registers[dest] = d_divide_d0(registers[a], registers[b])
		case D_Divide_D1:
			a, b, dest := three()
			registers[dest] = d_divide_d1(registers[a], registers[b])
		case D_Multiply_D0:
			a, b, dest := three()
			registers[dest] = d_multiply_d0(registers[a], registers[b])
		case D_Multiply_D1:
			a, b, dest := three()
			registers[dest] = d_multiply_d1(registers[a], registers[b])
		case D_Pow_D0:
			a, b, dest := three()
			registers[dest] = d_pow_d0(registers[a], registers[b])
		case D_Pow_D1:
			a, b, dest := three()
			registers[dest] = d_pow_d1(registers[a], registers[b])
		case D_Subtract_D0:
			a, b, dest := three()
			registers[dest] = d_subtract_d0(registers[a], registers[b])
		case D_Subtract_D1:
			a, b, dest := three()
			registers[dest] = d_subtract_d1(registers[a], registers[b])
		case Pow:
			a, b, dest := three()
			registers[dest] = math.Pow(registers[a], registers[b])

		default:
			return fmt.Errorf("unhandled op %s", VmOp(c))
		}
	}

	fmt.Printf("registers = %.3f\n", registers)
	return
}

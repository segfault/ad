// autogenerated, do not edit!
package vm

type VmOp uint64

const (
	_ VmOp = iota

	Abs       // absolute value
	Acos      //
	Add       //
	Asin      //
	Atan      //
	Cos       //
	Cosh      //
	Divide    //
	Exp       //
	Exp10     // 10^x
	Exp2      // 2^x
	Halt      //
	Inputs    // validate input dimension is large enough
	Literal   //
	Log       //
	Log10     //
	Log2      //
	Multiply  //
	Outputs   // validate output dimension is large enough
	Pow       //
	Registers // 1 argument, sets the number of registers
	SetOutput //
	Sin       //
	Sinh      //
	Sqrt      //
	Subtract  //
	Tan       //
	Tanh      //

)

func (o VmOp) String() string {
	switch o {
	case Abs:
		return "Abs"
	case Acos:
		return "Acos"
	case Add:
		return "Add"
	case Asin:
		return "Asin"
	case Atan:
		return "Atan"
	case Cos:
		return "Cos"
	case Cosh:
		return "Cosh"
	case Divide:
		return "Divide"
	case Exp:
		return "Exp"
	case Exp10:
		return "Exp10"
	case Exp2:
		return "Exp2"
	case Halt:
		return "Halt"
	case Inputs:
		return "Inputs"
	case Literal:
		return "Literal"
	case Log:
		return "Log"
	case Log10:
		return "Log10"
	case Log2:
		return "Log2"
	case Multiply:
		return "Multiply"
	case Outputs:
		return "Outputs"
	case Pow:
		return "Pow"
	case Registers:
		return "Registers"
	case SetOutput:
		return "SetOutput"
	case Sin:
		return "Sin"
	case Sinh:
		return "Sinh"
	case Sqrt:
		return "Sqrt"
	case Subtract:
		return "Subtract"
	case Tan:
		return "Tan"
	case Tanh:
		return "Tanh"
	}
	panic("illegal state")
}

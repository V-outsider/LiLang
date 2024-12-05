package compiler

import (
	"fmt"
	"strconv"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	generated "github.com/V-outsider/LiLang/generated"
)

func Test() {
	m := ir.NewModule()
	hello := constant.NewCharArrayFromString("Hello, world!\n\x00")
	str := m.NewGlobalDef("str", hello)

	puts := m.NewFunc("puts", types.I32, ir.NewParam("", types.NewPointer(types.I8)))
	main := m.NewFunc("main", types.I32)
	entry := main.NewBlock("")

	zero := constant.NewInt(types.I64, 0)
	gep := constant.NewGetElementPtr(hello.Typ, str, zero, zero)
	entry.NewCall(puts, gep)
	entry.NewRet(constant.NewInt(types.I32, 0))
	fmt.Println(m)
}

type CodegenState struct {
	Module           *ir.Module
	Current          *ir.Block
	Symbols          map[string]interface{}
	DefinedFunctions map[string]bool // Track defined functions
}

type IRListener struct {
	*generated.BaseLiLangListener
	State *CodegenState
}

func NewCodegenState() *CodegenState {
	return &CodegenState{
		Module:           ir.NewModule(),
		Symbols:          make(map[string]interface{}),
		DefinedFunctions: make(map[string]bool),
	}
}

func NewIRListener() *IRListener {
	state := NewCodegenState()

	return &IRListener{
		BaseLiLangListener: &generated.BaseLiLangListener{},
		State:              state,
	}
}

func (l *IRListener) EnterProgram(ctx *generated.ProgramContext) {}

func (l *IRListener) ExitProgram(ctx *generated.ProgramContext) {
	if l.State.Current != nil && l.State.Current.Term == nil {
		if xVal, ok := l.State.Symbols["x"].(*ir.Global); ok {
			loadX := l.State.Current.NewLoad(types.I32, xVal)
			l.State.Current.NewRet(loadX)
		} else {
			l.State.Current.NewRet(constant.NewInt(types.I32, 0))
		}
	}
}

func (l *IRListener) EnterStatement(ctx *generated.StatementContext) {
	if l.State.Current == nil {
		if !l.State.DefinedFunctions["main"] {
			mainFunc := l.State.Module.NewFunc("main", types.I32)
			l.State.Current = mainFunc.NewBlock("entry")
			l.State.DefinedFunctions["main"] = true
			l.State.Symbols["#currentFunc"] = mainFunc
		}
	}
}

func (l *IRListener) EnterVariableDeclaration(ctx *generated.VariableDeclarationContext) {
	name := ctx.IDENTIFIER().GetText()
	varType := ctx.Type_().GetText()

	llvmType := types.I32
	if varType != "int" {
		panic(fmt.Sprintf("Unsupported type: %s", varType))
	}

	globalVar := l.State.Module.NewGlobal(name, llvmType)
	globalVar.Init = constant.NewInt(types.I32, 0) // Default initialization

	l.State.Symbols[name] = globalVar
}

func (l *IRListener) EnterAssignment(ctx *generated.AssignmentContext) {
	name := ctx.IDENTIFIER().GetText()
	expr := ctx.Expression()

	value := l.evaluateExpression(expr.(*generated.ExpressionContext))

	varValue, ok := l.State.Symbols[name]
	if !ok {
		panic(fmt.Sprintf("Undefined variable: %s", name))
	}

	globalVar, ok := varValue.(*ir.Global)
	if !ok {
		panic(fmt.Sprintf("Variable is not of type *ir.Global: %s", name))
	}
	l.State.Current.NewStore(value, globalVar)
}

func (l *IRListener) evaluateExpression(ctx *generated.ExpressionContext) value.Value {
	if num := ctx.NUMBER(); num != nil {
		value, _ := strconv.ParseInt(num.GetText(), 10, 64)
		return constant.NewInt(types.I32, value)
	}

	if id := ctx.IDENTIFIER(); id != nil {
		name := id.GetText()
		if val, exists := l.State.Symbols[name]; exists {
			if global, ok := val.(*ir.Global); ok {
				// Load the value from the global variable
				return l.State.Current.NewLoad(types.I32, global)
			}
			return val.(value.Value)
		}
		panic(fmt.Sprintf("Undefined variable: %s", name))
	}

	if left := ctx.Expression(0); left != nil {
		right := ctx.Expression(1)
		leftValue := l.evaluateExpression(left.(*generated.ExpressionContext))
		rightValue := l.evaluateExpression(right.(*generated.ExpressionContext))

		if l.State.Current == nil {
			panic("No current block set for IR generation")
		}

		switch {
		case ctx.PLUS() != nil:
			return l.State.Current.NewAdd(leftValue, rightValue)
		case ctx.MINUS() != nil:
			return l.State.Current.NewSub(leftValue, rightValue)
		case ctx.MULT() != nil:
			return l.State.Current.NewMul(leftValue, rightValue)
		case ctx.DIV() != nil:
			return l.State.Current.NewSDiv(leftValue, rightValue)
		}
	}

	if fnCall := ctx.FunctionCall(); fnCall != nil {
		return l.handleFunctionCall(fnCall.(*generated.FunctionCallContext))
	}

	if ctx.LPAREN() != nil && ctx.RPAREN() != nil {
		return l.evaluateExpression(ctx.Expression(0).(*generated.ExpressionContext))
	}

	panic("Unsupported expression type")
}

func (l *IRListener) handleFunctionCall(ctx *generated.FunctionCallContext) value.Value {
	funcName := ctx.IDENTIFIER().GetText()

	var fn *ir.Func
	for _, f := range l.State.Module.Funcs {
		if f.Name() == funcName {
			fn = f
			break
		}
	}

	if fn == nil {
		panic(fmt.Sprintf("Undefined function: %s", funcName))
	}

	var args []value.Value
	if argList := ctx.ArgList(); argList != nil {
		for _, expr := range argList.AllExpression() {
			arg := l.evaluateExpression(expr.(*generated.ExpressionContext))
			args = append(args, arg)
		}
	}

	return l.State.Current.NewCall(fn, args...)
}

func (l *IRListener) EnterFunctionDeclaration(ctx *generated.FunctionDeclarationContext) {
	name := ctx.IDENTIFIER().GetText()

	oldBlock := l.State.Current
	oldFunc := l.State.Symbols["#currentFunc"]

	if name == "main" && l.State.Current != nil {
		l.State.Symbols["#prevBlock"] = oldBlock
		l.State.Symbols["#prevFunc"] = oldFunc
		return
	}

	if l.State.DefinedFunctions[name] {
		panic(fmt.Sprintf("Function %s is already defined", name))
	}

	returnType := types.I32

	var params []*ir.Param
	if paramList := ctx.ParamList(); paramList != nil {
		for _, param := range paramList.AllParam() {
			paramCtx := param.(*generated.ParamContext)
			paramName := paramCtx.IDENTIFIER().GetText()
			params = append(params, ir.NewParam(paramName, types.I32))
		}
	}

	fn := l.State.Module.NewFunc(name, returnType, params...)
	block := fn.NewBlock("entry")

	l.State.Current = block
	l.State.Symbols["#currentFunc"] = fn
	l.State.Symbols["#prevBlock"] = oldBlock
	l.State.Symbols["#prevFunc"] = oldFunc

	for _, param := range params {
		l.State.Symbols[param.Name()] = param
	}

	l.State.DefinedFunctions[name] = true
}

func (l *IRListener) EnterReturnStatement(ctx *generated.ReturnStatementContext) {
	if ctx.Expression() != nil {
		returnValue := l.evaluateExpression(ctx.Expression().(*generated.ExpressionContext))
		l.State.Current.NewRet(returnValue)
	} else {
		l.State.Current.NewRet(constant.NewInt(types.I32, 0))
	}
}

func (l *IRListener) EnterPrintCall(ctx *generated.PrintCallContext) {
	value := l.evaluateExpression(ctx.Expression().(*generated.ExpressionContext))

	var printf *ir.Func
	for _, f := range l.State.Module.Funcs {
		if f.Name() == "printf" {
			printf = f
			break
		}
	}
	if printf == nil {
		printf = l.State.Module.NewFunc("printf", types.I32,
			ir.NewParam("", types.NewPointer(types.I8)))
		printf.Sig.Variadic = true
	}

	formatStr := l.State.Module.NewGlobalDef("printf.str", constant.NewCharArrayFromString("%d\n"))
	formatPtr := l.State.Current.NewGetElementPtr(
		types.NewArray(uint64(len("%d\n")), types.I8),
		formatStr,
		constant.NewInt(types.I64, 0),
		constant.NewInt(types.I64, 0))

	l.State.Current.NewCall(printf, formatPtr, value)
}

func (l *IRListener) ExitFunctionDeclaration(ctx *generated.FunctionDeclarationContext) {
	if l.State.Current != nil && l.State.Current.Term == nil {
		l.State.Current.NewRet(constant.NewInt(types.I32, 0))
	}

	if prev, ok := l.State.Symbols["#prevBlock"].(*ir.Block); ok {
		l.State.Current = prev
	}
	if prevFunc, ok := l.State.Symbols["#prevFunc"]; ok {
		l.State.Symbols["#currentFunc"] = prevFunc
	}

	delete(l.State.Symbols, "#prevBlock")
	delete(l.State.Symbols, "#prevFunc")
}

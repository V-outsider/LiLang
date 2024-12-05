package compiler

import (
	"fmt"
	"strconv"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"

	generated "github.com/V-outsider/LiLang/generated"
)

func Test() {
	m := ir.NewModule()
	hello := constant.NewCharArrayFromString("Hello, world!\n\x00")
	str := m.NewGlobalDef("str", hello)
	// Add external function declaration of puts.
	puts := m.NewFunc("puts", types.I32, ir.NewParam("", types.NewPointer(types.I8)))
	main := m.NewFunc("main", types.I32)
	entry := main.NewBlock("")
	// Cast *[15]i8 to *i8.
	zero := constant.NewInt(types.I64, 0)
	gep := constant.NewGetElementPtr(hello.Typ, str, zero, zero)
	entry.NewCall(puts, gep)
	entry.NewRet(constant.NewInt(types.I32, 0))
	fmt.Println(m)
}

type CodegenState struct {
	Module  *ir.Module             // LLVM module being generated
	Current *ir.Block              // Current basic block for instructions
	Symbols map[string]interface{}
}

type IRListener struct {
    *generated.BaseLiLangListener
    State *CodegenState
}

func NewCodegenState() *CodegenState {
	return &CodegenState{
		Module:  ir.NewModule(),
		Symbols: make(map[string]interface{}),
	}
}

func NewIRListener() *IRListener {
    return &IRListener{State: NewCodegenState()}
}

// Handle variable declarations
func (l *IRListener) EnterVariableDeclaration(ctx *generated.VariableDeclarationContext) {
    name := ctx.IDENTIFIER().GetText()
    varType := ctx.Type_().GetText()

    // Determine LLVM type
    llvmType := types.I32 // Example for `int`
    if varType != "int" {
        panic(fmt.Sprintf("Unsupported type: %s", varType))
    }

    // Allocate variable in the module
    globalVar := l.State.Module.NewGlobal(name, llvmType)
    globalVar.Init = constant.NewInt(types.I32, 0) // Default initialization

    // Save the variable in the symbol table
    l.State.Symbols[name] = globalVar
}

// Handle assignments
func (l *IRListener) EnterAssignment(ctx *generated.AssignmentContext) {
    name := ctx.IDENTIFIER().GetText()
    expr := ctx.Expression() // Expression context

    // Get the value from the symbol table
    varValue, ok := l.State.Symbols[name]
    if !ok {
        panic(fmt.Sprintf("Undefined variable: %s", name))
    }

    // For simplicity, handle only constant assignments
    if num := expr.NUMBER(); num != nil {
        value := constant.NewInt(types.I32, strconv.ParseInt(num.GetText(), 10, 64))
        l.State.Current.NewStore(value, varValue) // Store the value
    }
}

// Handle function declarations
func (l *IRListener) EnterFunctionDeclaration(ctx *generated.FunctionDeclarationContext) {
    name := ctx.IDENTIFIER().GetText()
    fn := l.State.Module.NewFunc(name, types.I32) // Assuming return type is `int`

    l.State.Current = fn.NewBlock("entry") // Create a new entry block
}

// Handle return statements
func (l *IRListener) EnterReturnStatement(ctx *generated.ReturnStatementContext) {
    if ctx.Expression() != nil {
        returnValue := constant.NewInt(types.I32, 0) // Replace with expression evaluation
        l.State.Current.NewRet(returnValue)
    }
}

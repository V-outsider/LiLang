package analyzers

import (
	"fmt"

	generated "github.com/V-outsider/LiLang/generated"
)

type StaticAnalyzer struct {
	*generated.BaseLiLangListener
	variables map[string]string
}

func NewStaticAnalyzer() *StaticAnalyzer {
	return &StaticAnalyzer{
		variables: make(map[string]string),
	}
}

func (a *StaticAnalyzer) EnterVariableDeclaration(ctx *generated.VariableDeclarationContext) {
	varName := ctx.IDENTIFIER().GetText()
	varType := ctx.Type_().GetText()

	a.variables[varName] = varType
	fmt.Printf("Объявлена переменная: %s с типом: %s\n", varName, varType)
}

func (a *StaticAnalyzer) EnterAssignment(ctx *generated.AssignmentContext) {
	varName := ctx.IDENTIFIER().GetText()

	// Проверяем, объявлена ли переменная
	if _, exists := a.variables[varName]; !exists {
		fmt.Printf("Ошибка: переменная %s не была объявлена!\n", varName)
	} else {
		fmt.Printf("Присваивание переменной: %s\n", varName)
	}
}

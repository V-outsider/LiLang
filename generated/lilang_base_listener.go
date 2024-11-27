// Code generated from LiLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // LiLang

import "github.com/antlr4-go/antlr/v4"

// BaseLiLangListener is a complete listener for a parse tree produced by LiLangParser.
type BaseLiLangListener struct{}

var _ LiLangListener = &BaseLiLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLiLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLiLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLiLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLiLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseLiLangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseLiLangListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseLiLangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseLiLangListener) ExitStatement(ctx *StatementContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseLiLangListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseLiLangListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseLiLangListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseLiLangListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterPrintCall is called when production printCall is entered.
func (s *BaseLiLangListener) EnterPrintCall(ctx *PrintCallContext) {}

// ExitPrintCall is called when production printCall is exited.
func (s *BaseLiLangListener) ExitPrintCall(ctx *PrintCallContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseLiLangListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseLiLangListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BaseLiLangListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseLiLangListener) ExitParam(ctx *ParamContext) {}

// EnterType is called when production type is entered.
func (s *BaseLiLangListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseLiLangListener) ExitType(ctx *TypeContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseLiLangListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseLiLangListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseLiLangListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseLiLangListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLiLangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLiLangListener) ExitBlock(ctx *BlockContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseLiLangListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseLiLangListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseLiLangListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseLiLangListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *BaseLiLangListener) EnterCaseStatement(ctx *CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *BaseLiLangListener) ExitCaseStatement(ctx *CaseStatementContext) {}

// EnterListOperation is called when production listOperation is entered.
func (s *BaseLiLangListener) EnterListOperation(ctx *ListOperationContext) {}

// ExitListOperation is called when production listOperation is exited.
func (s *BaseLiLangListener) ExitListOperation(ctx *ListOperationContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseLiLangListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseLiLangListener) ExitMethod(ctx *MethodContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseLiLangListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseLiLangListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLiLangListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLiLangListener) ExitExpression(ctx *ExpressionContext) {}

// Generated from /Users/voutsider/Documents/Labs/LPIS/LiLang/LiLang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link LiLangParser}.
 */
public interface LiLangListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link LiLangParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(LiLangParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(LiLangParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(LiLangParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(LiLangParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterVariableDeclaration(LiLangParser.VariableDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitVariableDeclaration(LiLangParser.VariableDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterFunctionDeclaration(LiLangParser.FunctionDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitFunctionDeclaration(LiLangParser.FunctionDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#paramList}.
	 * @param ctx the parse tree
	 */
	void enterParamList(LiLangParser.ParamListContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#paramList}.
	 * @param ctx the parse tree
	 */
	void exitParamList(LiLangParser.ParamListContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#param}.
	 * @param ctx the parse tree
	 */
	void enterParam(LiLangParser.ParamContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#param}.
	 * @param ctx the parse tree
	 */
	void exitParam(LiLangParser.ParamContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#type}.
	 * @param ctx the parse tree
	 */
	void enterType(LiLangParser.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#type}.
	 * @param ctx the parse tree
	 */
	void exitType(LiLangParser.TypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#assignment}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(LiLangParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#assignment}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(LiLangParser.AssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterExpressionStatement(LiLangParser.ExpressionStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitExpressionStatement(LiLangParser.ExpressionStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(LiLangParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(LiLangParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(LiLangParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(LiLangParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStatement(LiLangParser.SwitchStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStatement(LiLangParser.SwitchStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#caseStatement}.
	 * @param ctx the parse tree
	 */
	void enterCaseStatement(LiLangParser.CaseStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#caseStatement}.
	 * @param ctx the parse tree
	 */
	void exitCaseStatement(LiLangParser.CaseStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void enterReturnStatement(LiLangParser.ReturnStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void exitReturnStatement(LiLangParser.ReturnStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LiLangParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(LiLangParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link LiLangParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(LiLangParser.ExpressionContext ctx);
}
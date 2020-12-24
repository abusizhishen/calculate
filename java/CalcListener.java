// Generated from Calc.g4 by ANTLR 4.7.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link CalcParser}.
 */
public interface CalcListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link CalcParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(CalcParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link CalcParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(CalcParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link CalcParser#row}.
	 * @param ctx the parse tree
	 */
	void enterRow(CalcParser.RowContext ctx);
	/**
	 * Exit a parse tree produced by {@link CalcParser#row}.
	 * @param ctx the parse tree
	 */
	void exitRow(CalcParser.RowContext ctx);
	/**
	 * Enter a parse tree produced by {@link CalcParser#last}.
	 * @param ctx the parse tree
	 */
	void enterLast(CalcParser.LastContext ctx);
	/**
	 * Exit a parse tree produced by {@link CalcParser#last}.
	 * @param ctx the parse tree
	 */
	void exitLast(CalcParser.LastContext ctx);
	/**
	 * Enter a parse tree produced by {@link CalcParser#file}.
	 * @param ctx the parse tree
	 */
	void enterFile(CalcParser.FileContext ctx);
	/**
	 * Exit a parse tree produced by {@link CalcParser#file}.
	 * @param ctx the parse tree
	 */
	void exitFile(CalcParser.FileContext ctx);
}
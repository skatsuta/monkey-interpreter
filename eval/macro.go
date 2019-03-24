package eval

import (
	"github.com/skatsuta/monkey-interpreter/ast"
	"github.com/skatsuta/monkey-interpreter/object"
)

// DefineMacros finds macro definitions in the program, saves them to a given environment and
// removes them from the AST.
func DefineMacros(program *ast.Program, env object.Environment) {
	defs := make([]int, 0)
	stmts := program.Statements

	for pos, stmt := range stmts {
		if isMacroDefinition(stmt) {
			addMacro(stmt, env)
			defs = append(defs, pos)
		}
	}

	for i := len(defs) - 1; i >= 0; i-- {
		pos := defs[i]
		program.Statements = append(stmts[:pos], stmts[pos+1:]...)
	}
}

func isMacroDefinition(node ast.Statement) bool {
	letStmt, ok := node.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = letStmt.Value.(*ast.MacroLiteral)
	return ok
}

func addMacro(stmt ast.Statement, env object.Environment) {
	letStmt := stmt.(*ast.LetStatement)
	macroLit := letStmt.Value.(*ast.MacroLiteral)
	macro := &object.Macro{
		Parameters: macroLit.Parameters,
		Env:        env,
		Body:       macroLit.Body,
	}
	env.Set(letStmt.Name.Value, macro)
}

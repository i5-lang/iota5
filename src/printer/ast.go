// SPDX-License-Identifier: GPL-3.0-or-later
package printer

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/io/console"
)

func printAst(_ast ast.Node) {
	_ast.Print()
	console.Println()
}

// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Builtin struct {
	Value string
}

func (b Builtin) StringValue() string {
	return "$" + b.Value
}

func (b Builtin) expression() {}

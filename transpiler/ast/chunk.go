package ast

import "strings"

type ASTChunk struct {
	items []IASTItem
}

func NewASTChunk() *ASTChunk {
	return &ASTChunk{}
}

func (c *ASTChunk) Add(item IASTItem) {
	c.items = append(c.items, item)
}

func (c *ASTChunk) String() string {
	lines := make([]string, len(c.items))

	for i, item := range c.items {
		lines[i] += item.String()
	}

	return strings.Join(lines, "\n")
}

package transpiler

import (
	"testing"

	"github.com/tereus-project/tereus-transpiler-std/core"
)

func assertTranspilation(t *testing.T, sourceCode string, targetCode string) {
	core.AssertTranspilation(t, &core.AssertTranspilationConfig{
		SourceFilename:    "test.c",
		SourceCode:        sourceCode,
		TargetCode:        targetCode,
		TranspileFunction: Transpile,
	})
}

func TestVariableDeclaration(t *testing.T) {
	source := `
a = 5
b = 3.14
`

	target := `
a = 5
b = 3.14
`

	assertTranspilation(t, source, target)
}

func TestAddition(t *testing.T) {
	source := `
a = 5 + 3
b = 3.14 + 2.71
`

	target := `
a = 5 + 3
b = 3.14 + 2.71
`

	assertTranspilation(t, source, target)
}

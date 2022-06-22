package main

import (
	"github.com/tereus-project/tereus-transpiler-std/core"
	"github.com/tereus-project/tereus-transpiler-template/transpiler"
)

func main() {
	core.InitTranspiler(&core.TranspilerContextConfig{
		SourceLanguage:              "lua",
		SourceLanguageFileExtension: ".lua",
		TargetLanguage:              "ruby",
		TargetLanguageFileExtension: ".rb",
		TranspileFunction:           transpiler.Transpile,
	})
}

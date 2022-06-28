package main

import (
	"github.com/tereus-project/tereus-transpiler-lua-ruby/transpiler"
	"github.com/tereus-project/tereus-transpiler-std/core"
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

/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package module_test

import (
	fmt "fmt"
	gen "github.com/craterdog/go-syntax-generation/v5"
	not "github.com/craterdog/go-syntax-notation/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

var syntaxFile = "../../go-test-framework/v5/Syntax.cdsn"

var modelFiles = []string{
	"../../go-test-framework/v5/ast/Package.go",
	"../../go-test-framework/v5/generator/Package.go",
	"../../go-test-framework/v5/grammar/Package.go",
	"../../go-test-framework/v5/example/Package.go",
}

func TestGeneration(t *tes.T) {
	fmt.Println("Reading and validating the following language syntax:")
	fmt.Printf("   %v\n", syntaxFile)
	var bytes, err = osx.ReadFile(syntaxFile)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)
	fmt.Println()

	fmt.Println("Generating the following class models:")
	var wiki = "github.com/craterdog/go-test-framework/wiki"
	for _, modelFile := range modelFiles {
		var classModel string
		fmt.Printf("   %v\n", modelFile)
		switch {
		case sts.Contains(modelFile, "ast/"):
			classModel = gen.GenerateAstModel(wiki, syntax)
		case sts.Contains(modelFile, "grammar/"):
		case sts.Contains(modelFile, "generator/"):
		case sts.Contains(modelFile, "example/"):
		}
		bytes = []byte(classModel)
		err = osx.WriteFile(modelFile, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Done.")
}

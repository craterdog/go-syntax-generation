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
	tes "testing"
)

func TestGeneration(t *tes.T) {
	fmt.Println("Reading and validating the following language syntax:")
	var syntaxFile = "../../go-test-framework/v5/Syntax.cdsn"
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

	var modelFile = "../../go-test-framework/v5/ast/Package.go"
	fmt.Printf("   %v\n", modelFile)
	var classModel = gen.GenerateAstModel(wiki, syntax)
	bytes = []byte(classModel)
	err = osx.WriteFile(modelFile, bytes, 0644)
	if err != nil {
		panic(err)
	}

	/*
		modelFile = "../../go-test-framework/v5/grammar/Package.go"
		fmt.Printf("   %v\n", modelFile)
		classModel = gen.GenerateGrammarModel(wiki, syntax)
		bytes = []byte(classModel)
		err = osx.WriteFile(modelFile, bytes, 0644)
		if err != nil {
			panic(err)
		}
	*/

	fmt.Println("Done.")
}

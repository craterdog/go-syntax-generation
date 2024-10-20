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

package generator_test

import (
	ge2 "github.com/craterdog/go-class-generation/v5"
	mod "github.com/craterdog/go-class-model/v5"
	gen "github.com/craterdog/go-syntax-generation/v5/generator"
	not "github.com/craterdog/go-syntax-notation/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var directory = "../testdata/"

func TestRoundTrips(t *tes.T) {
	// Read in the syntax notation and validate it.
	var filename = directory + "Syntax.cdsn"
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var syntax = not.ParseSource(source)
	not.ValidateSyntax(syntax)
	var actual = not.FormatSyntax(syntax)
	ass.Equal(t, source, actual)

	// Generate the AST Package.go file.
	filename = directory + "ast/Package.go"
	var wiki = "github.com/craterdog/go-test-framework/wiki"
	source = gen.AstGenerator().Make().GenerateAstModel(wiki, syntax)
	bytes = []byte(source)
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the AST classes.
	var model = mod.ParseSource(source)
	var classes = ge2.GenerateModelClasses(model).GetIterator()
	for classes.HasNext() {
		var class = classes.GetNext()
		var className = class.GetKey()
		filename = directory + "ast/" + className + ".go"
		source = class.GetValue()
		bytes = []byte(source)
		err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}

	// Generate the grammar Package.go file.

	// Generate the formatter class file.

	// Generate the parser class file.

	// Generate the processor class file.

	// Generate the scanner class file.

	// Generate the token class file.

	// Generate the validator class file.

	// Generate the visitor class file.
}

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

/*
Package "module" defines type aliases for the commonly used types defined in the
packages contained in this module.  It also provides a universal constructor for
the commonly used classes that are exported by this module.  Each constructor
delegates the actual construction process to its corresponding concrete class
defined in a package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-syntax-generator/wiki
*/
package module

import (
	gen "github.com/craterdog/go-syntax-generation/v5/generator"
	not "github.com/craterdog/go-syntax-notation/v5"
)

// TYPE ALIASES

type (
	AstGeneratorLike = gen.AstGeneratorLike
	//FormatterGeneratorLike = gen.FormatterGeneratorLike
	//GrammarGeneratorLike = gen.GrammarGeneratorLike
	//ParserGeneratorLike = gen.ParserGeneratorLike
	//ProcessorGeneratorLike = gen.ProcessorGeneratorLike
	//SyntaxGeneratorLike = gen.SyntaxGeneratorLike
	//ValidatorGeneratorLike = gen.ValidatorGeneratorLike
	//VisitorGeneratorLike = gen.VisitorGeneratorLike
)

// UNIVERSAL CONSTRUCTORS

func AstGenerator(args ...any) AstGeneratorLike {
	if len(args) > 0 {
		panic("The \"AST\" generator constructor does not take any arguments.")
	}
	var generator = gen.AstGenerator().Make()
	return generator
}

// GLOBAL FUNCTIONS

func GenerateAstModel(
	wiki string,
	syntax not.SyntaxLike,
) string {
	var generator = AstGenerator()
	var catalog = generator.GenerateAstModel(wiki, syntax)
	return catalog
}

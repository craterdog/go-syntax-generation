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
Package "generator" provides a template-based code generator that can generate
the class model packages for both the abstract syntax tree (AST) and the
language grammar tools for processing any language defined using Crater Dog
Syntax Notation™ (CDSN).

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-syntax-generation/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package generator

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-syntax-notation/v5/ast"
	gra "github.com/craterdog/go-syntax-notation/v5/grammar"
)

// Class Definitions

/*
AnalyzerClassLike defines the set of class constants, constructors and
functions that must be supported by all analyzer-class-like classes.
*/
type AnalyzerClassLike interface {
	// Constructor Methods
	Make() AnalyzerLike
}

/*
AstGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all ast-generator-class-like classes.
*/
type AstGeneratorClassLike interface {
	// Constructor Methods
	Make() AstGeneratorLike
}

/*
FormatterGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all formatter-generator-class-like classes.
*/
type FormatterGeneratorClassLike interface {
	// Constructor Methods
	Make() FormatterGeneratorLike
}

/*
GrammarGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all grammar-generator-class-like classes.
*/
type GrammarGeneratorClassLike interface {
	// Constructor Methods
	Make() GrammarGeneratorLike
}

/*
ParserGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all parser-generator-class-like classes.
*/
type ParserGeneratorClassLike interface {
	// Constructor Methods
	Make() ParserGeneratorLike
}

/*
ProcessorGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all processor-generator-class-like classes.
*/
type ProcessorGeneratorClassLike interface {
	// Constructor Methods
	Make() ProcessorGeneratorLike
}

/*
ScannerGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all scanner-generator-class-like classes.
*/
type ScannerGeneratorClassLike interface {
	// Constructor Methods
	Make() ScannerGeneratorLike
}

/*
SyntaxGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all syntax-generator-class-like classes.
*/
type SyntaxGeneratorClassLike interface {
	// Constructor Methods
	Make() SyntaxGeneratorLike
}

/*
TokenGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all token-generator-class-like classes.
*/
type TokenGeneratorClassLike interface {
	// Constructor Methods
	Make() TokenGeneratorLike
}

/*
ValidatorGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all validator-generator-class-like classes.
*/
type ValidatorGeneratorClassLike interface {
	// Constructor Methods
	Make() ValidatorGeneratorLike
}

/*
VisitorGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all visitor-generator-class-like classes.
*/
type VisitorGeneratorClassLike interface {
	// Constructor Methods
	Make() VisitorGeneratorLike
}

// Instance Definitions

/*
AnalyzerLike defines the set of aspects and methods that must be supported by
all analyzer-like instances.
*/
type AnalyzerLike interface {
	// Primary Methods
	GetClass() AnalyzerClassLike
	AnalyzeSyntax(
		syntax ast.SyntaxLike,
	)
	GetExpressions() abs.Sequential[abs.AssociationLike[string, string]]
	GetIdentifiers(
		ruleName string,
	) abs.Sequential[ast.IdentifierLike]
	GetNotice() string
	GetReferences(
		ruleName string,
	) abs.Sequential[ast.ReferenceLike]
	GetRuleNames() abs.Sequential[string]
	GetSyntaxMap() string
	GetSyntaxName() string
	GetTerms(
		ruleName string,
	) abs.Sequential[ast.TermLike]
	GetTokenNames() abs.Sequential[string]
	GetVariableName(
		reference ast.ReferenceLike,
	) string
	GetVariableNames(
		references abs.Sequential[ast.ReferenceLike],
	) abs.Sequential[string]
	GetVariableType(
		reference ast.ReferenceLike,
	) string
	HasPlurals() bool
	IsDelimited(
		ruleName string,
	) bool
	IsPlural(
		name string,
	) bool

	// Aspect Interfaces
	gra.Methodical
}

/*
AstGeneratorLike defines the set of aspects and methods that must be supported by
all ast-generator-like instances.
*/
type AstGeneratorLike interface {
	// Primary Methods
	GetClass() AstGeneratorClassLike
	GenerateAstModel(
		wiki string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
FormatterGeneratorLike defines the set of aspects and methods that must be supported by
all formatter-generator-like instances.
*/
type FormatterGeneratorLike interface {
	// Primary Methods
	GetClass() FormatterGeneratorClassLike
	GenerateFormatterClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
GrammarGeneratorLike defines the set of aspects and methods that must be supported by
all grammar-generator-like instances.
*/
type GrammarGeneratorLike interface {
	// Primary Methods
	GetClass() GrammarGeneratorClassLike
	GenerateGrammarModel(
		module string,
		wiki string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
ParserGeneratorLike defines the set of aspects and methods that must be supported by
all parser-generator-like instances.
*/
type ParserGeneratorLike interface {
	// Primary Methods
	GetClass() ParserGeneratorClassLike
	GenerateParserClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
ProcessorGeneratorLike defines the set of aspects and methods that must be supported by
all processor-generator-like instances.
*/
type ProcessorGeneratorLike interface {
	// Primary Methods
	GetClass() ProcessorGeneratorClassLike
	GenerateProcessorClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
ScannerGeneratorLike defines the set of aspects and methods that must be supported by
all scanner-generator-like instances.
*/
type ScannerGeneratorLike interface {
	// Primary Methods
	GetClass() ScannerGeneratorClassLike
	GenerateScannerClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
SyntaxGeneratorLike defines the set of aspects and methods that must be supported by
all syntax-generator-like instances.
*/
type SyntaxGeneratorLike interface {
	// Primary Methods
	GetClass() SyntaxGeneratorClassLike
	GenerateSyntaxNotation(
		syntax string,
		copyright string,
	) (
		implementation string,
	)
}

/*
TokenGeneratorLike defines the set of aspects and methods that must be supported by
all token-generator-like instances.
*/
type TokenGeneratorLike interface {
	// Primary Methods
	GetClass() TokenGeneratorClassLike
	GenerateTokenClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
ValidatorGeneratorLike defines the set of aspects and methods that must be supported by
all validator-generator-like instances.
*/
type ValidatorGeneratorLike interface {
	// Primary Methods
	GetClass() ValidatorGeneratorClassLike
	GenerateValidatorClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

/*
VisitorGeneratorLike defines the set of aspects and methods that must be supported by
all visitor-generator-like instances.
*/
type VisitorGeneratorLike interface {
	// Primary Methods
	GetClass() VisitorGeneratorClassLike
	GenerateVisitorClass(
		module string,
		syntax ast.SyntaxLike,
	) (
		implementation string,
	)
}

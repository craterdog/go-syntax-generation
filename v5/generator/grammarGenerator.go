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

package generator

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
	ast "github.com/craterdog/go-syntax-notation/v5/ast"
)

// CLASS INTERFACE

// Access Function

func GrammarGenerator() GrammarGeneratorClassLike {
	return grammarGeneratorReference()
}

// Constructor Methods

func (c *grammarGeneratorClass_) Make() GrammarGeneratorLike {
	var instance = &grammarGenerator_{
		// Initialize the instance attributes.
		analyzer_: Analyzer().Make(),
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *grammarGenerator_) GetClass() GrammarGeneratorClassLike {
	return grammarGeneratorReference()
}

func (v *grammarGenerator_) GenerateGrammarModel(
	module string,
	wiki string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = grammarGeneratorReference().modelTemplate_
	var notice = v.generateNotice()
	implementation = uti.ReplaceAll(implementation, "notice", notice)
	var packageDeclaration = v.generatePackageDeclaration(wiki)
	implementation = uti.ReplaceAll(
		implementation,
		"packageDeclaration",
		packageDeclaration,
	)
	var moduleImports = v.generateModuleImports(module)
	implementation = uti.ReplaceAll(
		implementation,
		"moduleImports",
		moduleImports,
	)
	var typeDefinitions = v.generateTypeDefinitions()
	implementation = uti.ReplaceAll(
		implementation,
		"typeDefinitions",
		typeDefinitions,
	)
	var classDefinitions = v.generateClassDefinitions()
	implementation = uti.ReplaceAll(
		implementation,
		"classDefinitions",
		classDefinitions,
	)
	var instanceDefinitions = v.generateInstanceDefinitions()
	implementation = uti.ReplaceAll(
		implementation,
		"instanceDefinitions",
		instanceDefinitions,
	)
	var aspectDefinitions = v.generateAspectDefinitions()
	implementation = uti.ReplaceAll(
		implementation,
		"aspectDefinitions",
		aspectDefinitions,
	)
	var syntaxName = v.analyzer_.GetSyntaxName()
	implementation = uti.ReplaceAll(
		implementation,
		"syntaxName",
		syntaxName,
	)
	return implementation
}

// PROTECTED INTERFACE

// Private Methods

func (v *grammarGenerator_) generateAspectDefinitions() (
	implementation string,
) {
	implementation = grammarGeneratorReference().aspectDefinitions_
	var processTokens = v.generateProcessTokens()
	implementation = uti.ReplaceAll(implementation, "processTokens", processTokens)
	var processRules = v.generateProcessRules()
	implementation = uti.ReplaceAll(implementation, "processRules", processRules)
	return implementation
}

func (v *grammarGenerator_) generateClassDefinitions() (
	implementation string,
) {
	implementation = grammarGeneratorReference().classDefinitions_
	return implementation
}

func (v *grammarGenerator_) generateInstanceDefinitions() (
	implementation string,
) {
	implementation = grammarGeneratorReference().instanceDefinitions_
	return implementation
}

func (v *grammarGenerator_) generateModuleImports(
	module string,
) (
	implementation string,
) {
	implementation = grammarGeneratorReference().moduleImports_
	implementation = uti.ReplaceAll(implementation, "Module", module)
	return implementation
}

func (v *grammarGenerator_) generateNotice() string {
	var notice = v.analyzer_.GetNotice()
	return notice
}

func (v *grammarGenerator_) generatePackageDeclaration(
	wiki string,
) (
	implementation string,
) {
	implementation = grammarGeneratorReference().packageDeclaration_
	implementation = uti.ReplaceAll(implementation, "wiki", wiki)
	return implementation
}

func (v *grammarGenerator_) generateProcessRule(
	ruleName string,
) (
	implementation string,
) {
	implementation = grammarGeneratorReference().processRule_
	if v.analyzer_.IsPlural(ruleName) {
		implementation = grammarGeneratorReference().processIndexedRule_
	}
	implementation = uti.ReplaceAll(implementation, "ruleName", ruleName)
	return implementation
}

func (v *grammarGenerator_) generateProcessRules() (
	implementation string,
) {
	var iterator = v.analyzer_.GetRuleNames().GetIterator()
	for iterator.HasNext() {
		var ruleName = iterator.GetNext()
		var processRule = v.generateProcessRule(ruleName)
		implementation += processRule
	}
	return implementation
}

func (v *grammarGenerator_) generateProcessToken(
	tokenName string,
) (
	implementation string,
) {
	if tokenName == "delimiter" {
		return implementation
	}
	implementation = grammarGeneratorReference().processToken_
	if v.analyzer_.IsPlural(tokenName) {
		implementation = grammarGeneratorReference().processIndexedToken_
	}
	implementation = uti.ReplaceAll(implementation, "tokenName", tokenName)
	return implementation
}

func (v *grammarGenerator_) generateProcessTokens() (
	implementation string,
) {
	var iterator = v.analyzer_.GetTokenNames().GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var processToken = v.generateProcessToken(tokenName)
		implementation += processToken
	}
	return implementation
}

func (v *grammarGenerator_) generateTokenTypes() (
	implementation string,
) {
	var iterator = v.analyzer_.GetTokenNames().GetIterator()
	for iterator.HasNext() {
		var tokenName = iterator.GetNext()
		var tokenType = grammarGeneratorReference().tokenType_
		tokenType = uti.ReplaceAll(tokenType, "tokenName", tokenName)
		implementation += tokenType
	}
	return implementation
}

func (v *grammarGenerator_) generateTypeDefinitions() (
	implementation string,
) {
	implementation = grammarGeneratorReference().typeDefinitions_
	var tokenTypes = v.generateTokenTypes()
	implementation = uti.ReplaceAll(implementation, "tokenTypes", tokenTypes)
	return implementation
}

// Instance Structure

type grammarGenerator_ struct {
	// Declare the instance attributes.
	analyzer_ AnalyzerLike
}

// Class Structure

type grammarGeneratorClass_ struct {
	// Declare the class constants.
	modelTemplate_       string
	packageDeclaration_  string
	moduleImports_       string
	typeDefinitions_     string
	tokenType_           string
	classDefinitions_    string
	instanceDefinitions_ string
	aspectDefinitions_   string
	processToken_        string
	processIndexedToken_ string
	processRule_         string
	processIndexedRule_  string
}

// Class Reference

func grammarGeneratorReference() *grammarGeneratorClass_ {
	return grammarGeneratorReference_
}

var grammarGeneratorReference_ = &grammarGeneratorClass_{
	// Initialize the class constants.
	modelTemplate_: `<Notice><PackageDeclaration><ModuleImports>

// Type Definitions<TypeDefinitions>

// Class Definitions<ClassDefinitions>

// Instance Definitions<InstanceDefinitions>

// Aspect Definitions<AspectDefinitions>
`,
	packageDeclaration_: `
/*
Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

For detailed documentation on this package refer to the wiki:
  - https://<wiki>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package grammar`,
	moduleImports_: `

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "<Module>/ast"
)`,

	typeDefinitions_: `

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota<TokenTypes>
)`,

	tokenType_: `
	<~TokenName>Token`,

	classDefinitions_: `

/*
FormatterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor Methods
	Make() FormatterLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor Methods
	Make() ParserLike
}

/*
ProcessorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor Methods
	Make() ProcessorLike
}

/*
ScannerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructor Methods
	Make(
		source string,
		tokens abs.QueueLike[TokenLike],
	) ScannerLike

	// Function Methods
	FormatToken(
		token TokenLike,
	) string
	FormatType(
		tokenType TokenType,
	) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructor Methods
	Make(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor Methods
	Make() ValidatorLike
}

/*
VisitorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor Methods
	Make(
		processor Methodical,
	) VisitorLike
}`,

	instanceDefinitions_: `

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Primary Methods
	GetClass() FormatterClassLike
	Format<~SyntaxName>(
		<syntaxName_> ast.<~SyntaxName>Like,
	) string

	// Aspect Interfaces
	Methodical
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Primary Methods
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.<~SyntaxName>Like
}

/*
ProcessorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Primary Methods
	GetClass() ProcessorClassLike

	// Aspect Interfaces
	Methodical
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Primary Methods
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Primary Methods
	GetClass() TokenClassLike

	// Attribute Methods
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Primary Methods
	GetClass() ValidatorClassLike
	Validate<~SyntaxName>(
		<syntaxName_> ast.<~SyntaxName>Like,
	)

	// Aspect Interfaces
	Methodical
}

/*
VisitorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Primary Methods
	GetClass() VisitorClassLike
	Visit<~SyntaxName>(
		<syntaxName_> ast.<~SyntaxName>Like,
	)
}`,

	aspectDefinitions_: `

/*
Methodical defines the set of method signatures that must be supported
by all methodical processors.
*/
type Methodical interface {<ProcessTokens><ProcessRules>
}`,

	processToken_: `
	Process<~TokenName>(
		<tokenName_> string,
	)`,

	processIndexedToken_: `
	Process<~TokenName>(
		<tokenName_> string,
		index uint,
		size uint,
	)`,

	processRule_: `
	Preprocess<~RuleName>(
		<ruleName_> ast.<~RuleName>Like,
	)
	Process<~RuleName>Slot(
		slot uint,
	)
	Postprocess<~RuleName>(
		<ruleName_> ast.<~RuleName>Like,
	)`,

	processIndexedRule_: `
	Preprocess<~RuleName>(
		<ruleName_> ast.<~RuleName>Like,
		index uint,
		size uint,
	)
	Process<~RuleName>Slot(
		slot uint,
	)
	Postprocess<~RuleName>(
		<ruleName_> ast.<~RuleName>Like,
		index uint,
		size uint,
	)`,
}

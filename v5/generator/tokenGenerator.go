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

func TokenGenerator() TokenGeneratorClassLike {
	return tokenGeneratorReference()
}

// Constructor Methods

func (c *tokenGeneratorClass_) Make() TokenGeneratorLike {
	var instance = &tokenGenerator_{
		// Initialize the instance attributes.
		analyzer_: Analyzer().Make(),
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *tokenGenerator_) GetClass() TokenGeneratorClassLike {
	return tokenGeneratorReference()
}

func (v *tokenGenerator_) GenerateTokenClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = tokenGeneratorReference().classTemplate_
	var notice = v.generateNotice()
	implementation = uti.ReplaceAll(implementation, "notice", notice)

	var packageDeclaration = v.generatePackageDeclaration()
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

	var accessFunction = v.generateAccessFunction()
	implementation = uti.ReplaceAll(
		implementation,
		"accessFunction",
		accessFunction,
	)

	var constructorMethods = v.generateConstructorMethods()
	implementation = uti.ReplaceAll(
		implementation,
		"constructorMethods",
		constructorMethods,
	)

	var attributeMethods = v.generateAttributeMethods()
	implementation = uti.ReplaceAll(
		implementation,
		"attributeMethods",
		attributeMethods,
	)

	var primaryMethods = v.generatePrimaryMethods()
	implementation = uti.ReplaceAll(
		implementation,
		"primaryMethods",
		primaryMethods,
	)

	var privateMethods = v.generatePrivateMethods()
	implementation = uti.ReplaceAll(
		implementation,
		"privateMethods",
		privateMethods,
	)

	var instanceStructure = v.generateInstanceStructure()
	implementation = uti.ReplaceAll(
		implementation,
		"instanceStructure",
		instanceStructure,
	)

	var classStructure = v.generateClassStructure()
	implementation = uti.ReplaceAll(
		implementation,
		"classStructure",
		classStructure,
	)

	var classReference = v.generateClassReference()
	implementation = uti.ReplaceAll(
		implementation,
		"classReference",
		classReference,
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

func (v *tokenGenerator_) generateAccessFunction() (
	implementation string,
) {
	implementation = tokenGeneratorReference().accessFunction_
	return implementation
}

func (v *tokenGenerator_) generateConstructorMethods() (
	implementation string,
) {
	implementation = tokenGeneratorReference().constructorMethods_
	return implementation
}

func (v *tokenGenerator_) generateAttributeMethods() (
	implementation string,
) {
	implementation = tokenGeneratorReference().attributeMethods_
	return implementation
}

func (v *tokenGenerator_) generatePrimaryMethods() (
	implementation string,
) {
	implementation = tokenGeneratorReference().primaryMethods_
	return implementation
}

func (v *tokenGenerator_) generatePrivateMethods() (
	implementation string,
) {
	implementation = tokenGeneratorReference().privateMethods_
	return implementation
}

func (v *tokenGenerator_) generateInstanceStructure() (
	implementation string,
) {
	implementation = tokenGeneratorReference().instanceStructure_
	return implementation
}

func (v *tokenGenerator_) generateClassStructure() (
	implementation string,
) {
	implementation = tokenGeneratorReference().classStructure_
	return implementation
}

func (v *tokenGenerator_) generateClassReference() (
	implementation string,
) {
	implementation = tokenGeneratorReference().classReference_
	return implementation
}

func (v *tokenGenerator_) generateModuleImports(
	module string,
) (
	implementation string,
) {
	implementation = tokenGeneratorReference().moduleImports_
	implementation = uti.ReplaceAll(implementation, "Module", module)
	return implementation
}

func (v *tokenGenerator_) generateNotice() string {
	var notice = v.analyzer_.GetNotice()
	return notice
}

func (v *tokenGenerator_) generatePackageDeclaration() (
	implementation string,
) {
	implementation = tokenGeneratorReference().packageDeclaration_
	return implementation
}

// Instance Structure

type tokenGenerator_ struct {
	// Declare the instance attributes.
	analyzer_ AnalyzerLike
}

// Class Structure

type tokenGeneratorClass_ struct {
	// Declare the class constants.
	classTemplate_      string
	packageDeclaration_ string
	moduleImports_      string
	accessFunction_     string
	constructorMethods_ string
	primaryMethods_     string
	attributeMethods_   string
	privateMethods_     string
	instanceStructure_  string
	classStructure_     string
	classReference_     string
}

// Class Reference

func tokenGeneratorReference() *tokenGeneratorClass_ {
	return tokenGeneratorReference_
}

var tokenGeneratorReference_ = &tokenGeneratorClass_{
	// Initialize the class constants.
	classTemplate_: `<Notice><PackageDeclaration><ModuleImports>

// CLASS INTERFACE

// Access Function<AccessFunction>

// Constructor Methods<ConstructorMethods>

// INSTANCE INTERFACE

// Primary Methods<PrimaryMethods>

// Attribute Methods<AttributeMethods>

// PROTECTED INTERFACE

// Private Methods<PrivateMethods>

// Instance Structure<InstanceStructure>

// Class Structure<ClassStructure>

// Class Reference<ClassReference>
`,

	packageDeclaration_: `
package grammar`,

	moduleImports_: `

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)`,

	accessFunction_: `

func Token() TokenClassLike {
	return tokenReference()
}`,

	constructorMethods_: `

func (c *tokenClass_) Make(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) TokenLike {
	if uti.IsUndefined(line) {
		panic("The \"line\" attribute is required by this class.")
	}
	if uti.IsUndefined(position) {
		panic("The \"position\" attribute is required by this class.")
	}
	if uti.IsUndefined(type_) {
		panic("The \"type\" attribute is required by this class.")
	}
	if uti.IsUndefined(value) {
		panic("The \"value\" attribute is required by this class.")
	}
	var instance = &token_{
		// Initialize the instance attributes.
		line_:     line,
		position_: position,
		type_:     type_,
		value_:    value,
	}
	return instance
}`,

	primaryMethods_: `

func (v *token_) GetClass() TokenClassLike {
	return tokenReference()
}`,

	attributeMethods_: `
// Attribute Methods

func (v *token_) GetLine() uint {
	return v.line_
}

func (v *token_) GetPosition() uint {
	return v.position_
}

func (v *token_) GetType() TokenType {
	return v.type_
}

func (v *token_) GetValue() string {
	return v.value_
}`,

	privateMethods_: ``,

	instanceStructure_: `

type token_ struct {
	// Declare the instance attributes.
	line_     uint
	position_ uint
	type_     TokenType
	value_    string
}`,

	classStructure_: `

type tokenClass_ struct {
	// Declare the class constants.
}`,

	classReference_: `

func tokenReference() *tokenClass_ {
	return tokenReference_
}

var tokenReference_ = &tokenClass_{
	// Initialize the class constants.
}`,
}

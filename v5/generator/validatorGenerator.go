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

func ValidatorGenerator() ValidatorGeneratorClassLike {
	return validatorGeneratorReference()
}

// Constructor Methods

func (c *validatorGeneratorClass_) Make() ValidatorGeneratorLike {
	var instance = &validatorGenerator_{
		// Initialize the instance attributes.
		analyzer_: Analyzer().Make(),
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *validatorGenerator_) GetClass() ValidatorGeneratorClassLike {
	return validatorGeneratorReference()
}

func (v *validatorGenerator_) GenerateValidatorClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = validatorGeneratorReference().classTemplate_
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

	var methodicalMethods = v.generateMethodicalMethods()
	implementation = uti.ReplaceAll(
		implementation,
		"methodicalMethods",
		methodicalMethods,
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

func (v *validatorGenerator_) generateAccessFunction() (
	implementation string,
) {
	implementation = validatorGeneratorReference().accessFunction_
	return implementation
}

func (v *validatorGenerator_) generateConstructorMethods() (
	implementation string,
) {
	implementation = validatorGeneratorReference().constructorMethods_
	return implementation
}

func (v *validatorGenerator_) generateMethodicalMethods() (
	implementation string,
) {
	implementation = validatorGeneratorReference().methodicalMethods_
	var processTokens = v.generateProcessTokens()
	implementation = uti.ReplaceAll(implementation, "processTokens", processTokens)
	var processRules = v.generateProcessRules()
	implementation = uti.ReplaceAll(implementation, "processRules", processRules)
	return implementation
}

func (v *validatorGenerator_) generatePrimaryMethods() (
	implementation string,
) {
	implementation = validatorGeneratorReference().primaryMethods_
	return implementation
}

func (v *validatorGenerator_) generatePrivateMethods() (
	implementation string,
) {
	implementation = validatorGeneratorReference().privateMethods_
	return implementation
}

func (v *validatorGenerator_) generateInstanceStructure() (
	implementation string,
) {
	implementation = validatorGeneratorReference().instanceStructure_
	return implementation
}

func (v *validatorGenerator_) generateClassStructure() (
	implementation string,
) {
	implementation = validatorGeneratorReference().classStructure_
	return implementation
}

func (v *validatorGenerator_) generateClassReference() (
	implementation string,
) {
	implementation = validatorGeneratorReference().classReference_
	return implementation
}

func (v *validatorGenerator_) generateModuleImports(
	module string,
) (
	implementation string,
) {
	implementation = validatorGeneratorReference().moduleImports_
	implementation = uti.ReplaceAll(implementation, "Module", module)
	return implementation
}

func (v *validatorGenerator_) generateNotice() string {
	var notice = v.analyzer_.GetNotice()
	return notice
}

func (v *validatorGenerator_) generatePackageDeclaration() (
	implementation string,
) {
	implementation = validatorGeneratorReference().packageDeclaration_
	return implementation
}

func (v *validatorGenerator_) generateProcessRule(
	ruleName string,
) (
	implementation string,
) {
	implementation = validatorGeneratorReference().processRule_
	if v.analyzer_.IsPlural(ruleName) {
		implementation = validatorGeneratorReference().processIndexedRule_
	}
	implementation = uti.ReplaceAll(implementation, "ruleName", ruleName)
	return implementation
}

func (v *validatorGenerator_) generateProcessRules() (
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

func (v *validatorGenerator_) generateProcessToken(
	tokenName string,
) (
	implementation string,
) {
	if tokenName == "delimiter" {
		return implementation
	}
	implementation = validatorGeneratorReference().processToken_
	if v.analyzer_.IsPlural(tokenName) {
		implementation = validatorGeneratorReference().processIndexedToken_
	}
	implementation = uti.ReplaceAll(implementation, "tokenName", tokenName)
	return implementation
}

func (v *validatorGenerator_) generateProcessTokens() (
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

// Instance Structure

type validatorGenerator_ struct {
	// Declare the instance attributes.
	analyzer_ AnalyzerLike
}

// Class Structure

type validatorGeneratorClass_ struct {
	// Declare the class constants.
	classTemplate_       string
	packageDeclaration_  string
	moduleImports_       string
	accessFunction_      string
	constructorMethods_  string
	primaryMethods_      string
	methodicalMethods_   string
	processToken_        string
	processIndexedToken_ string
	processRule_         string
	processIndexedRule_  string
	privateMethods_      string
	instanceStructure_   string
	classStructure_      string
	classReference_      string
}

// Class Reference

func validatorGeneratorReference() *validatorGeneratorClass_ {
	return validatorGeneratorReference_
}

var validatorGeneratorReference_ = &validatorGeneratorClass_{
	// Initialize the class constants.
	classTemplate_: `<Notice><PackageDeclaration><ModuleImports>

// CLASS INTERFACE

// Access Function<AccessFunction>

// Constructor Methods<ConstructorMethods>

// INSTANCE INTERFACE

// Primary Methods<PrimaryMethods>

// Methodical Methods<MethodicalMethods>

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
	fmt "fmt"
	ast "<Module>/ast"
)`,

	accessFunction_: `

func Validator() ValidatorClassLike {
	return validatorReference()
}`,

	constructorMethods_: `

func (c *validatorClass_) Make() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	instance.visitor_ = Visitor().Make(instance)
	return instance
}`,

	primaryMethods_: `

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorReference()
}

func (v *validator_) Validate<~SyntaxName>(
	<syntaxName_> ast.<~SyntaxName>Like,
) {
	v.visitor_.Visit<~SyntaxName>(<syntaxName_>)
}`,

	methodicalMethods_: `<ProcessTokens><ProcessRules>`,

	processToken_: `

func (v *validator_) Process<~TokenName>(
	<tokenName_> string,
) {
	v.validateToken(<tokenName_>, <~TokenName>Token)
}`,

	processIndexedToken_: `

func (v *validator_) Process<~TokenName>(
	<tokenName_> string,
	index uint,
	size uint,
) {
	v.validateToken(<tokenName_>, <~TokenName>Token)
}`,

	processRule_: `

func (v *validator_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add any validation checks.
}`,

	processIndexedRule_: `

func (v *validator_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}`,

	privateMethods_: `

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	if !Scanner().MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			Scanner().FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}`,

	instanceStructure_: `

type validator_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike

	// Declare the inherited aspects.
	Methodical
}`,

	classStructure_: `

type validatorClass_ struct {
	// Declare the class constants.
}`,

	classReference_: `

func validatorReference() *validatorClass_ {
	return validatorReference_
}

var validatorReference_ = &validatorClass_{
	// Initialize the class constants.
}`,
}

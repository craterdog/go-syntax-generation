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

func ProcessorGenerator() ProcessorGeneratorClassLike {
	return processorGeneratorReference()
}

// Constructor Methods

func (c *processorGeneratorClass_) Make() ProcessorGeneratorLike {
	var instance = &processorGenerator_{
		// Initialize the instance attributes.
		analyzer_: Analyzer().Make(),
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *processorGenerator_) GetClass() ProcessorGeneratorClassLike {
	return processorGeneratorReference()
}

func (v *processorGenerator_) GenerateProcessorClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = processorGeneratorReference().classTemplate_
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

func (v *processorGenerator_) generateAccessFunction() (
	implementation string,
) {
	implementation = processorGeneratorReference().accessFunction_
	return implementation
}

func (v *processorGenerator_) generateConstructorMethods() (
	implementation string,
) {
	implementation = processorGeneratorReference().constructorMethods_
	return implementation
}

func (v *processorGenerator_) generateMethodicalMethods() (
	implementation string,
) {
	implementation = processorGeneratorReference().methodicalMethods_
	var processTokens = v.generateProcessTokens()
	implementation = uti.ReplaceAll(implementation, "processTokens", processTokens)
	var processRules = v.generateProcessRules()
	implementation = uti.ReplaceAll(implementation, "processRules", processRules)
	return implementation
}

func (v *processorGenerator_) generatePrimaryMethods() (
	implementation string,
) {
	implementation = processorGeneratorReference().primaryMethods_
	return implementation
}

func (v *processorGenerator_) generatePrivateMethods() (
	implementation string,
) {
	implementation = processorGeneratorReference().privateMethods_
	return implementation
}

func (v *processorGenerator_) generateInstanceStructure() (
	implementation string,
) {
	implementation = processorGeneratorReference().instanceStructure_
	return implementation
}

func (v *processorGenerator_) generateClassStructure() (
	implementation string,
) {
	implementation = processorGeneratorReference().classStructure_
	return implementation
}

func (v *processorGenerator_) generateClassReference() (
	implementation string,
) {
	implementation = processorGeneratorReference().classReference_
	return implementation
}

func (v *processorGenerator_) generateModuleImports(
	module string,
) (
	implementation string,
) {
	implementation = processorGeneratorReference().moduleImports_
	implementation = uti.ReplaceAll(implementation, "Module", module)
	return implementation
}

func (v *processorGenerator_) generateNotice() string {
	var notice = v.analyzer_.GetNotice()
	return notice
}

func (v *processorGenerator_) generatePackageDeclaration() (
	implementation string,
) {
	implementation = processorGeneratorReference().packageDeclaration_
	return implementation
}

func (v *processorGenerator_) generateProcessRule(
	ruleName string,
) (
	implementation string,
) {
	implementation = processorGeneratorReference().processRule_
	if v.analyzer_.IsPlural(ruleName) {
		implementation = processorGeneratorReference().processIndexedRule_
	}
	implementation = uti.ReplaceAll(implementation, "ruleName", ruleName)
	return implementation
}

func (v *processorGenerator_) generateProcessRules() (
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

func (v *processorGenerator_) generateProcessToken(
	tokenName string,
) (
	implementation string,
) {
	if tokenName == "delimiter" {
		return implementation
	}
	implementation = processorGeneratorReference().processToken_
	if v.analyzer_.IsPlural(tokenName) {
		implementation = processorGeneratorReference().processIndexedToken_
	}
	implementation = uti.ReplaceAll(implementation, "tokenName", tokenName)
	return implementation
}

func (v *processorGenerator_) generateProcessTokens() (
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

type processorGenerator_ struct {
	// Declare the instance attributes.
	analyzer_ AnalyzerLike
}

// Class Structure

type processorGeneratorClass_ struct {
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

func processorGeneratorReference() *processorGeneratorClass_ {
	return processorGeneratorReference_
}

var processorGeneratorReference_ = &processorGeneratorClass_{
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
	ast "<Module>/ast"
)`,

	accessFunction_: `

func Processor() ProcessorClassLike {
	return processorReference()
}`,

	constructorMethods_: `

func (c *processorClass_) Make() ProcessorLike {
	var instance = &processor_{
		// Initialize the instance attributes.
	}
	return instance
}`,

	primaryMethods_: `

func (v *processor_) GetClass() ProcessorClassLike {
	return processorReference()
}`,

	methodicalMethods_: `<ProcessTokens><ProcessRules>`,

	processToken_: `

func (v *processor_) Process<~TokenName>(
	<tokenName_> string,
) {
}`,

	processIndexedToken_: `

func (v *processor_) Process<~TokenName>(
	<tokenName_> string,
	index uint,
	size uint,
) {
}`,

	processRule_: `

func (v *processor_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
}

func (v *processor_) Process<~RuleName>Slot(
	slot uint,
) {
}

func (v *processor_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
}`,

	processIndexedRule_: `

func (v *processor_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
}

func (v *processor_) Process<~RuleName>Slot(
	slot uint,
) {
}

func (v *processor_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
}`,

	privateMethods_: ``,

	instanceStructure_: `

type processor_ struct {
	// Declare the instance attributes.
}`,

	classStructure_: `

type processorClass_ struct {
	// Declare the class constants.
}`,

	classReference_: `

func processorReference() *processorClass_ {
	return processorReference_
}

var processorReference_ = &processorClass_{
	// Initialize the class constants.
}`,
}

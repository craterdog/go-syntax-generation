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

func AstGenerator() AstGeneratorClassLike {
	return astGeneratorReference()
}

// Constructor Methods

func (c *astGeneratorClass_) Make() AstGeneratorLike {
	var instance = &astGenerator_{
		// Initialize the instance attributes.
		analyzer_: Analyzer().Make(),
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *astGenerator_) GetClass() AstGeneratorClassLike {
	return astGeneratorReference()
}

func (v *astGenerator_) GenerateAstModel(
	wiki string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = astGeneratorReference().modelTemplate_
	var notice = v.generateNotice()
	implementation = uti.ReplaceAll(implementation, "notice", notice)
	var packageDeclaration = v.generatePackageDeclaration(wiki)
	implementation = uti.ReplaceAll(
		implementation,
		"packageDeclaration",
		packageDeclaration,
	)
	var moduleImports = v.generateModuleImports()
	implementation = uti.ReplaceAll(
		implementation,
		"moduleImports",
		moduleImports,
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
	return implementation
}

// PROTECTED INTERFACE

// Private Methods

func (v *astGenerator_) generateParameter(
	isPlural bool,
	parameterName string,
	parameterType string,
) (
	parameter string,
) {
	parameter = astGeneratorReference().singularRuleParameter_
	if parameterType == "string" {
		parameter = astGeneratorReference().singularTokenParameter_
		if isPlural {
			parameter = astGeneratorReference().pluralTokenParameter_
		}
	} else {
		if isPlural {
			parameter = astGeneratorReference().pluralRuleParameter_
		}
	}
	parameter = uti.ReplaceAll(parameter, "parameterName", parameterName)
	parameter = uti.ReplaceAll(parameter, "parameterType", parameterType)
	return parameter
}

func (v *astGenerator_) generateClassDefinition(
	className string,
) (
	implementation string,
) {
	var parameters string
	var variables = v.analyzer_.GetVariables(className)
	if uti.IsDefined(variables) {
		// This class represents an inline rule.
		var attributes = variables.GetIterator()
		for attributes.HasNext() {
			var attribute = attributes.GetNext()
			var isPlural = v.isPlural(attribute)
			var attributeName = v.analyzer_.GetVariableName(attribute)
			var attributeType = v.analyzer_.GetVariableType(attribute)
			parameters += v.generateParameter(
				isPlural,
				attributeName,
				attributeType,
			)
		}
		if attributes.GetSize() > 0 {
			parameters += "\n\t"
		}
	} else {
		// This class represents a multiline rule.
		parameters += "\n\t\tany_ any,\n\t"
	}
	implementation = astGeneratorReference().classDefinition_
	implementation = uti.ReplaceAll(implementation, "parameters", parameters)
	implementation = uti.ReplaceAll(implementation, "className", className)
	return implementation
}

func (v *astGenerator_) generateClassDefinitions() (
	implementation string,
) {
	var rules = v.analyzer_.GetRuleNames().GetIterator()
	for rules.HasNext() {
		var className = rules.GetNext()
		implementation += v.generateClassDefinition(className)
	}
	return implementation
}

func (v *astGenerator_) generateGetterMethod(
	isPlural bool,
	attributeName string,
	attributeType string,
) (
	implementation string,
) {
	implementation = astGeneratorReference().ruleGetterMethod_
	if attributeType == "string" {
		implementation = astGeneratorReference().tokenGetterMethod_
		if isPlural {
			implementation = astGeneratorReference().pluralTokenGetterMethod_
		}
	} else {
		if isPlural {
			implementation = astGeneratorReference().pluralRuleGetterMethod_
		}
	}
	implementation = uti.ReplaceAll(implementation, "attributeName", attributeName)
	implementation = uti.ReplaceAll(implementation, "attributeType", attributeType)
	return implementation
}

func (v *astGenerator_) generatePackageDeclaration(
	wiki string,
) (
	implementation string,
) {
	implementation = astGeneratorReference().packageDeclaration_
	implementation = uti.ReplaceAll(implementation, "wiki", wiki)
	return implementation
}

func (v *astGenerator_) generateModuleImports() (
	implementation string,
) {
	if v.analyzer_.HasPlurals() {
		implementation = astGeneratorReference().moduleImports_
	}
	return implementation
}

func (v *astGenerator_) generateInstanceDefinition(
	className string,
) (
	implementation string,
) {
	var getterMethods string
	var variables = v.analyzer_.GetVariables(className)
	if uti.IsDefined(variables) {
		// This instance represents an inline rule.
		var attributes = variables.GetIterator()
		for attributes.HasNext() {
			var attribute = attributes.GetNext()
			var isPlural = v.isPlural(attribute)
			var attributeName = v.analyzer_.GetVariableName(attribute)
			var attributeType = v.analyzer_.GetVariableType(attribute)
			getterMethods += v.generateGetterMethod(
				isPlural,
				attributeName,
				attributeType,
			)
		}
	} else {
		// This instance represents a multiline rule.
		getterMethods += "\n\tGetAny() any"
	}
	implementation = astGeneratorReference().instanceDefinition_
	implementation = uti.ReplaceAll(
		implementation,
		"primaryMethods",
		astGeneratorReference().primaryMethods_,
	)
	var template string
	if uti.IsDefined(getterMethods) {
		template = astGeneratorReference().attributeMethods_
		template = uti.ReplaceAll(template, "getterMethods", getterMethods)
	}
	implementation = uti.ReplaceAll(implementation, "attributeMethods", template)
	implementation = uti.ReplaceAll(implementation, "className", className)
	return implementation
}

func (v *astGenerator_) generateInstanceDefinitions() (
	implementation string,
) {
	var rules = v.analyzer_.GetRuleNames().GetIterator()
	for rules.HasNext() {
		var className = rules.GetNext()
		implementation += v.generateInstanceDefinition(className)
	}
	return implementation
}

func (v *astGenerator_) generateNotice() string {
	var notice = v.analyzer_.GetNotice()
	return notice
}

func (v *astGenerator_) isPlural(reference ast.ReferenceLike) bool {
	var cardinality = reference.GetOptionalCardinality()
	if uti.IsUndefined(cardinality) {
		return false
	}
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		if actual.GetAny().(string) == "?" {
			return false
		}
	}
	return true
}

// Instance Structure

type astGenerator_ struct {
	// Declare the instance attributes.
	analyzer_ AnalyzerLike
}

// Class Structure

type astGeneratorClass_ struct {
	// Declare the class constants.
	modelTemplate_           string
	packageDeclaration_      string
	moduleImports_           string
	classDefinition_         string
	singularRuleParameter_   string
	pluralRuleParameter_     string
	singularTokenParameter_  string
	pluralTokenParameter_    string
	instanceDefinition_      string
	primaryMethods_          string
	attributeMethods_        string
	ruleGetterMethod_        string
	pluralRuleGetterMethod_  string
	tokenGetterMethod_       string
	pluralTokenGetterMethod_ string
}

// Class Reference

func astGeneratorReference() *astGeneratorClass_ {
	return astGeneratorReference_
}

var astGeneratorReference_ = &astGeneratorClass_{
	// Initialize the class constants.
	modelTemplate_: `<Notice><PackageDeclaration><ModuleImports>

// Class Definitions<ClassDefinitions>

// Instance Definitions<InstanceDefinitions>
`,
	packageDeclaration_: `
/*
Package "ast" provides the abstract syntax tree (AST) classes for this module.
Each AST class manages the attributes associated with the rule definition found
in the syntax grammar with the same rule name as the class.

For detailed documentation on this package refer to the wiki:
  - https://<wiki>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki
*/
package ast`,
	moduleImports_: `

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)`,

	classDefinition_: `

/*
<~ClassName>ClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete <~class-name>-like class.
*/
type <~ClassName>ClassLike interface {
	// Constructor Methods
	Make(<Parameters>) <~ClassName>Like
}`,
	singularRuleParameter_: `
		<parameterName_> <ParameterType>,`,
	pluralRuleParameter_: `
		<parameterName_> abs.Sequential[<ParameterType>],`,
	singularTokenParameter_: `
		<parameterName_> string,`,
	pluralTokenParameter_: `
		<parameterName_> abs.Sequential[string],`,
	instanceDefinition_: `

/*
<~ClassName>Like is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete <~class-name>-like class.
*/
type <~ClassName>Like interface {<PrimaryMethods><AttributeMethods>}`,
	primaryMethods_: `
	// Primary Methods
	GetClass() <~ClassName>ClassLike
`,
	attributeMethods_: `
	// Attribute Methods<GetterMethods>
`,
	ruleGetterMethod_: `
	Get<~AttributeName>() <AttributeType>`,
	pluralRuleGetterMethod_: `
	Get<~AttributeName>() abs.Sequential[<AttributeType>]`,
	tokenGetterMethod_: `
	Get<~AttributeName>() string`,
	pluralTokenGetterMethod_: `
	Get<~AttributeName>() abs.Sequential[string]`,
}

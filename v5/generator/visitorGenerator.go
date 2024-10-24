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
	not "github.com/craterdog/go-syntax-notation/v5"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func VisitorGenerator() VisitorGeneratorClassLike {
	return visitorGeneratorReference()
}

// Constructor Methods

func (c *visitorGeneratorClass_) Make() VisitorGeneratorLike {
	var instance = &visitorGenerator_{
		// Initialize the instance attributes.
		analyzer_: Analyzer().Make(),
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *visitorGenerator_) GetClass() VisitorGeneratorClassLike {
	return visitorGeneratorReference()
}

func (v *visitorGenerator_) GenerateVisitorClass(
	module string,
	syntax not.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = visitorGeneratorReference().classTemplate_
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

func (v *visitorGenerator_) generateAccessFunction() (
	implementation string,
) {
	implementation = visitorGeneratorReference().accessFunction_
	return implementation
}

func (v *visitorGenerator_) generateConstructorMethods() (
	implementation string,
) {
	implementation = visitorGeneratorReference().constructorMethods_
	return implementation
}

func (v *visitorGenerator_) generatePlurality(
	reference not.ReferenceLike,
) (
	plurality string,
) {
	var name = reference.GetIdentifier().GetAny().(string)
	var cardinality = reference.GetOptionalCardinality()
	if uti.IsUndefined(cardinality) {
		if v.analyzer_.IsPlural(name) {
			plurality = "singular"
		}
		return plurality
	}
	switch actual := cardinality.GetAny().(type) {
	case not.ConstrainedLike:
		var token = actual.GetAny().(string)
		switch {
		case not.MatchesType(token, not.OptionalToken):
			plurality = "optional"
			if v.analyzer_.IsPlural(name) {
				plurality = "singular"
			}
		case not.MatchesType(token, not.RepeatedToken):
			plurality = "repeated"
		}
	case not.QuantifiedLike:
		plurality = "repeated"
	}
	return plurality
}

func (v *visitorGenerator_) generatePrimaryMethods() (
	implementation string,
) {
	implementation = visitorGeneratorReference().primaryMethods_
	return implementation
}

func (v *visitorGenerator_) generatePrivateMethods() (
	implementation string,
) {
	var ruleNames = v.analyzer_.GetRuleNames().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		var visitMethod = v.generateVisitMethod(ruleName)
		implementation += visitMethod
	}
	return implementation
}

func (v *visitorGenerator_) generateVisitMethod(
	ruleName string,
) (
	implementation string,
) {
	var methodImplementation string
	switch {
	case uti.IsDefined(v.analyzer_.GetIdentifiers(ruleName)):
		methodImplementation = v.generateMultilineImplementation(ruleName)
	case uti.IsDefined(v.analyzer_.GetReferences(ruleName)):
		methodImplementation = v.generateInlineImplementation(ruleName)
	}
	implementation = visitorGeneratorReference().visitMethod_
	implementation = uti.ReplaceAll(
		implementation,
		"methodImplementation",
		methodImplementation,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"targetName",
		ruleName,
	)
	return implementation
}

func (v *visitorGenerator_) generateMultilineImplementation(
	ruleName string,
) (
	implementation string,
) {
	var tokenCases, ruleCases string
	var identifiers = v.analyzer_.GetIdentifiers(ruleName).GetIterator()
	for identifiers.HasNext() {
		var identifier = identifiers.GetNext().GetAny().(string)
		switch {
		case not.MatchesType(identifier, not.LowercaseToken):
			tokenCases += v.generateMultilineToken(identifier)
		case not.MatchesType(identifier, not.UppercaseToken):
			ruleCases += v.generateMultilineRule(identifier)
		}
	}
	implementation = visitorGeneratorReference().multilineCases_
	implementation = uti.ReplaceAll(implementation, "ruleCases", ruleCases)
	implementation = uti.ReplaceAll(implementation, "tokenCases", tokenCases)
	return implementation
}

func (v *visitorGenerator_) generateMultilineRule(
	ruleName string,
) (
	implementation string,
) {
	implementation = visitorGeneratorReference().ruleCase_
	if v.analyzer_.IsPlural(ruleName) {
		implementation = visitorGeneratorReference().singularRuleCase_
	}
	implementation = uti.ReplaceAll(
		implementation,
		"ruleName",
		ruleName,
	)
	return implementation
}

func (v *visitorGenerator_) generateMultilineToken(
	tokenName string,
) (
	implementation string,
) {
	implementation = visitorGeneratorReference().tokenCase_
	if v.analyzer_.IsPlural(tokenName) {
		implementation = visitorGeneratorReference().singularTokenCase_
	}
	implementation = uti.ReplaceAll(
		implementation,
		"tokenName",
		tokenName,
	)
	return implementation
}

func (v *visitorGenerator_) generateInlineImplementation(
	ruleName string,
) (
	implementation string,
) {
	var references = v.analyzer_.GetReferences(ruleName).GetIterator()
	var variables = v.analyzer_.GetVariables(ruleName).GetIterator()
	for references.HasNext() && variables.HasNext() {
		var slot uint = uint(references.GetSlot())
		if slot > 0 {
			implementation += v.generateInlineSlot(ruleName, slot)
		}
		var reference = references.GetNext()
		var variable = variables.GetNext()
		var variableName = v.analyzer_.GetVariableName(variable)
		implementation += v.generateInlineReference(reference, variableName)
	}
	return implementation
}

func (v *visitorGenerator_) generateInlineReference(
	reference not.ReferenceLike,
	variableName string,
) (
	implementation string,
) {
	var identifier = reference.GetIdentifier().GetAny().(string)
	switch {
	case not.MatchesType(identifier, not.LowercaseToken):
		implementation = v.generateInlineToken(reference, variableName)
	case not.MatchesType(identifier, not.UppercaseToken):
		implementation = v.generateInlineRule(reference, variableName)
	}
	return implementation
}

func (v *visitorGenerator_) generateInlineRule(
	reference not.ReferenceLike,
	variableName string,
) (
	implementation string,
) {
	switch v.generatePlurality(reference) {
	case "singular":
		implementation = visitorGeneratorReference().singularRuleBlock_
	case "optional":
		implementation = visitorGeneratorReference().optionalRuleBlock_
	case "repeated":
		implementation = visitorGeneratorReference().repeatedRuleBlock_
	default:
		implementation = visitorGeneratorReference().ruleBlock_
	}
	var ruleName = reference.GetIdentifier().GetAny().(string)
	implementation = uti.ReplaceAll(implementation, "ruleName", ruleName)
	implementation = uti.ReplaceAll(implementation, "variableName", variableName)
	return implementation
}

func (v *visitorGenerator_) generateInlineToken(
	reference not.ReferenceLike,
	variableName string,
) (
	implementation string,
) {
	switch v.generatePlurality(reference) {
	case "singular":
		implementation = visitorGeneratorReference().singularTokenBlock_
	case "optional":
		implementation = visitorGeneratorReference().optionalTokenBlock_
	case "repeated":
		implementation = visitorGeneratorReference().repeatedTokenBlock_
	default:
		implementation = visitorGeneratorReference().tokenBlock_
	}
	var tokenName = reference.GetIdentifier().GetAny().(string)
	implementation = uti.ReplaceAll(implementation, "tokenName", tokenName)
	implementation = uti.ReplaceAll(implementation, "variableName", variableName)
	return implementation
}

func (v *visitorGenerator_) generateInlineSlot(
	ruleName string,
	slot uint,
) (
	implementation string,
) {
	implementation = visitorGeneratorReference().slotBlock_
	implementation = uti.ReplaceAll(implementation, "slot", stc.Itoa(int(slot)))
	return implementation
}

func (v *visitorGenerator_) generateInstanceStructure() (
	implementation string,
) {
	implementation = visitorGeneratorReference().instanceStructure_
	return implementation
}

func (v *visitorGenerator_) generateClassStructure() (
	implementation string,
) {
	implementation = visitorGeneratorReference().classStructure_
	return implementation
}

func (v *visitorGenerator_) generateClassReference() (
	implementation string,
) {
	implementation = visitorGeneratorReference().classReference_
	return implementation
}

func (v *visitorGenerator_) generateModuleImports(
	module string,
) (
	implementation string,
) {
	implementation = visitorGeneratorReference().moduleImports_
	implementation = uti.ReplaceAll(implementation, "Module", module)
	return implementation
}

func (v *visitorGenerator_) generateNotice() string {
	var notice = v.analyzer_.GetNotice()
	return notice
}

func (v *visitorGenerator_) generatePackageDeclaration() (
	implementation string,
) {
	implementation = visitorGeneratorReference().packageDeclaration_
	return implementation
}

// Instance Structure

type visitorGenerator_ struct {
	// Declare the instance attributes.
	analyzer_ AnalyzerLike
}

// Class Structure

type visitorGeneratorClass_ struct {
	// Declare the class constants.
	classTemplate_      string
	packageDeclaration_ string
	moduleImports_      string
	accessFunction_     string
	constructorMethods_ string
	primaryMethods_     string
	visitMethod_        string
	multilineCases_     string
	ruleCase_           string
	singularRuleCase_   string
	tokenCase_          string
	singularTokenCase_  string
	ruleBlock_          string
	singularRuleBlock_  string
	optionalRuleBlock_  string
	repeatedRuleBlock_  string
	tokenBlock_         string
	singularTokenBlock_ string
	optionalTokenBlock_ string
	repeatedTokenBlock_ string
	slotBlock_          string
	instanceStructure_  string
	classStructure_     string
	classReference_     string
}

// Class Reference

func visitorGeneratorReference() *visitorGeneratorClass_ {
	return visitorGeneratorReference_
}

var visitorGeneratorReference_ = &visitorGeneratorClass_{
	// Initialize the class constants.
	classTemplate_: `<Notice><PackageDeclaration><ModuleImports>

// CLASS INTERFACE

// Access Function<AccessFunction>

// Constructor Methods<ConstructorMethods>

// INSTANCE INTERFACE

// Primary Methods<PrimaryMethods>

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
	uti "github.com/craterdog/go-missing-utilities/v2"
)`,

	accessFunction_: `

func Visitor() VisitorClassLike {
	return visitorReference()
}`,

	constructorMethods_: `

func (c *visitorClass_) Make(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		processor_: processor,
	}
	return instance
}`,

	primaryMethods_: `

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorReference()
}

func (v *visitor_) Visit<~SyntaxName>(
	<syntaxName_> ast.<~SyntaxName>Like,
) {
	v.processor_.Preprocess<~SyntaxName>(<syntaxName_>)
	v.visit<~SyntaxName>(<syntaxName_>)
	v.processor_.Postprocess<~SyntaxName>(<syntaxName_>)
}`,

	visitMethod_: `

func (v *visitor_) visit<~TargetName>(<targetName_> ast.<~TargetName>Like) {<MethodImplementation>}`,

	multilineCases_: `
	// Visit the possible <~targetName> types.
	switch actual := <targetName_>.GetAny().(type) {<RuleCases>
	case string:
		switch {<TokenCases>
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
`,

	ruleCase_: `
	case ast.<~RuleName>Like:
		v.processor_.Preprocess<~RuleName>(actual)
		v.visit<~RuleName>(actual)
		v.processor_.Postprocess<~RuleName>(actual)`,

	singularRuleCase_: `
	case ast.<~RuleName>Like:
		v.processor_.Preprocess<~RuleName>(actual, 1, 1)
		v.visit<~RuleName>(actual)
		v.processor_.Postprocess<~RuleName>(actual, 1, 1)`,

	tokenCase_: `
		case Scanner().MatchesType(actual, <~TokenName>Token):
			v.processor_.Process<~TokenName>(actual)`,

	singularTokenCase_: `
	case Scanner().MatchesType(actual, <~TokenName>Token):
			v.processor_.Process<~TokenName>(actual, 1, 1)`,

	ruleBlock_: `
	// Visit a single <~ruleName> rule.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	v.processor_.Preprocess<~RuleName>(<variableName_>)
	v.visit<~RuleName>(<variableName_>)
	v.processor_.Postprocess<~RuleName>(<variableName_>)
`,

	singularRuleBlock_: `
	// Visit a single <~ruleName> rule.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.Preprocess<~RuleName>(<variableName_>, 1, 1)
		v.visit<~RuleName>(<variableName_>)
		v.processor_.Postprocess<~RuleName>(<variableName_>, 1, 1)
	}
`,

	optionalRuleBlock_: `
	// Visit an optional <~ruleName> rule.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.Preprocess<~RuleName>(<variableName_>)
		v.visit<~RuleName>(<variableName_>)
		v.processor_.Postprocess<~RuleName>(<variableName_>)
	}
`,

	repeatedRuleBlock_: `
	// Visit each <~ruleName> rule.
	var <~ruleName>Index uint
	var <variableName_> = <targetName_>.Get<~VariableName>().GetIterator()
	var <variableName>Size = uint(<variableName_>.GetSize())
	for <variableName_>.HasNext() {
		<~ruleName>Index++
		var <ruleName_> = <variableName_>.GetNext()
		v.processor_.Preprocess<~RuleName>(
			<ruleName_>,
			<~ruleName>Index,
			<variableName>Size,
		)
		v.visit<~RuleName>(<ruleName_>)
		v.processor_.Postprocess<~RuleName>(
			<ruleName_>,
			<~ruleName>Index,
			<variableName>Size,
		)
	}
`,

	tokenBlock_: `
	// Visit a single <~tokenName> token.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	v.processor_.Process<~TokenName>(<variableName_>)
`,

	singularTokenBlock_: `
	// Visit a single <~tokenName> token.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.Process<~TokenName>(<variableName_>, 1, 1)
	}
`,

	optionalTokenBlock_: `
	// Visit an optional <~tokenName> token.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.Process<~TokenName>(<variableName_>)
	}
`,

	repeatedTokenBlock_: `
	// Visit each <~tokenName> token.
	var <~tokenName>Index uint
	var <variableName_> = <targetName_>.Get<~VariableName>().GetIterator()
	var <variableName>Size = uint(<variableName_>.GetSize())
	for <variableName_>.HasNext() {
		<~tokenName>Index++
		var <tokenName_> = <variableName_>.GetNext()
		v.processor_.Process<~TokenName>(
			<tokenName_>,
			<~tokenName>Index,
			<variableName>Size,
		)
	}
`,

	slotBlock_: `
	// Visit slot <Slot> between references.
	v.processor_.Process<~TargetName>Slot(<Slot>)
`,

	instanceStructure_: `

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}`,

	classStructure_: `

type visitorClass_ struct {
	// Declare the class constants.
}`,

	classReference_: `

func visitorReference() *visitorClass_ {
	return visitorReference_
}

var visitorReference_ = &visitorClass_{
	// Initialize the class constants.
}`,
}

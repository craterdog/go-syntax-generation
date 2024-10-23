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

func ScannerGenerator() ScannerGeneratorClassLike {
	return scannerGeneratorReference()
}

// Constructor Methods

func (c *scannerGeneratorClass_) Make() ScannerGeneratorLike {
	var instance = &scannerGenerator_{
		// Initialize the instance attributes.
		analyzer_: Analyzer().Make(),
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *scannerGenerator_) GetClass() ScannerGeneratorClassLike {
	return scannerGeneratorReference()
}

func (v *scannerGenerator_) GenerateScannerClass(
	module string,
	syntax ast.SyntaxLike,
) (
	implementation string,
) {
	v.analyzer_.AnalyzeSyntax(syntax)
	implementation = scannerGeneratorReference().classTemplate_
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

	var functionMethods = v.generateFunctionMethods()
	implementation = uti.ReplaceAll(
		implementation,
		"functionMethods",
		functionMethods,
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

func (v *scannerGenerator_) generateAccessFunction() (
	implementation string,
) {
	implementation = scannerGeneratorReference().accessFunction_
	return implementation
}

func (v *scannerGenerator_) generateConstructorMethods() (
	implementation string,
) {
	implementation = scannerGeneratorReference().constructorMethods_
	return implementation
}

func (v *scannerGenerator_) generateExpressions() (
	implementation string,
) {
	var expressions = v.analyzer_.GetExpressions().GetIterator()
	for expressions.HasNext() {
		var association = expressions.GetNext()
		var expressionName = association.GetKey()
		var regexp = association.GetValue()
		var expression = scannerGeneratorReference().expression_
		expression = uti.ReplaceAll(expression, "expressionName", expressionName)
		expression = uti.ReplaceAll(expression, "regexp", regexp)
		implementation += expression
	}
	return implementation
}

func (v *scannerGenerator_) generateFoundCases() (
	implementation string,
) {
	var tokenNames = v.analyzer_.GetTokenNames().GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		var foundCase = scannerGeneratorReference().foundCase_
		foundCase = uti.ReplaceAll(foundCase, "tokenName", tokenName)
		implementation += foundCase
	}
	return implementation
}

func (v *scannerGenerator_) generateFunctionMethods() (
	implementation string,
) {
	implementation = scannerGeneratorReference().functionMethods_
	return implementation
}

func (v *scannerGenerator_) generatePrimaryMethods() (
	implementation string,
) {
	implementation = scannerGeneratorReference().primaryMethods_
	return implementation
}

func (v *scannerGenerator_) generatePrivateMethods() (
	implementation string,
) {
	implementation = scannerGeneratorReference().privateMethods_
	var foundCases = v.generateFoundCases()
	implementation = uti.ReplaceAll(implementation, "foundCases", foundCases)
	return implementation
}

func (v *scannerGenerator_) generateTokenMatchers() (
	implementation string,
) {
	var tokenNames = v.analyzer_.GetTokenNames().GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		var tokenMatcher = scannerGeneratorReference().tokenMatcher_
		tokenMatcher = uti.ReplaceAll(tokenMatcher, "tokenName", tokenName)
		implementation += tokenMatcher
	}
	return implementation
}

func (v *scannerGenerator_) generateTokenIdentifiers() (
	implementation string,
) {
	implementation = `ErrorToken: "error",`
	var tokenNames = v.analyzer_.GetTokenNames().GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		var tokenIdentifier = scannerGeneratorReference().tokenIdentifier_
		tokenIdentifier = uti.ReplaceAll(tokenIdentifier, "tokenName", tokenName)
		implementation += tokenIdentifier
	}
	return implementation
}

func (v *scannerGenerator_) generateInstanceStructure() (
	implementation string,
) {
	implementation = scannerGeneratorReference().instanceStructure_
	return implementation
}

func (v *scannerGenerator_) generateClassStructure() (
	implementation string,
) {
	implementation = scannerGeneratorReference().classStructure_
	return implementation
}

func (v *scannerGenerator_) generateClassReference() (
	implementation string,
) {
	implementation = scannerGeneratorReference().classReference_
	var tokenIdentifiers = v.generateTokenIdentifiers()
	implementation = uti.ReplaceAll(
		implementation,
		"tokenIdentifiers",
		tokenIdentifiers,
	)
	var tokenMatchers = v.generateTokenMatchers()
	implementation = uti.ReplaceAll(
		implementation,
		"tokenMatchers",
		tokenMatchers,
	)
	var expressions = v.generateExpressions()
	implementation = uti.ReplaceAll(
		implementation,
		"expressions",
		expressions,
	)
	return implementation
}

func (v *scannerGenerator_) generateModuleImports(
	module string,
) (
	implementation string,
) {
	implementation = scannerGeneratorReference().moduleImports_
	implementation = uti.ReplaceAll(implementation, "module", module)
	return implementation
}

func (v *scannerGenerator_) generateNotice() string {
	var notice = v.analyzer_.GetNotice()
	return notice
}

func (v *scannerGenerator_) generatePackageDeclaration() (
	implementation string,
) {
	implementation = scannerGeneratorReference().packageDeclaration_
	return implementation
}

func (v *scannerGenerator_) generateProcessRule(
	ruleName string,
) (
	implementation string,
) {
	implementation = scannerGeneratorReference().processRule_
	if v.analyzer_.IsPlural(ruleName) {
		implementation = scannerGeneratorReference().processIndexedRule_
	}
	implementation = uti.ReplaceAll(implementation, "ruleName", ruleName)
	return implementation
}

func (v *scannerGenerator_) generateProcessRules() (
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

func (v *scannerGenerator_) generateProcessToken(
	tokenName string,
) (
	implementation string,
) {
	if tokenName == "delimiter" {
		return implementation
	}
	implementation = scannerGeneratorReference().processToken_
	if v.analyzer_.IsPlural(tokenName) {
		implementation = scannerGeneratorReference().processIndexedToken_
	}
	implementation = uti.ReplaceAll(implementation, "tokenName", tokenName)
	return implementation
}

func (v *scannerGenerator_) generateProcessTokens() (
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

type scannerGenerator_ struct {
	// Declare the instance attributes.
	analyzer_ AnalyzerLike
}

// Class Structure

type scannerGeneratorClass_ struct {
	// Declare the class constants.
	classTemplate_       string
	packageDeclaration_  string
	moduleImports_       string
	accessFunction_      string
	constructorMethods_  string
	functionMethods_     string
	methodicalMethods_   string
	processToken_        string
	processIndexedToken_ string
	processRule_         string
	processIndexedRule_  string
	primaryMethods_      string
	privateMethods_      string
	foundCase_           string
	instanceStructure_   string
	classStructure_      string
	classReference_      string
	tokenIdentifier_     string
	tokenMatcher_        string
	expression_          string
}

// Class Reference

func scannerGeneratorReference() *scannerGeneratorClass_ {
	return scannerGeneratorReference_
}

var scannerGeneratorReference_ = &scannerGeneratorClass_{
	// Initialize the class constants.
	classTemplate_: `<Notice><PackageDeclaration><ModuleImports>

// CLASS INTERFACE

// Access Function<AccessFunction>

// Constructor Methods<ConstructorMethods><FunctionMethods>

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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	reg "regexp"
	sts "strings"
	uni "unicode"
)`,

	accessFunction_: `

func Scanner() ScannerClassLike {
	return scannerReference()
}`,

	constructorMethods_: `

func (c *scannerClass_) Make(
	source string,
	tokens abs.QueueLike[TokenLike],
) ScannerLike {
	if uti.IsUndefined(source) {
		panic("The \"source\" attribute is required by this class.")
	}
	if uti.IsUndefined(tokens) {
		panic("The \"tokens\" attribute is required by this class.")
	}
	var instance = &scanner_{
		// Initialize the instance attributes.
		line_:     1,
		position_: 1,
		runes_:    []rune(source),
		tokens_:   tokens,
	}
	go instance.scanTokens() // Do scanning in the background...
	return instance
}`,

	functionMethods_: `

func (c *scannerClass_) FormatToken(token TokenLike) string {
	var result_ string
	var value = token.GetValue()
	value = fmt.Sprintf("%q", value)
	if len(value) > 40 {
		value = fmt.Sprintf("%.40q...", value)
	}
	result_ = fmt.Sprintf(
		"Token [type: %s, line: %d, position: %d]: %s",
		c.tokens_[token.GetType()],
		token.GetLine(),
		token.GetPosition(),
		value,
	)
	return result_
}

func (c *scannerClass_) FormatType(tokenType TokenType) string {
	var result_ = c.tokens_[tokenType]
	return result_
}

func (c *scannerClass_) MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var result_ bool
	var matcher = c.matchers_[tokenType]
	var match = matcher.FindString(tokenValue)
	result_ = uti.IsDefined(match)
	return result_
}`,

	methodicalMethods_: `<ProcessTokens><ProcessRules>`,

	processToken_: `

func (v *scanner_) Process<~TokenName>(
	<tokenName_> string,
) {
	v.validateToken(<tokenName_>, <~TokenName>Token)
}`,

	processIndexedToken_: `

func (v *scanner_) Process<~TokenName>(
	<tokenName_> string,
	index uint,
	size uint,
) {
	v.validateToken(<tokenName_>, <~TokenName>Token)
}`,

	processRule_: `

func (v *scanner_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add any validation checks.
}

func (v *scanner_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *scanner_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add any validation checks.
}`,

	processIndexedRule_: `

func (v *scanner_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *scanner_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *scanner_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}`,

	primaryMethods_: `

func (v *scanner_) GetClass() ScannerClassLike {
	return scannerReference()
}`,

	privateMethods_: `

func (v *scanner_) emitToken(tokenType TokenType) {
	var value = string(v.runes_[v.first_:v.next_])
	switch value {
	case "\x00":
		value = "<NULL>"
	case "\a":
		value = "<BELL>"
	case "\b":
		value = "<BKSP>"
	case "\t":
		value = "<HTAB>"
	case "\f":
		value = "<FMFD>"
	case "\r":
		value = "<CRTN>"
	case "\v":
		value = "<VTAB>"
	}
	var token = Token().Make(v.line_, v.position_, tokenType, value)
	//fmt.Println(Scanner().FormatToken(token)) // Uncomment when debugging.
	v.tokens_.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundError() {
	v.next_++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(tokenType TokenType) bool {
	// Attempt to match the specified token type.
	var text = string(v.runes_[v.next_:])
	var matcher = scannerReference().matchers_[tokenType]
	var match = matcher.FindString(text)
	if len(match) == 0 {
		return false
	}

	// Check for false delimiter matches.
	var token = []rune(match)
	var length = uint(len(token))
	var previous = token[length-1]
	if tokenType == DelimiterToken && uint(len(v.runes_)) > v.next_+length {
		var next = v.runes_[v.next_+length]
		if (uni.IsLetter(previous) || uni.IsNumber(previous)) &&
			(uni.IsLetter(next) || uni.IsNumber(next) || next == '_') {
			return false
		}
	}

	// Found the requested token type.
	v.next_ += length
	v.emitToken(tokenType)
	var count = uint(sts.Count(match, "\n"))
	if count > 0 {
		v.line_ += count
		v.position_ = v.indexOfLastEol(token)
	} else {
		v.position_ += v.next_ - v.first_
	}
	v.first_ = v.next_
	return true
}

func (v *scanner_) indexOfLastEol(runes []rune) uint {
	var length = uint(len(runes))
	for index := length; index > 0; index-- {
		if runes[index-1] == '\n' {
			return length - index + 1
		}
	}
	return 0
}

func (v *scanner_) scanTokens() {
loop:
	for v.next_ < uint(len(v.runes_)) {
		switch {
		// Find the next token type.<FoundCases>
		default:
			v.foundError()
			break loop
		}
	}
	v.tokens_.CloseQueue()
}`,

	foundCase_: `
		case v.foundToken(<~TokenName>Token):`,

	instanceStructure_: `

type scanner_ struct {
	// Declare the instance attributes.
	first_    uint // A zero based index of the first possible rune in the next token.
	next_     uint // A zero based index of the next possible rune in the next token.
	line_     uint // The line number in the source string of the next rune.
	position_ uint // The position in the current line of the next rune.
	runes_    []rune
	tokens_   abs.QueueLike[TokenLike]
}`,

	classStructure_: `

type scannerClass_ struct {
	// Declare the class constants.
	tokens_   map[TokenType]string
	matchers_ map[TokenType]*reg.Regexp
}`,

	classReference_: `

func scannerReference() *scannerClass_ {
	return scannerReference_
}

var scannerReference_ = &scannerClass_{
	// Initialize the class constants.
	tokens_: map[TokenType]string{
		// Define identifiers for each type of token.<TokenIdentifiers>
	},
	matchers_: map[TokenType]*reg.Regexp{
		// Define pattern matchers for each type of token.<TokenMatchers>
	},
}

// Private Constants

/*
NOTE:
These private constants define the regular expression sub-patterns that make up
the intrinsic types and token types.  Unfortunately there is no way to make them
private to the scanner class since they must be TRUE Go constants to be used in
this way.  We append an underscore to each name to lessen the chance of a name
collision with other private Go class constants in this package.
*/
const (
	// Define the regular expression patterns for each intrinsic type.
	any_     = "." // This does NOT include newline characters.
	control_ = "\\p{Cc}"
	digit_   = "\\p{Nd}"
	eol_     = "\\r?\\n"
	lower_   = "\\p{Ll}"
	upper_   = "\\p{Lu}"

	// Define the regular expression patterns for each token type."<Expressions>
)`,

	tokenIdentifier_: `
		<~TokenName>Token: "<~tokenName>",`,

	tokenMatcher_: `
		<~TokenName>Token: reg.MustCompile("^" + <~tokenName>_),`,

	expression_: `
	<~expressionName>_ = <Regexp>`,
}

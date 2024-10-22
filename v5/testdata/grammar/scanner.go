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

package grammar

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Scanner() ScannerClassLike {
	return scannerReference()
}

// Constructor Methods

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
		source_: source,
		tokens_: tokens,
	}
	return instance

}

// Function Methods

func (c *scannerClass_) FormatToken(
	token TokenLike,
) string {
	var result_ string
	// TBD - Add the function implementation.
	return result_
}

func (c *scannerClass_) FormatType(
	tokenType TokenType,
) string {
	var result_ string
	// TBD - Add the function implementation.
	return result_
}

func (c *scannerClass_) MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var result_ bool
	// TBD - Add the function implementation.
	return result_
}

// INSTANCE INTERFACE

// Primary Methods

func (v *scanner_) GetClass() ScannerClassLike {
	return scannerReference()
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type scanner_ struct {
	// Declare the instance attributes.
	source_ string
	tokens_ abs.QueueLike[TokenLike]
}

// Class Structure

type scannerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func scannerReference() *scannerClass_ {
	return scannerReference_
}

var scannerReference_ = &scannerClass_{
	// Initialize the class constants.
}

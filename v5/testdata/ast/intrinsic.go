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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Intrinsic() IntrinsicClassLike {
	return intrinsicReference()
}

// Constructor Methods

func (c *intrinsicClass_) Make(
	any_ any,
) IntrinsicLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &intrinsic_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *intrinsic_) GetClass() IntrinsicClassLike {
	return intrinsicReference()
}

// Attribute Methods

func (v *intrinsic_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type intrinsic_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type intrinsicClass_ struct {
	// Declare the class constants.
}

// Class Reference

func intrinsicReference() *intrinsicClass_ {
	return intrinsicReference_
}

var intrinsicReference_ = &intrinsicClass_{
	// Initialize the class constants.
}

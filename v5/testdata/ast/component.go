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

func Component() ComponentClassLike {
	return componentReference()
}

// Constructor Methods

func (c *componentClass_) Make(
	any_ any,
) ComponentLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &component_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *component_) GetClass() ComponentClassLike {
	return componentReference()
}

// Attribute Methods

func (v *component_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type component_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type componentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func componentReference() *componentClass_ {
	return componentReference_
}

var componentReference_ = &componentClass_{
	// Initialize the class constants.
}

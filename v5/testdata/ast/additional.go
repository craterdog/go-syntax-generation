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

func Additional() AdditionalClassLike {
	return additionalReference()
}

// Constructor Methods

func (c *additionalClass_) Make(
	component ComponentLike,
) AdditionalLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &additional_{
		// Initialize the instance attributes.
		component_: component,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *additional_) GetClass() AdditionalClassLike {
	return additionalReference()
}

// Attribute Methods

func (v *additional_) GetComponent() ComponentLike {
	return v.component_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type additional_ struct {
	// Declare the instance attributes.
	component_ ComponentLike
}

// Class Structure

type additionalClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalReference() *additionalClass_ {
	return additionalReference_
}

var additionalReference_ = &additionalClass_{
	// Initialize the class constants.
}

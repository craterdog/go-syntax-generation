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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func List() ListClassLike {
	return listReference()
}

// Constructor Methods

func (c *listClass_) Make(
	component ComponentLike,
	additionals abs.Sequential[AdditionalLike],
) ListLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionals) {
		panic("The \"additionals\" attribute is required by this class.")
	}
	var instance = &list_{
		// Initialize the instance attributes.
		component_:   component,
		additionals_: additionals,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *list_) GetClass() ListClassLike {
	return listReference()
}

// Attribute Methods

func (v *list_) GetComponent() ComponentLike {
	return v.component_
}

func (v *list_) GetAdditionals() abs.Sequential[AdditionalLike] {
	return v.additionals_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type list_ struct {
	// Declare the instance attributes.
	component_   ComponentLike
	additionals_ abs.Sequential[AdditionalLike]
}

// Class Structure

type listClass_ struct {
	// Declare the class constants.
}

// Class Reference

func listReference() *listClass_ {
	return listReference_
}

var listReference_ = &listClass_{
	// Initialize the class constants.
}

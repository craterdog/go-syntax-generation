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

func Document() DocumentClassLike {
	return documentReference()
}

// Constructor Methods

func (c *documentClass_) Make(
	components abs.Sequential[ComponentLike],
) DocumentLike {
	if uti.IsUndefined(components) {
		panic("The \"components\" attribute is required by this class.")
	}
	var instance = &document_{
		// Initialize the instance attributes.
		components_: components,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *document_) GetClass() DocumentClassLike {
	return documentReference()
}

// Attribute Methods

func (v *document_) GetComponents() abs.Sequential[ComponentLike] {
	return v.components_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type document_ struct {
	// Declare the instance attributes.
	components_ abs.Sequential[ComponentLike]
}

// Class Structure

type documentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func documentReference() *documentClass_ {
	return documentReference_
}

var documentReference_ = &documentClass_{
	// Initialize the class constants.
}

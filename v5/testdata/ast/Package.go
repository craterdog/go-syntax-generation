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

/*
Package "ast" provides the abstract syntax tree (AST) classes for this module.
Each AST class manages the attributes associated with the rule definition found
in the syntax grammar with the same rule name as the class.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-test-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki
*/
package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// Class Definitions

/*
AdditionalClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete additional-like class.
*/
type AdditionalClassLike interface {
	// Constructor Methods
	Make(
		component ComponentLike,
	) AdditionalLike
}

/*
ComponentClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) ComponentLike
}

/*
DocumentClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructor Methods
	Make(
		components abs.Sequential[ComponentLike],
	) DocumentLike
}

/*
IntrinsicClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete intrinsic-like class.
*/
type IntrinsicClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) IntrinsicLike
}

/*
ListClassLike is a class interface that defines the complete set
of class constructors, constants and functions that must be supported by
each concrete list-like class.
*/
type ListClassLike interface {
	// Constructor Methods
	Make(
		component ComponentLike,
		additionals abs.Sequential[AdditionalLike],
	) ListLike
}

// Instance Definitions

/*
AdditionalLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete additional-like class.
*/
type AdditionalLike interface {
	// Primary Methods
	GetClass() AdditionalClassLike

	// Attribute Methods
	GetComponent() ComponentLike
}

/*
ComponentLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Primary Methods
	GetClass() ComponentClassLike

	// Attribute Methods
	GetAny() any
}

/*
DocumentLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete document-like class.
*/
type DocumentLike interface {
	// Primary Methods
	GetClass() DocumentClassLike

	// Attribute Methods
	GetComponents() abs.Sequential[ComponentLike]
}

/*
IntrinsicLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete intrinsic-like class.
*/
type IntrinsicLike interface {
	// Primary Methods
	GetClass() IntrinsicClassLike

	// Attribute Methods
	GetAny() any
}

/*
ListLike is an instance interface that defines the complete set
of primary, attribute and aspect methods that must be supported by each
instance of a concrete list-like class.
*/
type ListLike interface {
	// Primary Methods
	GetClass() ListClassLike

	// Attribute Methods
	GetComponent() ComponentLike
	GetAdditionals() abs.Sequential[AdditionalLike]
}

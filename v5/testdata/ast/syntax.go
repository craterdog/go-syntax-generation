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

func Syntax() SyntaxClassLike {
	return syntaxReference()
}

// Constructor Methods

func (c *syntaxClass_) Make(
	notice NoticeLike,
	comment string,
	rules abs.Sequential[RuleLike],
	comment string,
	expressions abs.Sequential[ExpressionLike],
) SyntaxLike {
	if uti.IsUndefined(notice) {
		panic("The \"notice\" attribute is required by this class.")
	}
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	if uti.IsUndefined(rules) {
		panic("The \"rules\" attribute is required by this class.")
	}
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	if uti.IsUndefined(expressions) {
		panic("The \"expressions\" attribute is required by this class.")
	}
	var instance = &syntax_{
		// Initialize the instance attributes.
		notice_:      notice,
		comment_:     comment,
		rules_:       rules,
		comment_:     comment,
		expressions_: expressions,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *syntax_) GetClass() SyntaxClassLike {
	return syntaxReference()
}

// Attribute Methods

func (v *syntax_) GetNotice() NoticeLike {
	return v.notice_
}

func (v *syntax_) GetComment() string {
	return v.comment_
}

func (v *syntax_) GetRules() abs.Sequential[RuleLike] {
	return v.rules_
}

func (v *syntax_) GetComment() string {
	return v.comment_
}

func (v *syntax_) GetExpressions() abs.Sequential[ExpressionLike] {
	return v.expressions_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type syntax_ struct {
	// Declare the instance attributes.
	notice_      NoticeLike
	comment_     string
	rules_       abs.Sequential[RuleLike]
	expressions_ abs.Sequential[ExpressionLike]
}

// Class Structure

type syntaxClass_ struct {
	// Declare the class constants.
}

// Class Reference

func syntaxReference() *syntaxClass_ {
	return syntaxReference_
}

var syntaxReference_ = &syntaxClass_{
	// Initialize the class constants.
}

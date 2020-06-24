package aaaaaaaa

import (
	s "github.com/qeetell-aaaa/aaaaaaab" // A package for working with slice.
)


// ==== //

// Object type Sentence represents a language sentence.
type Sentence []byte


func Sentence_New () (*Sentence) {
	return &Sentence {}
}


// Method add () adds a new byte to the end of a sentence.
func (obj *Sentence) add (fig byte) {
	*obj = append (*obj, fig)
}

// Method removeLast () removes the last byte of a sentence.
func (obj *Sentence) removeLast () {
	*obj = s.RemoveLastByte (*obj)
}

// Method clear () deletes all the bytes in a sentence, making the sentence empty.
func (obj *Sentence) clear () {
	*obj = []byte {}
}

// Method Sentence () supplies the bytes in a sentence.
func (obj *Sentence) Sentence () ([]byte) {
	return *obj
}

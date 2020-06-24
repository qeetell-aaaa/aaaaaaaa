package aaaaaaaa

import (
	"fmt"
	"testing"
)

func TestSentence (t *testing.T) {
	sent := *Sentence_New ()
	fmt.Println ("a:", string (sent.Sentence ()))
	
	sent.add (byte ('h'))
	sent.add (byte ('e'))
	sent.add (byte ('l'))
	sent.add (byte ('l'))
	sent.add (byte ('o'))
	fmt.Println ("b:", string (sent.Sentence ()))

	sent.removeLast ()
	sent.removeLast ()
	fmt.Println ("c:", string (sent.Sentence ()))

	sent.clear ()
	fmt.Println ("d:", string (sent.Sentence ()))
}

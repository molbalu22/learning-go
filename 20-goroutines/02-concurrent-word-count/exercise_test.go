package wordcount

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{"Single text", []string{"the quick brown fox jumps over the lazy dog"},
			map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}},
		{"Multiple texts", []string{"hello world", "world hello", "hello hello"},
			map[string]int{"hello": 4, "world": 2}},
		{"Even more texts", []string{
			"the cat sat quietly on the windowsill",
			"the cat chased a mouse through the tall grass",
			"a mouse can squeeze through very small holes",
			"the small box was hidden behind the bookshelf",
			"she found an old book on the top shelf",
			"he opened the dusty book and began to read",
			"as he read a strange sound echoed in the hallway",
			"a loud sound startled the sleeping dog",
			"the dog barked at the passing stranger",
			"that stranger was wearing a red coat",
			"the coat was torn and stained with paint",
			"bright paint covered the walls of the nursery",
			"the cheerful walls made the room feel warm and inviting",
		},
			map[string]int{
				"the":        17,
				"a":          5,
				"was":        3,
				"and":        3,
				"cat":        2,
				"on":         2,
				"mouse":      2,
				"through":    2,
				"small":      2,
				"book":       2,
				"he":         2,
				"read":       2,
				"sound":      2,
				"dog":        2,
				"stranger":   2,
				"coat":       2,
				"paint":      2,
				"walls":      2,
				"sat":        1,
				"quietly":    1,
				"windowsill": 1,
				"chased":     1,
				"tall":       1,
				"grass":      1,
				"can":        1,
				"squeeze":    1,
				"very":       1,
				"holes":      1,
				"box":        1,
				"hidden":     1,
				"behind":     1,
				"bookshelf":  1,
				"she":        1,
				"found":      1,
				"an":         1,
				"old":        1,
				"top":        1,
				"shelf":      1,
				"opened":     1,
				"dusty":      1,
				"began":      1,
				"to":         1,
				"as":         1,
				"strange":    1,
				"echoed":     1,
				"in":         1,
				"hallway":    1,
				"loud":       1,
				"startled":   1,
				"sleeping":   1,
				"barked":     1,
				"at":         1,
				"passing":    1,
				"that":       1,
				"wearing":    1,
				"red":        1,
				"torn":       1,
				"stained":    1,
				"with":       1,
				"bright":     1,
				"covered":    1,
				"of":         1,
				"nursery":    1,
				"cheerful":   1,
				"made":       1,
				"room":       1,
				"feel":       1,
				"warm":       1,
				"inviting":   1,
			}},
		{"Empty input", []string{}, map[string]int{}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CountWords(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("CountWords(%v) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}

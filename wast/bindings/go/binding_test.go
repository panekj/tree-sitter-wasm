package tree_sitter_wast_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-wast"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_wast.Language())
	if language == nil {
		t.Errorf("Error loading Wast grammar")
	}
}

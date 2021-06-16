/*
#
# @Time : 2019/9/24 10:23
# @Author : Qian Lu
*/

package t208_tree_trie

import (
	"fmt"
	"testing"
)

func Test_TreeTrie(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")

	// 返回 true
	fmt.Println(trie.Search("apple"))
	// 返回 false
	fmt.Println(trie.Search("app"))

	trie.Insert("app")
	// 返回 true
	fmt.Println(trie.Search("app"))

	// 返回 true
	fmt.Println(trie.StartsWith("app"))
	fmt.Println(trie.StartsWith("appe"))

}

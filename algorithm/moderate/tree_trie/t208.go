/*
#
# @Time : 2019/9/24 18:32
# @Author : Qian Lu
*/

package tree_trie

/**
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
说明:

你可以假设所有的输入都是由小写字母 a-z 构成的。
保证所有输入均为非空字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type Trie struct {
	//Val  string
	Word string
	Next map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	trie := Trie{}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if this.Next == nil {
		this.Next = map[rune]*Trie{}
	}

	parent := this
	for i := 0; i < len(word); i++ {
		if parent.Next == nil {
			parent.Next = map[rune]*Trie{}
		}

		if parent.Next[rune(word[i])] == nil {
			trieChild := Trie{}
			//trieChild.Val = string(word[i])
			parent.Next[rune(word[i])] = &trieChild
			parent = &trieChild
		} else {
			parent = parent.Next[rune(word[i])]
		}

		if i == len(word)-1 {
			parent.Word = word
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for i := 0; i < len(word); i++ {
		if parent.Next[rune(word[i])] == nil {
			return false
		} else {
			parent = parent.Next[rune(word[i])]
		}

		if i == len(word)-1 && parent.Word == word {
			return true
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for i := 0; i < len(prefix); i++ {
		if parent.Next[rune(prefix[i])] == nil {
			return false
		} else {
			parent = parent.Next[rune(prefix[i])]
		}

		if i == len(prefix)-1 {
			return true
		}
	}
	return false
}

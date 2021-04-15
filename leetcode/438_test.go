package leetcode

import (
	"fmt"
	"testing"
)

// 给定一个字符串s和一个非空字符串p，找到s中所有是p的字母异位词的子串，返回这些子串的起始索引。
// 字符串只包含小写英文字母，并且字符串s和 p的长度都不超过 20100。
// 说明：
// 字母异位词指字母相同，但排列不同的字符串。
// 不考虑答案输出的顺序。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string

func Test_Problem438(t *testing.T) {

	flas := problem("aacblcfaancbasxcbaba", "cba")
	fmt.Printf("%v", flas)
}

func problem(s, p string) []int {
	m := [256]int{}
	count := 0
	flas := []int{}
	for k := range p {
		m[p[k]-'a']++
		count++
	}
	fmt.Printf("%v", m)
	for k := range s {
		i := 0
		fmt.Printf("%v\n", flas)

		for count > 0 {
			println(k, i, count, m[s[k+i]-'a'])

			if m[s[k+i]-'a'] > 0 {
				println(4444)
				count--
				i++
			} else {
				println(3333333)
				count += i
				break
			}
		}
		println(count, i)
		if count == i {
			flas = append(flas, k)
		}

	}
	return flas
}

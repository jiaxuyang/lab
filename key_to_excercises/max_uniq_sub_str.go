package main

/*
复盘：
1. 滑动窗口还是有点懵，虽然做出来了
2. 系统设计还是有点不知道面试官想要什么。最后的字典树+bitmap秀了一下，但是面试官想要的答案一直不对，很长的字符串，到底如何计数
*/
func maxSubStrLen(s string) int {
	countMap := make(map[byte]int)
	maxLen := 0
	left, right := 0, 0
	for right < len(s) {
		ch := s[right]
		countMap[ch]++
		if countMap[ch] == 1 {
			// no dup
			maxLen = max(maxLen, right+1-left)
		} else {
			for left < right && countMap[ch] > 1 {
				countMap[s[left]]--
				left++
			}
			// left==right or countMap[ch]->1
		}
		right++
	}
	return maxLen
}

func max(l, r int) int {
	if r > l {
		l = r
	}
	return l
}

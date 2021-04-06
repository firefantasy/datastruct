package kmp

func KmpFind(source string, target string) int {
	m, n := len(source), len(target)
	if n == 0 {
		return 0
	}

	lps := genLPS(target)
	i, j := 0, 0
	for i < m {
		if source[i] == target[j] {
			i++
			j++ 
			if j == n {
				return i - n
			}
		} else if j > 0 {
			j = lps[j-1]
		} else {
			i++
		}
	}
	
	return -1
}

// 用来生成KMP算法的longest prefix | suffix 数组
// abcdabc 生成的结果为[0, 0, 0, 0, 1, 2, 3]
func genLPS(target string) []int{
	n := len(target)
	i, k := 1, 0
	var lps = make([]int, n)
	// lps[0] 只有一个字符串，没有前缀和后缀
	for i < n {
		if target[k] == target[i] {
			k++
			lps[i] = k
			i++
		} else if k > 0 {
            // ...    k = 3 i
			//        |     |  
			// ... abyxxxabyy ...       
			//     k = lps[2] = 0

		    // ... k = 0    i
			//     |        |  
			// ... abyxxxabyy ...
			k = lps[k-1] + 1 // 回退后好前缀子串对齐了，直接+1移到下个位置
		} else {
			i++
		}
	}

	return lps
}
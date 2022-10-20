package k_th_symbol

func kthGrammar(n, k int) (ans int) {
	for k--; k > 0; k &= k - 1 {
		ans ^= 1
	}
	return
}

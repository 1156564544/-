package Question_by_day

type ATM struct {
	nums []int
}

func Constructor() ATM {
	return ATM{nums: make([]int, 5)}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < 5; i++ {
		this.nums[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	idx := 4
	for idx >= 0 {
		if this.nums[idx] == 0 || help(idx) > amount {
			idx -= 1
			continue
		}
		n := amount / help(idx)
		nn := min(n, this.nums[idx])
		amount = amount - help(idx)*nn
		this.nums[idx] -= nn
		res[idx] += nn
	}
	if amount == 0 {
		return res
	} else {
		for i := 0; i < 5; i++ {
			this.nums[i] += res[i]
		}
		return []int{-1}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func help(idx int) int {
	switch idx {
	case 0:
		return 20
	case 1:
		return 50
	case 2:
		return 100
	case 3:
		return 200
	case 4:
		return 500
	default:
		return -1
	}
}

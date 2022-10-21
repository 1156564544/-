package Question_by_day

type StockSpanner struct {
	prices []int
	kuadu  []int
}

func Constructor() StockSpanner {
	return StockSpanner{[]int{1000000}, []int{0}}
}

func (this *StockSpanner) Next(price int) int {
	if price < this.prices[len(this.prices)-1] {
		this.prices = append(this.prices, price)
		this.kuadu = append(this.kuadu, 1)
		return 1
	}
	this.prices = append(this.prices, price)
	i := len(this.prices) - this.kuadu[len(this.kuadu)-1] - 1
	for ; i > 0; i-- {
		if this.prices[i] > price {
			break
		}
	}
	this.kuadu = append(this.kuadu, len(this.prices)-i-1)
	return len(this.prices) - i - 1
}

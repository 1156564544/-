package Question_by_day

import "strconv"

func solveEquation(equation string) string {
	a, b := 0, 0
	flag := 1
	for i := 0; i < len(equation); i++ {
		if equation[i] == '=' || equation[i] == '+' || equation[i] == '-' {
			if i > 0 && equation[i-1] == 'x' {
				if equation[i] == '=' {
					flag = -1
				}
				continue
			}
			var j int
			for j = i - 1; j >= 0; j-- {
				if equation[j] < '0' || equation[j] > '9' {
					break
				}
			}
			v, _ := strconv.Atoi(equation[j+1 : i])
			if j == -1 || equation[j] == '=' || equation[j] == '+' {
				b -= v * flag
			} else {
				b += v * flag
			}
			if equation[i] == '=' {
				flag = -1
				continue
			}
		}
		if equation[i] == 'x' {
			var j int
			for j = i - 1; j >= 0; j-- {
				if equation[j] < '0' || equation[j] > '9' {
					break
				}
			}
			var v int
			if j == i-1 {
				v = 1
			} else {
				v, _ = strconv.Atoi(equation[j+1 : i])
			}
			if j == -1 || equation[j] == '=' || equation[j] == '+' {
				a += v * flag
			} else {
				a -= v * flag
			}
		}

	}
	i := len(equation) - 1
	if equation[i] != 'x' {
		var j int
		for j = i - 1; j >= 0; j-- {
			if equation[j] < '0' || equation[j] > '9' {
				break
			}
		}
		v, _ := strconv.Atoi(equation[j+1 : i+1])
		if j == -1 || equation[j] == '=' || equation[j] == '+' {
			b -= v * flag
		} else {
			b += v * flag
		}
	}
	if a == 0 {
		if b == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	} else {
		res := b / a
		return "x=" + strconv.Itoa(res)
	}
}

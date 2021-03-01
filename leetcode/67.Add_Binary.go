package leetcode

func addBinary(a string, b string) string {

	answer := ""
	i := len(a)-1
	j := len(b)-1

	carry := byte(0)
	for i>-1 || j>-1 {
		var sum byte
		if i>-1 {
			sum = sum + a[i]-'0'
			i = i-1
		}

		if j>-1 {
			sum = sum + b[j]-'0'
			j = j-1
		}

		sum = sum + carry
		ap := sum%2
		carry = sum/2
		answer = answer + string(ap+'0')
	}
	if carry == 1 {
		answer = answer + "1"
	}
	answer = reverse(answer)
	return answer
}

func reverse(s string) string {
	reverseStr := ""
	// for i:=0; i<len(s) i++ 이렇게하면 안됨.
	for _, value := range s {
		reverseStr = string(value) + reverseStr
	}
	return reverseStr
}

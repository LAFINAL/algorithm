package leetcode

func ReconstructQueue(people [][]int) [][] int {
	answer := make([][]int, 0, 0)
	length := len(people)



	for len(answer) != length { // 조건
		max := 0
		for i:=0; i<len(people); i++ {
			if people[i][0] > max { // height
				max = people[i][0]
			}
		}

		tempPeople := make([][]int, 0, 0)
		maxPeople := make([][]int, 0, 0)
		for i:=0; i<len(people); i++ {
			if people[i][0] == max {
				maxPeople = append(maxPeople, people[i])
			} else {
				tempPeople = append(tempPeople, people[i])
			}
		}

		for i:=0; i<len(maxPeople); i++ {
			for j:=i+1; j<len(maxPeople); j++ {
				if maxPeople[i][1] > maxPeople[j][1] {
					temp := maxPeople[i][1]
					maxPeople[i][1] = maxPeople[j][1]
					maxPeople[j][1] = temp
				}
			}
		}

		for i:=0; i<len(maxPeople); i++ {
			answer = append(answer, make([]int, 2, 2))
			if answer[maxPeople[i][1]][0] == 0 { // 자리가 비어있으면
				answer[maxPeople[i][1]] = maxPeople[i] // 거기 들어가고
			} else { // 자리 없으면 한 칸씩 뒤로 물리기.
				copy(answer[maxPeople[i][1]+1:], answer[maxPeople[i][1]:])
				answer[maxPeople[i][1]] = maxPeople[i]
			}
		}
		people = tempPeople
	}
	return answer
}


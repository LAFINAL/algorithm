package sorting

func SelectionSort(source []int) []int {

	// selection sort는 추가 공간 사용이 없다.
	// selection sort는 min 값을 찾아서 맨 앞에서 부터 쓰는 것이다. min의 위치도 알아야겠네? -> min의 위치'만' 알면된다.
	// for loop 안에서 minIndex 생성과 소멸이 반복되는건가?
	// 일단 for문 자체도 중괄호, 함수고 for문 돌때 반복 한 번마다 전부 초기화 된다.
	// 그리고 반복문에 쓰이는 변수는 무조건 그 안쪽에서 쓰길 바란다. 성능도 그렇게 좋지 않다고 나옴.
	// for문 scope안에서만 쓰는 변수라는게 더 중요함.

	for i:=0; i<len(source)-1; i++ {
		minIndex := i
		for j:=i+1; j<len(source); j++ {
			if source[j] < source[minIndex] {
				minIndex = j
			}
		}
		source[i], source[minIndex] = source[minIndex], source[i]
	}

	return source
}

/*
	패키지 외부 공개

	패키지내 타입, 전역변수, 상수, 함수, 메서드 등을 외부로 공개해서 사용할 수 있음
	대문자로 시작하면 외부 공개
	소문자로 시작하면 비공개


	type Student struct { // 첫글자가 대문자이므로 외부로 공개됨
		Name string
		Age int
		score int
	}

	var Ratio int // 외부로 공개
	var ttt int // 외부로 공개되지 않음

	const PI = 3.14 // 외부로 공개

*/
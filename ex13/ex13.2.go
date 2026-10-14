/*
var house House // 모든 필드값 기본값
var house House = House { "서울시 강동구", 28, 9.80, "아파트"} // 이렇게 넣으며 초기화 

var house House = House {
	"서울시 강동구", 
	28, 
	9.80, 
	"아파트", // go lang은 마지막에 쉼표 넣어줘야 된다. --> 규칙임
}

var house House = House{ Size: 28, Type: "아파트"}  // 특정 값만 초기화 나머진 기본값

var house House = House{ 
	Size: 28, 
	Type: "아파트", // 마지막 쉼표 잊지 말기 --> Go lang 문법
}  

*/
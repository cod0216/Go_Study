/*
객체로 진화

객체(Object)란 데이터(state)와 기능(Function)을 묶은 것

절차 중심에서 관계 중심(OOP) <-- 클래스 다이어그램이 중요하다.

클래스간 상호 작용
	func (s *Student) EnrollClass(c *Subject) {
		...
	}

	func(s *Student) SendReport(p *Professor, r *Report) { // func이랑 똑같은데 개념 상 차이가 있을 뿐
		...
	}

*/
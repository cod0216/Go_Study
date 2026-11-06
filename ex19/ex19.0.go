/*
메서드는 타입에 속한 함수

메서드는 타입에 종속적

Type 안에 메서드가 포함되어 있다.

	func (r Rabbit) // 리시버 info() // 메서드 명 int{
		return r.width * r.height
	}

리시버는 모든 패키지 지역 타입이 가능하다
 구조체, 별칭 타입 등
*/

/*
	다른 언어는 클래스 안에 메서드가 정의된다.

class Info {
	public static void main(){
		// 내용
	}
}

type Rabbit struct {

}

func (r Rabbit) info() {
	리시버에 info()메서드가 포함되어 있게 된다.
}

info는 Rabbit에 포함된다.

리시버는 모든 지역 타입이 가능하다.

type MyInt int

func (m MyInt) fff() {

}

	메서드는 함수다 (근데 리시버가 있는 함수다.)

*/
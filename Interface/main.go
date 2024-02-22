package main

import "fmt"

// 4. 그러나 인터페이스를 사용하게 되면 두 객체는 서로의 내부 구조를 몰라도 된다.
// 인터페이스를 통해 서로의 메소드를 호출할 수 있기 때문이다.
// 인터페이스는 메소드만 정의한다.
type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type SpoonOfJam interface {
	String() string
}

// 5. 4번을 진행했을 때는 Jam인터페이스만 공유하면 된다.
func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()

}

// 1. 브레드라는 하나의 객체와
type Bread struct {
	val string
}

// 2. 딸기잼이라는 하나의 객체가 있다.
type StrawberryJam struct {
}

// 6. 오렌지잼이라는 구조체를 만들어도 Jam인터페이스를 구현하기만 하면 된다.
type OrangeJam struct {
}

// 7. 애플잼을 만들어도 Jam인터페이스를 구현하기만 하면 된다.
type AppleJam struct {
}

type SpoonOfStrawberryJam struct {
}

type SpoonOfOrangeJam struct {
}

type SpoonOfAppleJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ orange"
}

func (s *SpoonOfAppleJam) String() string {
	return "+ apple"
}

func (b *Bread) String() string {
	return "bread " + b.val
}

// 3. 별개의 두 객체가 있다. 하지만 두 객체는 서로의 메소드를 호출한다.
// 이는 두 객체가 서로의 내부 구조를 알고 있어야 하기 때문에 결합도가 높아진다.
func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (o *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (a *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

// func (b *Bread) PutJam(jam *StrawberryJam) {
// 	spoon := jam.GetOneSpoon()
// 	b.val += spoon.String()

// }

func main() {
	// Call the function
	bread := &Bread{}
	sjam := &StrawberryJam{}
	bread.PutJam(sjam)
	ojam := &OrangeJam{}
	bread.PutJam(ojam)
	ajam := &AppleJam{}
	bread.PutJam(ajam)

	fmt.Println(bread.String())
}

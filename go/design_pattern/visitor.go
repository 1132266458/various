package main

import "fmt"

//VipVisitor 会员，打折购买
type VipVisitor struct {
	discount float32
}

func (t VipVisitor) BuyApple(a Apple) {
	fmt.Println("苹果原价:", a.GetPrice(), "会员价格:", t.discount*float32(a.GetPrice()))
}

func (t VipVisitor) BuyBanana(b Banana) {
	fmt.Println("香蕉原价:", b.GetPrice(), "会员价格购:", t.discount*float32(b.GetPrice()))
}

func (t VipVisitor) BuyStawberry(c Strawberry) {
	fmt.Println("草莓原价:", c.GetPrice(), "会员价格", t.discount*float32(c.GetPrice()))
}

//NormalVisitor 普通用户，原价购买
type NormalVisitor struct {
}

func (t NormalVisitor) BuyApple(a Apple) {
	fmt.Println("苹果原价:", a.GetPrice(), "零售价格:", a.GetPrice())
}

func (t NormalVisitor) BuyBanana(b Banana) {
	fmt.Println("香蕉原价:", b.GetPrice(), "零售价格:", b.GetPrice())
}

func (t NormalVisitor) BuyStawberry(c Strawberry) {
	fmt.Println("草莓原价:", c.GetPrice(), "零售价格:", c.GetPrice())
}


type Visitor interface {
	BuyApple(a Apple)
	BuyBanana(b Banana)
	BuyStawberry(c Strawberry)
}
type Fruit interface {
	Accept(v Visitor)
	GetPrice() int
}
type Apple struct {
	price int
}

func (t Apple) Accept(v Visitor) {
	v.BuyApple(t)
}
func (t Apple) GetPrice() int {
	return t.price
}

//Banana 香蕉类
type Banana struct {
	price int
}

func (t Banana) Accept(v Visitor) {
	v.BuyBanana(t)
}
func (t Banana) GetPrice() int {
	return t.price
}

//Strawberry 草莓类
type Strawberry struct {
	price int
}

func (t Strawberry) Accept(v Visitor) {
	v.BuyStawberry(t)
}
func (t Strawberry) GetPrice() int {
	return t.price
}

func main() {
	f := make([]Fruit, 3)
	f[0] = Apple{}
	f[1] = Banana{}
	f[2] = Strawberry{}
	fmt.Println("-----------访问者模式--------------")
	//for _, fruit := range f {
	//	fruit.Accept(v)
	//}
}
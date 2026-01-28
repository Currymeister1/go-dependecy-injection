package main

import (
	"fmt"
)

type PaymentMethoder interface {
	pay()
	printDetails()
}

type ApiInfo struct {
	apiUrl string
	apiKey string
}

type MasterCard struct {
	ApiInfo
	name       string
	bankNumber string
	cvc        int
}

func (m MasterCard) pay() {
	//Simulating

	fmt.Printf("Calling the endpoint: %s\n", m.apiUrl)

	fmt.Println("Api Key accepted!")

	fmt.Printf("Dear %s Please complete the payment via your Bank App for the account %s\n", m.name, m.bankNumber)

	fmt.Println("Thank you for using Mastercard!")
}

func (m MasterCard) printDetails() {
	fmt.Printf("Name: %s Banknumber: %s CVC: %d\n", m.name, m.bankNumber, m.cvc)
}

type Paypal struct {
	ApiInfo
	email    string
	password string
}

func (p Paypal) pay() {
	//Simulating

	fmt.Printf("Calling the endpoint: %s \n", p.apiKey)

	fmt.Println("Api Key accepted!")

	fmt.Printf("The money will be deducted from your paypal account with email %s\n", p.email)

	fmt.Println("Thank you for using Paypal!")
}

func (p Paypal) printDetails() {
	fmt.Printf("Email: %s Password: %s\n", p.email, p.password)
}

type Checkout struct {
	totalAmount   int
	cartSize      int
	paymentMethod PaymentMethoder
}

func (c *Checkout) setPaymentMethod(paymentMethod PaymentMethoder) {
	c.paymentMethod = paymentMethod
}

func (c Checkout) toString() string {
	return fmt.Sprintf(
		"Total amount: %d with cart size of %d with payment",
		c.totalAmount,
		c.cartSize)
}

func main() {
	checkout := Checkout{totalAmount: 15, cartSize: 4}

	p := Paypal{
		ApiInfo: ApiInfo{
			apiUrl: "dummypal.com",
			apiKey: "dummypalkey",
		},
		email:    "dummy@email.com",
		password: "dummypassword",
	}

	m := MasterCard{
		ApiInfo: ApiInfo{
			apiUrl: "dummyMC.com",
			apiKey: "dummyMCKey",
		},
		name:       "James",
		bankNumber: "12454321243",
		cvc:        123,
	}

	paymentMethods := [...]PaymentMethoder{p, m}

	fmt.Println(checkout.toString())
	fmt.Println()

	for _, pm := range paymentMethods {
		checkout.setPaymentMethod(pm)
		checkout.paymentMethod.pay()
		fmt.Println()
	}
}

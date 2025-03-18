package main

import "fmt"

const FULLTIME_SALARY = 15000  //30 days
const CONTRACTOR_SALARY = 3000 //30 days
const FREELANCER_SALARY = 2000 //20 hrs

type Employee interface {
	calcSalary(duration float32) float32
}

type Fulltime struct {
	// day wise
	sal float32
}

type Contractor struct {
	// day wise
	sal float32
}

type Freelancer struct {
	// hourly basis
	sal float32
}

func (emp Fulltime) calcSalary(durationDays float32) float32 {
	return durationDays * FULLTIME_SALARY / 30
}

func (emp Contractor) calcSalary(durationDays float32) float32 {
	return durationDays * CONTRACTOR_SALARY / 30
}

func (emp Freelancer) calcSalary(durationHours float32) float32 {
	return durationHours * FREELANCER_SALARY / 20
}

func measure(e Employee, d float32) float32 {
	return e.calcSalary(d)
}

func main() {
	e1 := Fulltime{}
	e2 := Contractor{}
	e3 := Freelancer{}

	e1.sal = measure(e1, 10)
	e2.sal = measure(e2, 30)
	e3.sal = measure(e3, 20)
	fmt.Print(e1, e2, e3)
}

package employee

import (
	"fmt"
	"strings"
)

type Employee interface {
	GetDetails() string
}
type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}
type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (fte FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Salary: %d Tenge", fte.ID, fte.Name, fte.Salary)
}
func (pte PartTimeEmployee) GetDetails() string {
	totalEarnings := pte.HourlyRate * uint64(pte.HoursWorked)
	return fmt.Sprintf("ID: %d, Name: %s, Hourly Rate: %d Tenge, Hours Worked: %.2f, Total Earnings: %d Tenge", pte.ID, pte.Name, pte.HourlyRate, pte.HoursWorked, totalEarnings)
}

type Company struct {
	Employees map[string]Employee
}

func (c *Company) AddEmployee(emp Employee) {

	empID := fmt.Sprintf("%s-%d", strings.ToLower(emp.GetDetails()), len(c.Employees)+1)
	c.Employees[empID] = emp
}
func (c *Company) ListEmployees() {
	if len(c.Employees) == 0 {
		fmt.Println("No employees found.")
		return
	}
	for id, emp := range c.Employees {
		fmt.Printf("Employee ID: %s\n%s\n\n", id, emp.GetDetails())
	}
}

func main() {

	company := Company{
		Employees: make(map[string]Employee),
	}

	fte := FullTimeEmployee{
		ID:     1,
		Name:   "Alice",
		Salary: 150000,
	}

	pte := PartTimeEmployee{
		ID:          2,
		Name:        "Bob",
		HourlyRate:  5000,
		HoursWorked: 20.5,
	}

	company.AddEmployee(fte)
	company.AddEmployee(pte)

	company.ListEmployees()
}

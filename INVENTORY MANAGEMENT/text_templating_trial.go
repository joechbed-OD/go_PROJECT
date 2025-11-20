package main

import (
	"fmt"
	// "text/template"
	"time"
)

type Employee struct {
	Name              string
	Department        string
	Salary            int
	yearsOfExperience int
	Performance       float64
	age               int
}

type Report struct {
	companyName string
	Date        string
	Employees   []Employee
}

func yearsUntilRetirement(e *Employee) {
	if e.age < 65 {
		retirement := 65 - e.age
		fmt.Printf("%s has %d years to retire", e.Name, retirement)
	}
}

// func formatSalary(e *Employee) string {
// 	tmpl := template.Must(template.New("salary").Parse("{{.e.Name}} earns {{.e.Salary}}"))
// 	return tmpl
// }

func isTopPerformer(e *Employee) bool {
	if e.Performance >= 4.5 {
		fmt.Printf("%s is a top Performer", e.Name)
		return true
	}
	return false
}

func isBonusEligible(e *Employee) bool {
	if e.Performance >= 4.0 && e.yearsOfExperience >= 3 {
		fmt.Printf("%s is Bonus Eligible", e.Name)
		return true
	}
	return false
}

func main() {

	employees := []Employee{
		{"Alice Johnson", "Engineering", 95000, 5, 4.8, 32},
		{"Bob Smith", "Engineering", 75000, 2, 4.2, 28},
		{"Carol White", "Sales", 68000, 4, 4.6, 35},
		{"David Brown", "Sales", 82000, 6, 3.9, 41},
		{"Eve Davis", "HR", 70000, 3, 4.1, 29},
	}

	report := Report{
		companyName: "ACME Corp",
		Date:        time.Now().Format("2006-01-02"),
		Employees:   employees,
	}

	fmt.Println(report)
	fmt.Println(isBonusEligible(&Employee{}))
	fmt.Println(isTopPerformer(&Employee{}))
	// fmt.Println(yearsUntilRetirement(&Employee{e}))
	// fmt.Println()

}

package types

import "fmt"

type Employee struct {
	UserID       int    `json:"userId"`
	Age          int    `json:"age"`
	Name         string `json:"name"`
	DepartmentID int    `json:"departmentId"`
}

func (e Employee) String() string {
	str := "User ID: %d; Age: %d; Name: %s; Department ID: %d; "
	return fmt.Sprintf(str, e.UserID, e.Age, e.Name, e.DepartmentID)
}

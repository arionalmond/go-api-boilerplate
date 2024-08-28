package datastore

// Employee ...
type Employee struct {
	id int
}

// GetEmployeeByID ...
func (m *MySQLDS) GetEmployeeByID(empID int) (emp Employee, err error) {
	// sqlx query
	return Employee{}, nil
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's go!!")

	attendanceStatusTable := CreateSpecifiedAttendanceStatusTable()
	mostSuitableDate := attendanceStatusTable.Decide()

	fmt.Printf("mostSuitableDate is %+v\n", mostSuitableDate)
}

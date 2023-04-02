package academy
import "math" 

type Student struct {
	Name      string
	Grades    []int
	Project   int
	Attendace []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	if len(grades) == 0 {
		return 0
	}
	var sum float64
	for _,v := range grades {
		sum += float64(v)
	}
	avg := sum / float64(len(grades))
	return int(math.Round(avg))

}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from  0 to 1,
// with 2 digits of precision.
func AttendancePercentage(attendance []bool) float64 {
	if len(attendance) == 0 {
		return 0
	}
	amount := 0
	for x := 0; x < len(attendance); x++ {
		if attendance[x] == true {
			amount++;
		}
	}
	var percentage float64 = float64(amount) / float64(len(attendance))
	return math.Round(percentage*1000)/1000
}


// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	if s.Project == 0 || len(s.Grades) == 0 || len(s.Attendace) == 0 {
		return 0
	}
	attendance := AttendancePercentage(s.Attendace)
	avg := AverageGrade(s.Grades)
	grade := int(math.Round(float64(avg + s.Project)/2))

	if attendance < 0.6 || avg == 1 || s.Project == 1 {
		return 1
	}
	
	if attendance < 0.8 {
		grade--
	}

	return grade
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	m := make(map[string]uint8)
	for _, student := range students {
		m[student.Name] = uint8(FinalGrade(student));
	}
	return m
}


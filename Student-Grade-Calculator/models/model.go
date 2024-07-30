package models

type Subject struct {
	Name string;
	Score float32;
	Grade string;
	Grade_value float32;
}

type Student struct {
	Name string;
	Number_of_subjects int;
	Subjets []Subject;
}


func calculate_grade(score float32) string {
	if score > 90 {
		return "A+"
	}else if score > 85 {
		return "A"
	}else if score > 80 {
		return "A-"
	}else if score > 75 {
		return "B+"
	}else if score > 70 {
		return "B";
	}else if score > 65 {
		return "B-"
	}else if score > 60 {
		return "C+"
	}else if score > 50 {
		return "C"
	}else if score > 45 {
		return "C-"
	}else if score > 40 {
		return "D"
	}else {
		return "F"
	}
}

func calculate_grade_value(score float32) float32{
	if score > 85{
		return 4.0
	}else if score > 80 {
		return 3.75
	}else if score > 75 {
		return 3.3
	}else if score > 70 {
		return 3.0
	}else if score > 65 {
		return 2.7
	}else if score > 60 {
		return 2.3
	}else if score > 50 {
		return 2.0
	}else if score > 45 {
		return 1.7
	}else if score > 40 {
		return 1.0
	}else {
		return 0
	}
}

func (subject *Subject) Init_subject(name string , score float32) {
	subject.Name = name;
	subject.Score = score;
	subject.Grade = calculate_grade(score);
	subject.Grade_value = calculate_grade_value(score);
}

func (student *Student) Init_student(name string , subjectNo int , subjects []Subject) {
	student.Name = name;
	student.Number_of_subjects = subjectNo;
	student.Subjets = append(student.Subjets, subjects...)
}

func (student *Student) Ave_calculate_grade() float32 {
	var grade float32;
	for i := 0; i < student.Number_of_subjects; i++ {
		grade += student.Subjets[i].Grade_value;
	}
	return (grade/float32(student.Number_of_subjects));
} 


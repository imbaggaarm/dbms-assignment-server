package models

type Course struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Commitment string `json:"commitment"`
	Description string `json:"description"`
	SpecializationID uint `json:"specialization_id"`
	MinGrade float32 `json:"min_grade"`
	CoursePrice float32 `json:"course_price"`
	Active int `json:"active"`
}

type OverallCourse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	InstitutionName string `json:"institution_name" gorm:"column:ins_name"`
}

func (Course) TableName() string {
	return "course"
}

func GetCourse(id uint) *Course {
	course := Course{ID: id}
	err := GetDB().First(&course).Error
	if err != nil {
		return nil
	}
	return &course
}

func GetAllCourses(offset uint) []*OverallCourse {
	courses := make([]*OverallCourse, 0)
	err := db.Raw("call get_all_courses(?)", offset).Scan(&courses).Error
	if err != nil {
		return nil
	}
	return courses
}

func GetUserCourses(userID uint) []*Course {
	courses := make([]*Course, 0)
	err := db.Raw("call student_list_course(?)", userID).Scan(&courses).Error
	if err != nil {
		return nil
	}
	return courses
}
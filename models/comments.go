package models

import u "src/utils"

type Comment struct {
	ID          uint   `json:"id"`
	CourseID    uint   `json:"course_id"`
	StudentID   uint   `json:"student_id"`
	Content     string `json:"content"`
	CommentTime string `json:"cmmt_time" gorm:"column:cmmt_time"`
}

type ResponseComment struct {
	Comment
	StdLastName  string `json:"std_last_name"`
	StdFirstName string `json:"std_first_name"`
	StdImageUrl  string `json:"std_image_url"`
}

func (Comment) TableName() string {
	return "comment"
}

func GetComments(courseID uint, offSet uint) []*ResponseComment {
	comments := make([]*ResponseComment, 0)
	err := GetDB().Raw("call get_comments(?,?)", courseID, offSet).Scan(&comments).Error

	if err != nil {
		return nil
	}
	return comments
}

func (comment *Comment) Create() map[string]interface{} {
	GetDB().Create(comment)
	if comment.ID <= 0 {
		return u.Message(false, "Failed to create comment, connection error.")
	}

	resp := u.Message(true, "")
	resp["data"] = comment
	return resp
}

func (comment *Comment) Update(content string) map[string]interface{} {
	GetDB().First(&comment)
	comment.Content = content
	GetDB().Save(&comment)
	resp := u.Message(true, "")
	resp["data"] = comment
	return resp
}

func DeleteComment(id uint) map[string]interface{} {
	comment := Comment{ID: id}
	GetDB().Delete(&comment)
	response := u.Message(true, "Comment deleted")
	return response
}

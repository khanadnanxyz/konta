package model

// Question represents a question model
type Question2 struct {
	ID         uint64    `gorm:"AUTO_INCREMENT;primary_key"`
	QText      string    `gorm:"type:varchar(255);not null"`
}

// TableName set a custom table name for Question model
func (*Question2) TableName() string {
	return "poll.question2"
}

func (q *Question2) Validate() error {
	if q.QText == ""{
		ve := ValidationError{"QText", "is required"}
		return &ve
	}
	return nil
}

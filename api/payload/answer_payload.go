package payload

import (
	"github.com/khanadnanxyz/konta/errors"
)

type AnswerPayload struct {
	UserID uint64 `json:"user_id"`
	QuestionID uint64 `json:"question_id"`
	OptionID uint64 `json:"option_id"`
}

func (ans *AnswerPayload) Validate() error {
	if ans.QuestionID == 0.0 {
		return &errors.ValidationError{"Question Id", "can't be 0"}
	}
	return nil
}

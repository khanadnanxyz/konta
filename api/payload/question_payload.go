package payload

import (
	"github.com/khanadnanxyz/konta/errors"
)

type QuestionPayload struct {
	QText string `json:"q_text"`
	CategoryId uint64 `json:"category_id"`
	Options []string `json:"options"`
}

func (qp *QuestionPayload) Validate() error {

	if qp.QText == "" {
		return &errors.ValidationError{"q_text", "is required"}
	}
	//if (!govalidator.IsNumeric(string(qp.CategoryId))) {
	//	return &errors.ValidationError{"Number", "is required"}
	//}
	if len(qp.Options) < 2 {
		return &errors.ValidationError{"options", "less than 2"}
	}
	return nil
}
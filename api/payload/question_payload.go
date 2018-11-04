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
	if len(qp.Options) < 2 {
		return &errors.ValidationError{"options", "very few options"}
	}
	optionsOk := true
	for _, option := range qp.Options {
		if option == "" {
			optionsOk = false
			break
		}
	}
	if !optionsOk {
		return &errors.ValidationError{"options", "Option must not be empty"}
	}
	return nil
}
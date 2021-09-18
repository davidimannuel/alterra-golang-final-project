package label

import "keep-remind-app/businesses/label"

type AddLabelRequest struct {
	Name string `json:"name"`
}

type EditLabelRequest struct {
	Name string `json:"name"`
}

func (req *AddLabelRequest) toDomain() *label.LabelDomain {
	return &label.LabelDomain{
		Name: req.Name,
	}
}

func (req *EditLabelRequest) toDomain() *label.LabelDomain {
	return &label.LabelDomain{
		Name: req.Name,
	}
}

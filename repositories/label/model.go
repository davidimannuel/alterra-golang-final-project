package label

import (
	"keep-remind-app/businesses/label"

	"gorm.io/gorm"
)

var tableName string = "labels"

type LabelModel struct {
	gorm.Model
	UserID uint
	Name   string
}

func (model *LabelModel) TableName() string {
	return tableName
}

func fromDomain(domain *label.LabelDomain) *LabelModel {
	return &LabelModel{
		Model: gorm.Model{
			ID:        uint(domain.ID),
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		UserID: uint(domain.UserID),
		Name:   domain.Name,
	}
}

func (model *LabelModel) toDomain() label.LabelDomain {
	return label.LabelDomain{
		ID:        int(model.ID),
		UserID:    int(model.UserID),
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func toDomains(models []LabelModel) (domains []label.LabelDomain) {
	for i := range models {
		domains = append(domains, models[i].toDomain())
	}
	return domains
}

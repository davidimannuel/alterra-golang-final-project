package businesses

import "gorm.io/gorm"

// context usecase struct for business uc
// used for passing values ex: config for cross businness process
type ContextUC struct {
	AppHost string
	DB      *gorm.DB
}

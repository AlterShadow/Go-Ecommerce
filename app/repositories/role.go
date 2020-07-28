package repositories

import (
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"

	"goshop/app/models"
	"goshop/app/schema"
	"goshop/dbs"
)

type RoleRepository interface {
	CreateRole(req *schema.RoleBodyParam) (*models.Role, error)
}

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepository() RoleRepository {
	return &roleRepo{db: dbs.Database}
}

func (r *roleRepo) CreateRole(req *schema.RoleBodyParam) (*models.Role, error) {
	var role models.Role
	copier.Copy(&role, &req)

	if err := r.db.Create(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

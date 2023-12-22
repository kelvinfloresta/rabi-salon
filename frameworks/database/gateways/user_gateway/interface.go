package user_gateway

import (
	"rabi-salon/frameworks/database"
	"rabi-salon/usecases/auth_case/role"
)

type UserGateway interface {
	Create(input CreateInput) (string, error)
	GetByID(id string) (*GetByIDOutput, error)
	Patch(filter PatchFilter, values PatchValues) (bool, error)
	Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error)
	Delete(id string) (bool, error)
}

type CreateInput struct {
	State          string
	ZIP            string
	Phone          string
	City           string
	Photo          string
	TaxID          string
	SocialID       string
	Street         string
	Complement     string
	EmergencyPhone string
	Name           string
	Email          string
	Neighborhood   string
	Role           role.Role
}

type GetByIDOutput struct {
	Phone          string
	City           string
	State          string
	ZIP            string
	Name           string
	Email          string
	Photo          string
	TaxID          string
	SocialID       string
	Street         string
	Complement     string
	EmergencyPhone string
}

type PatchFilter struct {
	ID string
}

type PatchValues struct {
	ZIP            string
	Phone          string
	City           string
	State          string
	TaxID          string
	SocialID       string
	Street         string
	Complement     string
	EmergencyPhone string
	Name           string
	Email          string
	Photo          string
}

type PaginateFilter struct {
	State *string
	City  *string
	Name  *string
}

type PaginateData struct {
	ID    string
	Photo string
	Name  string
	State string
	City  string
}

type PaginateOutput struct {
	Data     []PaginateData
	MaxPages int
}

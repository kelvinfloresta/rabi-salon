package user_case

import (
	"rabi-salon/app_context"
	g "rabi-salon/frameworks/database/gateways/user_gateway"
)

type PatchFilter struct {
	ID             *string
	Phone          *string
	City           *string
	State          *string
	ZIP            *string
	SocialID       *string
	Street         *string
	Complement     *string
	EmergencyPhone *string
	Name           *string
	Email          *string
	Photo          *string
	TaxID          *string
}

type PatchValues struct {
	Phone          string
	City           string
	State          string
	ZIP            string
	SocialID       string
	Street         string
	Complement     string
	EmergencyPhone string
	Name           string
	Email          string
	Photo          string
	TaxID          string
}

func (c UserCase) Patch(ctx *app_context.AppContext, filter PatchFilter, values PatchValues) (bool, error) {
	return c.gateway.Patch(
		g.PatchFilter{
			ID:    filter.ID,
			Phone: filter.Phone, City: filter.City, State: filter.State, ZIP: filter.ZIP, SocialID: filter.SocialID, Street: filter.Street, Complement: filter.Complement, EmergencyPhone: filter.EmergencyPhone, Name: filter.Name, Email: filter.Email, Photo: filter.Photo, TaxID: filter.TaxID,
		}, g.PatchValues{
			Phone: values.Phone, City: values.City, State: values.State, ZIP: values.ZIP, SocialID: values.SocialID, Street: values.Street, Complement: values.Complement, EmergencyPhone: values.EmergencyPhone, Name: values.Name, Email: values.Email, Photo: values.Photo, TaxID: values.TaxID,
		})
}

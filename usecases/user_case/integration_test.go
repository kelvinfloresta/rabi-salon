package user_case_test

import (
	"net/http"
	"rabi-salon/fixtures"
	"rabi-salon/frameworks/database/gateways/user_gateway"
	"rabi-salon/usecases/user_case"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Integration_should_create(t *testing.T) {
	fixtures.CleanDatabase()

	Body := user_case.CreateInput{
		Name:           "Name",
		TaxID:          "TaxID",
		City:           "City",
		State:          "State",
		Phone:          "Phone",
		ZIP:            "ZIP",
		SocialID:       "SocialID",
		Email:          "Email",
		EmergencyPhone: "EmergencyPhone",
		Neighborhood:   "Neighborhood",
		Street:         "Street",
		Complement:     "Complement",
		Photo:          "Photo",
	}

	id := ""
	statusCode := fixtures.Post(t, fixtures.PostInput{
		Body:     Body,
		URI:      fixtures.User.URI,
		Response: &id,
	})

	require.Equal(t, http.StatusCreated, statusCode)
	require.NotEmpty(t, id)
}

func Test_Integration_should_be_able_to_retrive_by_id(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.User.Create(t, nil)

	responseBody := user_gateway.GetByIDOutput{}

	token := fixtures.User.Login(t, nil)
	statusCode := fixtures.Get(t, fixtures.GetInput{
		URI:      fixtures.User.URI + id,
		Response: &responseBody,
		Token:    token,
	})
	require.Equal(t, http.StatusOK, statusCode)

	EXPECTED := user_gateway.GetByIDOutput{
		City:           "City",
		State:          "State",
		ZIP:            "ZIP",
		Name:           "Name",
		Email:          "Email",
		Phone:          "Phone",
		Photo:          "Photo",
		TaxID:          "TaxID",
		SocialID:       "SocialID",
		Street:         "Street",
		Complement:     "Complement",
		EmergencyPhone: "EmergencyPhone",
	}

	require.Equal(t, EXPECTED, responseBody)
}

func Test_Integration_should_be_able_to_paginate_if_is_a_backoffice_user(t *testing.T) {
	fixtures.CleanDatabase()

	for i := 0; i < 5; i++ {
		fixtures.User.Create(t, nil)
	}

	id := fixtures.User.Create(t, nil)
	backofficeToken := fixtures.NewBackofficeToken(t, id)

	responseBody := user_gateway.PaginateOutput{}
	statusCode := fixtures.Get(t, fixtures.GetInput{
		URI:      fixtures.User.URI,
		Response: &responseBody,
		Token:    backofficeToken,
	})

	require.Equal(t, http.StatusOK, statusCode)
	require.Len(t, responseBody.Data, 6)
	require.Equal(t, 1, responseBody.MaxPages)

	for i := range responseBody.Data {
		require.NotEmpty(t, responseBody.Data[i].ID)
		require.Equal(t, "Name", responseBody.Data[i].Name)
		require.Equal(t, "City", responseBody.Data[i].City)
		require.Equal(t, "State", responseBody.Data[i].State)
		require.Equal(t, "Photo", responseBody.Data[i].Photo)
	}
}

func Test_Integration_should_not_be_able_to_paginate_if_is_a_common_user(t *testing.T) {
	fixtures.CleanDatabase()

	for i := 0; i < 5; i++ {
		fixtures.User.Create(t, nil)
	}

	id := fixtures.User.Create(t, nil)
	token := fixtures.User.Login(t, &id)

	responseBody := user_gateway.PaginateOutput{}
	statusCode := fixtures.Get(t, fixtures.GetInput{
		URI:      fixtures.User.URI,
		Response: &responseBody,
		Token:    token,
	})

	require.Equal(t, http.StatusOK, statusCode)
	require.Len(t, responseBody.Data, 0)
	require.Equal(t, 0, responseBody.MaxPages)
}

func Test_Integration_should_be_able_to_update(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.User.Create(t, nil)

	Body := user_case.PatchValues{
		ZIP:            "NewZIP",
		Phone:          "NewPhone",
		Email:          "NewEmail",
		EmergencyPhone: "NewEmergencyPhone",
		Street:         "NewStreet",
		SocialID:       "NewSocialID",
		TaxID:          "NewTaxID",
		City:           "NewCity",
		State:          "NewState",
		Complement:     "NewComplement",
		Name:           "NewName",
		Photo:          "NewPhoto",
	}

	token := fixtures.User.Login(t, &id)
	statusCode := fixtures.Patch(t, fixtures.PatchInput{
		Body:  Body,
		URI:   fixtures.User.URI + id,
		Token: token,
	})

	require.Equal(t, http.StatusOK, statusCode)

	found, statusCode := fixtures.User.GetByID(t, id, token)
	require.Equal(t, http.StatusOK, statusCode)

	EXPECTED := user_gateway.GetByIDOutput{
		ZIP:            "NewZIP",
		Phone:          "NewPhone",
		Email:          "NewEmail",
		EmergencyPhone: "NewEmergencyPhone",
		Street:         "NewStreet",
		SocialID:       "NewSocialID",
		TaxID:          "NewTaxID",
		City:           "NewCity",
		State:          "NewState",
		Complement:     "NewComplement",
		Name:           "NewName",
		Photo:          "NewPhoto",
	}

	require.Equal(t, EXPECTED, found)
}

func Test_Integration_should_be_able_to_delete(t *testing.T) {
	fixtures.CleanDatabase()

	id := fixtures.User.Create(t, nil)

	token := fixtures.User.Login(t, &id)
	respBody, statusCode := fixtures.Delete(t, fixtures.DeleteInput{
		URI:   fixtures.User.URI + id,
		Token: token,
	})

	require.Equal(t, statusCode, http.StatusNoContent)
	require.Empty(t, respBody)

	found, statusCode := fixtures.User.GetByID(t, id, token)
	require.Equal(t, statusCode, http.StatusNotFound)

	EXPECTED := user_gateway.GetByIDOutput{}

	require.Equal(t, EXPECTED, found)
}

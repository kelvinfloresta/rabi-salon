package app_context

import (
	"context"
	"rabi-salon/usecases/auth_case/role"
)

var System *AppContext

func init() {
	System = New(context.Background())
	System.Session = &UserSession{
		UserID: "system",
		Name:   "system",
		Login:  "system",
		Role:   role.System,
	}
}

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUser struct {
	FirstName  *string `json:"firstName"`
	MiddleName *string `json:"middleName"`
	LastName   *string `json:"lastName"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
	Number     *string `json:"number"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Review struct {
	Rating *float64 `json:"rating"`
	Review *string  `json:"review"`
}

type ReviewInput struct {
	Token  string  `json:"token"`
	Rating float64 `json:"rating"`
	Review string  `json:"review"`
}

type User struct {
	FirstName  *string `json:"firstName"`
	MiddleName *string `json:"middleName"`
	LastName   *string `json:"lastName"`
	Email      *string `json:"email"`
	Number     *string `json:"number"`
}

type UserInfo struct {
	CurrentUser *bool `json:"currentUser"`
}

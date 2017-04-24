package harvest

import (
	"fmt"
	"time"
)

type UserResponse struct {
	User *User `json:"user"`
}

type User struct {
	ID                           int64     `json:"id"`
	Email                        string    `json:"email"`
	CreatedAt                    time.Time `json:"created_at"`
	IsAdmin                      bool      `json:"is_admin"`
	FirstName                    string    `json:"first_name"`
	LastName                     string    `json:"last_name"`
	Timezone                     string    `json:"timezone"`
	IsContractor                 bool      `json:"is_contractor"`
	Telephone                    string    `json:"telephone"`
	IsActive                     bool      `json:"is_active"`
	HasAccessToAllFutureProjects bool      `json:"has_access_to_all_future_projects"`
	DefaultHourlyRate            float64   `json:"default_hourly_rate"`
	Department                   string    `json:"department"`
	WantsNewsletter              bool      `json:"wants_newsletter"`
	UpdatedAt                    time.Time `json:"updated_at"`
	CostRate                     float64   `json:"cost_rate"`
	WeeklyCapacity               int64     `json:"weekly_capacity"`
	IdentityAccountID            int64     `json:"identity_account_id"`
	IdentityUserID               int64     `json:"identity_user_id"`
}

func (a *API) GetUser(userID int64, args Arguments) (user *User, err error) {
	userResponse := UserResponse{}
	path := fmt.Sprintf("/people/%v", userID)
	err = a.Get(path, args, &userResponse)
	return userResponse.User, err
}

func (a *API) GetUsers(args Arguments) (users []*User, err error) {
	usersResponse := make([]*UserResponse, 0)
	path := fmt.Sprintf("/people")
	err = a.Get(path, args, &usersResponse)
	for _, ur := range usersResponse {
		users = append(users, ur.User)
	}
	return users, err
}

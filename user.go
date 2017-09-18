package harvest

import (
	"fmt"
	"time"
)

type UsersResponse struct {
	Users        []*User `json:"users"`
	PerPage      int64   `json:"per_page"`
	TotalPages   int64   `json:"total_pages"`
	TotalEntries int64   `json:"total_entries"`
	NextPage     *int64  `json:"next_page"`
	PreviousPage *int64  `json:"previous_page"`
	Page         int64   `json:"page"`
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
	user = &User{}
	path := fmt.Sprintf("/users/%v", userID)
	err = a.Get(path, args, &user)
	return user, err
}

func (a *API) GetUsers(args Arguments) (users []*User, err error) {
	usersResponse := UsersResponse{}
	path := fmt.Sprintf("/users")
	err = a.Get(path, args, &usersResponse)
	return usersResponse.Users, err
}

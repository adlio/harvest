package harvest

type User struct {
	ID                int    `xml:"id"`
	FirstName         string `xml:"first-name"`
	LastName          string `xml:"last-name"`
	Email             string `xml:"email"`
	Admin             bool   `xml:"admin"`
	AvatarURL         string `xml:"avatar-url"`
	Timezone          string `xml:"timezone"`
	TimezoneUTCOffset int    `xml:"timezone-utc-offset"`
	TimestampTimers   bool   `xml:"timestamp-timers"`
}

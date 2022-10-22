package ctx

type User struct {
	Id                    int64
	IsBot                 bool
	FirstName             string
	LastName              string
	Username              string
	LanguageCode          string
	IsPremium             bool
	AddedToAttachmentMenu bool
}

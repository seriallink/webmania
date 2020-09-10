package webmania

type Api struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func NewApi(consumerKey, consumerSecret, accessToken, accessSecret string) *Api {

	return &Api{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		AccessToken:    accessToken,
		AccessSecret:   accessSecret,
	}

}

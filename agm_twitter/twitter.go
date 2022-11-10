package agm_twitter

type TwitterConfig struct {
	APIKey            string `json:"ApiKey"`
	APISecretKey      string `json:"ApiSecretKey"`
	AccessToken       string `json:"AccessToken"`
	AccessTokenSecret string `json:"AccessTokenSecret"`
}

func show() {
	//config := oauth1.NewConfig(ApiKey, ApiSecretKey)
	//token := oauth1.NewToken(AccessToken, AccessTokenSecret)
	//httpClient := config.Client(oauth1.NoContext, token)
	//
	//// Twitter client
	//client := twitter.NewClient(httpClient)
	//
	//// Home Timeline
	//tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
	//	Count: 1,
	//})
	//if err != nil {
	//	fmt.Printf("%v", err)
	//	return
	//}
	//fmt.Printf("%v", tweets[0].Text)
}

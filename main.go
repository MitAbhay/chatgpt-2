 package main
 
 import (
	 "github.com/spf13/viper"
	 "github.com/PullRequestInc/go-gpt3"
	 "github.com/spf13/cobra"

 )

 func main() {

	viper.SetConfigFile('.env');
	viper.ReadInConfig();
	apikey := viper.GetString("API_KEY");

	if(apikey == "") {
		panic("API_KEY not found in .env")
	}

	cntxt := context.Background()
	client := gpt3.NewClient(apikey)
	&Cobra.Command{}


 }
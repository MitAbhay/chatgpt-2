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
	rootCmd :=&Cobra.Command{
		Use: "gpt3",
		Short: "ChatGPT in console",
		Run : func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("Say something: ")
				scanner.Scan()
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true

				default:
					GetResponse(client,cntxt,question)
				}
			}

		}
	}



 }
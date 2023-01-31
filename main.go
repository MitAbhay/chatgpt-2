 package main
 
 import (
	 "bufio"
	 "context"
	 "fmt"
	 "os"
	 "github.com/spf13/viper"
	 "github.com/PullRequestInc/go-gpt3"
	 "github.com/spf13/cobra"
 )

 func GetResponse(client gpt3.Client, cntxt context.Context, question string) {
	err := client.CompletionStreamWithEngine(cntxt , gpt3.TextDavinci003Engine , gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens : gpt3.IntPtr(3000),
		Temperature : gpt3.Float32Ptr(0),
		},
		func(resp *gpt3.CompletionResponse){
		fmt.Println(resp.Choices[0].Text)
	})

	if err != nil {
		fmt.Println(err)
	}

 }

 type NullWriter int

 func (NullWriter) Write([]byte) (int, error) { return 0, nil}

 func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig();
	apikey := viper.GetString("API_KEY");
	if(apikey == "") {
		panic("API_KEY not found in .env")
	}
	cntxt := context.Background()
	client := gpt3.NewClient(apikey)
	rootCmd :=&cobra.Command{
		Use: "gpt3",
		Short: "ChatGPT in console",
		Run : func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false
			for !quit {
				fmt.Print("Use ChatGPT in your console : ")
				scanner.Scan()
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true
				default:
					GetResponse(client,cntxt,question)
				}
			}
		},
	}
	rootCmd.Execute()
 }
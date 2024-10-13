package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	app_data "github.com/storozhenko98/Q/pkg/data"
	completion "github.com/storozhenko98/Q/pkg/open_ai"
)

func main() {
	//ensure config file exists
	if err := app_data.SetupAppData(); err != nil {
		printGreeting()
		fmt.Printf("Failed to setup app data: %v\n", err)
		os.Exit(1)
	}
	//check if config file is empty
	empty, err := app_data.CheckIfConfigFileEmpty()
	if err != nil {
		printGreeting()
		fmt.Printf("Failed to check if config file is empty: %v\n", err)
		os.Exit(1)
	}
	if empty {
		printGreeting()
		//ask user for api key
		fmt.Print("Enter your OpenAI API key: ")
		reader := bufio.NewReader(os.Stdin)
		apiKey, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read API key: %v\n", err)
			os.Exit(1)
		}
		apiKey = strings.TrimSpace(apiKey)
		//delete all white spaces at the end or prior to the first non space character
		apiKey = strings.Trim(apiKey, " ")
		if apiKey == "" {
			fmt.Println("API key cannot be empty")
			os.Exit(1)
		}
		if err := app_data.WriteApiKeyToConfig(apiKey); err != nil {
			fmt.Printf("Failed to write API key to config: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("API key saved successfully \n\n All set! You can now use the Q command.")
	}
	if len(os.Args) > 1 {
		// Check for help flags
		if os.Args[1] == "--help" || os.Args[1] == "-h" || os.Args[1] == "-help" || os.Args[1] == "--h" {
			printUsage()
			os.Exit(0)
		}

		// Check for update flag
		if os.Args[1] == "--update" || os.Args[1] == "-update" {
			updateApiKey()
			os.Exit(0)
		}

		// Join all arguments after the program name into a single string
		question := strings.Join(os.Args[1:], " ")
		// Remove surrounding quotes if present
		question = strings.Trim(question, "\"")
		completion.GetCompletion(question)
		fmt.Println()
		os.Exit(0)
	} else {
		// prompt user for question
		fmt.Print("Enter your question: ")
		reader := bufio.NewReader(os.Stdin)
		question, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read question: %v\n", err)
			os.Exit(1)
		}
		question = strings.TrimSpace(question)
		if question == "" {
			fmt.Println("Question cannot be empty")
			os.Exit(1)
		}
		completion.GetCompletion(question)
		fmt.Println()
		os.Exit(0)
	}

}

func updateApiKey() {
	fmt.Print("Enter your new OpenAI API key: ")
	reader := bufio.NewReader(os.Stdin)
	apiKey, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read API key: %v\n", err)
		os.Exit(1)
	}
	apiKey = strings.TrimSpace(apiKey)
	apiKey = strings.Trim(apiKey, " ")
	if apiKey == "" {
		fmt.Println("API key cannot be empty")
		os.Exit(1)
	}
	if err := app_data.UpdateApiKeyInConfig(apiKey); err != nil {
		fmt.Printf("Failed to update API key: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("API key updated successfully")
}

func printUsage() {
	printGreeting()
	fmt.Println("Usage:")
	fmt.Println("  Q [question]                    Ask a question directly")
	fmt.Println("  If your question includes a '?', use quotes: \"query?\"")
	fmt.Println("  For questions without '?' in query, quotes are optional")
	fmt.Println("  Q                           Will ask you for a question")
	fmt.Println("  Q --help or -help or -h or --h   Show this help message")
	fmt.Println("  Q --update or -update        Update your OpenAI API key")
}

func printGreeting() {
	fmt.Println(" --------------------------------------------")
	fmt.Println("|Q (version 0.0.1) says hello!               |")
	fmt.Println("|For help, use Q --help or -help or -h or --h|")
	fmt.Println("|Made with ❤️ by github.com/storozhenko98     |")
	fmt.Println(" -------------------------------------------- ")
}

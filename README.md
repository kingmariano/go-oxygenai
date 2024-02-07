# ⚗️ go-oxygenai
![Build Status](https://github.com/hupe1980/go-huggingface/workflows/build/badge.svg) 
[![Go Reference](https://pkg.go.dev/badge/github.com/hupe1980/go-huggingface.svg)](https://pkg.go.dev/github.com/hupe1980/go-huggingface)
> The OxygenAI  Client in Golang is a modul designed to interact with the openai model repository for free and perform inference tasks using state-of-the-art natural language processing models. Developed in Golang, it provides a seamless and efficient way to integrate openai  models into your Golang applications.

## Installation
```
go get github.com/charlesozo/go-oxygenai
```
## Oxygen AI documentation
> Visit [Discord server](https://oxyapi.uk/discord "Visit OxygenAi discord server") to get the $OXYGEN_API_KEY

___[OxygenAI's Website](https://docs.oxyapi.uk "Visit OxygenAi")___

## How to use
```golang
package main

import (
	"context"
	"fmt"

	oxygenai "github.com/charlesozo/go-oxygenai"
)

func main(){
    client := oxygenai.NewClient("oxy-Jr6vMN7zlVf5Pql0Ua8nTBvfgeoLmzdnf1BPGmllofYu5")
	response, err := client.ChatCompletion(context.Background(), &oxygenai.ChatRequest{
		Messages: []oxygenai.ChatMessage{
		  {
			Role: "user",
			Content: "write an essay on global warming",
		  },
		},
	})

	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(response.Choices[0].Message.Content)
}
```
Output:
```text
response goes here
```

For more example usage, see [_examples](./_examples).

## License
[MIT](LICENCE)
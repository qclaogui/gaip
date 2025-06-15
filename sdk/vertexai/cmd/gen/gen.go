// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	genai2 "github.com/qclaogui/gaip/sdk/vertexai/genai"
	"log"
	"os"
	"strings"

	"google.golang.org/api/iterator"
)

var (
	project   = flag.String("project", "", "project ID")
	location  = flag.String("location", "", "location")
	model     = flag.String("model", "", "model")
	streaming = flag.Bool("stream", false, "using the streaming API")
)

func main() {
	flag.Parse()
	if *project == "" || *location == "" || *model == "" {
		log.Fatal("need -project, -location, and -model")
	}

	ctx := context.Background()
	client, err := genai2.NewClient(ctx, *project, *location)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel(*model)
	model.SafetySettings = []*genai2.SafetySetting{
		{
			Category:  genai2.HarmCategorySexuallyExplicit,
			Threshold: genai2.HarmBlockLowAndAbove,
		},
		{
			Category:  genai2.HarmCategoryDangerousContent,
			Threshold: genai2.HarmBlockLowAndAbove,
		},
	}

	text := strings.Join(flag.Args(), " ")
	if *streaming {
		iter := model.GenerateContentStream(ctx, genai2.Text(text))
		for {
			res, err := iter.Next()
			if errors.Is(err, iterator.Done) {
				break
			}
			if err != nil {
				showError(err)
			}
			showJSON(res)
			fmt.Println("---")
		}
	} else {
		res, err := model.GenerateContent(ctx, genai2.Text(text))
		if err != nil {
			showError(err)
		}
		showJSON(res)
	}
}

func showJSON(x any) {
	bs, err := json.MarshalIndent(x, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bs)
}

func showError(err error) {
	var berr *genai2.BlockedError
	if errors.As(err, &berr) {
		fmt.Println("ERROR:")
		showJSON(err)
	} else {
		fmt.Printf("ERROR: %s\n", err)
	}
	os.Exit(1)
}

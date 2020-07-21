# Mailchimp-go
[![GoDoc][godoc-img]][godoc-url]

## Introduction
Golang client for [MailChimp API 3.0](http://developer.mailchimp.com/documentation/mailchimp/).

## Install
Install with `go get`:

```bash
$ go get github.com/kozjat/mailchimp-go
```

## Usage
```go
package main

import (
	"fmt"
	"os"

	"github.com/kozjat/mailchimp-go"
)

const apiKey = "YOUR_API_KEY_HERE"

func main() {
	client := gochimp3.New(apiKey)

	// Audience ID
	// https://mailchimp.com/help/find-audience-id/
	listID := "7f12f9b3fz"

	// Fetch list
	list, err := client.GetList(listID, nil)
	if err != nil {
		fmt.Printf("Failed to get list %s", listID)
		os.Exit(1)
	}

	// Add subscriber
	req := &gochimp3.MemberRequest{
		EmailAddress: "test@mail.com",
		Status:       "subscribed",
	}

	if _, err := list.CreateMember(req); err != nil {
		fmt.Printf("Failed to subscribe %s", req.EmailAddress)
		os.Exit(1)
	}
}
```

[godoc-img]:      https://godoc.org/github.com/kozjat/mailchimp-go?status.svg
[godoc-url]:      https://godoc.org/github.com/kozjat/mailchimp-go

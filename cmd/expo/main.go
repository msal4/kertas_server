package main

import (
	"fmt"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

func main() {
	pushToken, err := expo.NewExponentPushToken("ExponentPushToken[vJSvQUEEgBsw22mHC6RFkq]")
	if err != nil {
		panic(err)
	}

	client := expo.NewPushClient(nil)

	res, err := client.Publish(&expo.PushMessage{
		To:       []expo.ExponentPushToken{pushToken},
		Title:    "of course it will",
		Body:     "hope this is working",
		Data:     map[string]string{"something": "wow"},
		Sound:    "default",
		Priority: expo.DefaultPriority,
	})
	if err != nil {
		panic(err)
	}

	if err := res.ValidateResponse(); err != nil {
		panic(err)
	}

	fmt.Println(res.Details)
}

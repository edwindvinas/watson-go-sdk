package main

import (
	"fmt"

	"github.com/IBM/go-sdk-core/core"
	"github.com/edwindvinas/watson-go-sdk/assistantv2"
)

func main() {
	// Instantiate the Watson AssistantV2 service
	authenticator := &core.IamAuthenticator{
		ApiKey: "MbThJzh3DPOgqRvF9kljaE9f2RNocPc68L1EZmB2P1DQ",
	}
	service, serviceErr := assistantv2.
		NewAssistantV2(&assistantv2.AssistantV2Options{
			URL:           "https://gateway.watsonplatform.net/assistant/api",
			Version:       "2017-04-21",
			Authenticator: authenticator,
		})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	/* CREATE SESSION */

	assistantID := "c7019059-edd8-4915-b89b-461318e79590"
	// Call the assistant CreateSession method
	createSessionResult, _, responseErr := service.
		CreateSession(&assistantv2.CreateSessionOptions{
			AssistantID: core.StringPtr(assistantID),
		})

	if responseErr != nil {
		panic(responseErr)
	}
	sessionID := createSessionResult.SessionID

	// 	/* MESSAGE */

	// Call the assistant Message method
	_, response, responseErr := service.
		Message(&assistantv2.MessageOptions{
			AssistantID: core.StringPtr(assistantID),
			SessionID:   sessionID,
			Input: &assistantv2.MessageInput{
				Text: core.StringPtr("I have a slow internet"),
			},
			Context: &assistantv2.MessageContext{
				Global: &assistantv2.MessageContextGlobal{
					System: &assistantv2.MessageContextGlobalSystem{
						UserID: core.StringPtr("dummy"),
					},
				},
			},
		})

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	core.PrettyPrint(response.GetResult(), "Message")

	// 	/* DELETE SESSION */

	// Call the assistant DeleteSession method
	_, responseErr = service.
		DeleteSession(&assistantv2.DeleteSessionOptions{
			AssistantID: core.StringPtr(assistantID),
			SessionID:   sessionID,
		})

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}
	fmt.Println("Session deleted successfully")
}

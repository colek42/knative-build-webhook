package main

import (
	"fmt"

	"net/http"

	"gopkg.in/go-playground/webhooks.v5/github"
)

const (
	path = "/webhook"
)

func main() {
	hook, _ := github.New(github.Options.Secret(""))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			prAction := pullRequest.Action
			if prAction == "closed" && pullRequest.PullRequest.Merged == true {
				fmt.Println("Merged")
			} else {
				fmt.Println("Not merged")
			}
		}
	})
	http.ListenAndServe(":8787", nil)
}

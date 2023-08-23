package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/crowdstrike/gofalcon/examples"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/incidents"
)

func main() {
	commonAuthFlags := examples.SetupAuthFlags()
	// filter := flag.String("filter", "", "Filter for alerts")

	flag.Parse()
	commonAuthFlags.PromptForRequiredFlags()

	client, err := falcon.NewClient(&falcon.ApiConfig{
		ClientId:     commonAuthFlags.ClientId,
		ClientSecret: commonAuthFlags.ClientSecret,
		MemberCID:    commonAuthFlags.MemberCID,
		Cloud:        falcon.Cloud(commonAuthFlags.Cloud),
		Context:      context.Background(),
	})

	if err != nil {
		panic(err)
	}

	desc := "timestamp.desc"
	res, err := client.Incidents.CrowdScore(&incidents.CrowdScoreParams{
		Context: context.Background(),
		Sort:    &desc,
	})
	if err != nil {
		panic(err)
	}
	payload := res.GetPayload()
	if err = falcon.AssertNoError(payload.Errors); err != nil {
		panic(err)
	}
	fmt.Printf("As of %s your CrowdScore is %d.\n",
		payload.Resources[0].Timestamp.String(), *payload.Resources[0].Score)
}

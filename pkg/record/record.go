package record

import (
	"context"
	gocorona "github.com/itsksaurabh/go-corona"
)

type Record struct {
	Confirmed int
	Deaths    int
	Recovered int
}

func Create() Record {
	client := gocorona.Client{}
	backgroundContext := context.Background()
	collectedData, retrieveFailure := client.GetLatestData(backgroundContext)
	latestData := collectedData.Data
	expectNoError(retrieveFailure)
	return Record{
		Confirmed: latestData.Confirmed,
		Deaths:    latestData.Deaths,
		Recovered: latestData.Recovered,
	}
}

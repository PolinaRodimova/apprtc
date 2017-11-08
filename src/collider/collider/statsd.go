package collider


import (
	"log"
	"github.com/DataDog/datadog-go/statsd"
)


func statsClient() *statsd.Client  {
	c, err := statsd.New("dd-agent:8125")
	if err != nil {
		log.Fatal(err)
	}
	return c;
}



func writeStats(action string) {
	c := statsClient()
	c.Namespace = "blitz.collider."
	err := c.Incr(action, nil, 1)
	if err != nil {
		log.Fatal(err)

	}
}


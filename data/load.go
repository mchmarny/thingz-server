package data

import (
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/mchmarny/thingz-server/config"
)

const (
	KAFKA_CONSUMER_GROUP = "thingzgroup"
)

// LoadFromKafka
func LoadFromKafka() {

	topic := config.Config.KafkaTopic
	brokers := config.Config.KafkaBrokers

	if len(topic) < 1 || len(brokers) < 1 {
		panic("Invalid arguments. Both topic and brokers required")
	}

	clientId, _ := os.Hostname()
	brokerList := strings.Split(brokers, ",")
	log.Printf("Subscribing to topic:%s > %v", topic, brokerList)

	client, err := sarama.NewClient(clientId, brokerList, nil)
	if err != nil {
		log.Fatalf("Error while creating client: %v", err)
		panic(err)
	}
	defer client.Close()

	// make sure we get messages from all partitions
	partIDs, err := client.Partitions(topic)
	if err != nil {
		log.Printf("Error on listing partitions: %v", err)
		panic(err)
	}

	consumers := make([]*sarama.Consumer, 0)

	for i, p := range partIDs {
		log.Printf("Creating consumer[%d]%d...", i, p)
		cons, err := sarama.NewConsumer(
			client,
			topic,
			p,
			KAFKA_CONSUMER_GROUP,
			sarama.NewConsumerConfig(),
		)

		if err != nil {
			log.Printf("Consumer[%d] failed: %v", i, err)
			panic(err)
		} else {
			consumers = append(consumers, cons)
			log.Printf("Consumer[%d] ready", i)
			defer cons.Close()
		}
	}

	for {
		//consumerLoop:
		for i, c := range consumers {
			select {
			case e := <-c.Events():
				if e.Err != nil {
					log.Printf("Event error: %d - %v", i, e.Err.Error())
				} else {
					// TODO: send to DB load channel
					log.Printf("T:%s K:%s O:%d P:%d V:%s",
						e.Topic,
						string(e.Key),
						e.Offset,
						e.Partition,
						string(e.Value),
					)
				} // error
			} // select
		} // for each consumer
	} // for

}

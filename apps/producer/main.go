// example of producer
package main

import (
	"log"
	"os"
	"strings"

	kafka "github.com/yoshihiro-shu/financial-bot/repository/apache_kafka"
)

var (
	brokers = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
)

func main() {
	producer, err := kafka.NewProducer(brokers)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	log.Printf("==> start producing %#v\n", producer)
	log.Printf("==> status of producing %#v", producer.TxnStatus())

	topic := "news"
	key := []byte("key")
	value := []byte("value")
	err = producer.SendMessage(topic, key, value)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("success")
}

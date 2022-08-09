package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
)

type DataBase struct {
	Port         int
	Db_url       string
	Jaeger_url   string
	Sentry_url   string
	Kafka_broker string
	Some_app_id  string
	Some_app_key string
}

func main() {
	Connect := DataBase{}

	flag.IntVar(&Connect.Port, "port", 8080, "port")
	flag.StringVar(&Connect.Db_url, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "db_url")
	flag.StringVar(&Connect.Jaeger_url, "jaeger_url", "http://jaeger:16686", "jaeger_url")
	flag.StringVar(&Connect.Sentry_url, "sentry_url", "http://sentry:9000", "sentry_url")
	flag.StringVar(&Connect.Kafka_broker, "kafka_broker", "kafka:9092", "kafka_broker")
	flag.StringVar(&Connect.Some_app_id, "some_app_id", "testid", "some_app_id")
	flag.StringVar(&Connect.Some_app_key, "some_app_key", "testkey", "some_app_key")
	flag.Parse()

	_, err := url.Parse(Connect.Db_url)
	if err != nil {
		panic(err)
	}
	_, err = url.Parse(Connect.Jaeger_url)
	if err != nil {
		panic(err)
	}
	_, err = url.Parse(Connect.Sentry_url)
	if err != nil {
		panic(err)
	}

	Bufer, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer Bufer.Close()
	DataJsonFile := DataBase{}

	err = json.NewDecoder(Bufer).Decode(&DataJsonFile)
	if err != nil {
		fmt.Println("Пустой файл, записаны стандартные значения")
		DataJsonFile = Connect
		DataJsonFlag, _ := json.Marshal(DataJsonFile)
		os.WriteFile("data.json", DataJsonFlag, 0777)

	}
	_, err = url.Parse(DataJsonFile.Db_url)
	if err != nil {
		panic(err)
	}
	_, err = url.Parse(DataJsonFile.Jaeger_url)
	if err != nil {
		panic(err)
	}
	_, err = url.Parse(DataJsonFile.Sentry_url)
	if err != nil {
		panic(err)
	}

	fmt.Println(DataJsonFile)

}

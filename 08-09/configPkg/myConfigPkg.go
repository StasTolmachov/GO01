package MyConfigPkg

import (
	"flag"
	"fmt"
	"net/url"
)

func MyConfigFunc() {
	fmt.Println("MyConfigFunc")

	type DataBase struct {
		port         int
		db_url       string
		jaeger_url   string
		sentry_url   string
		kafka_broker string
		some_app_id  string
		some_app_key string
	}

	connect := DataBase{}
	flag.IntVar(&connect.port, "port", 8080, "port")
	flag.StringVar(&connect.db_url, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "db_url")
	flag.StringVar(&connect.jaeger_url, "jaeger_url", "http://jaeger:16686", "jaeger_url")
	flag.StringVar(&connect.sentry_url, "sentry_url", "http://sentry:9000", "sentry_url")
	flag.StringVar(&connect.kafka_broker, "kafka_broker", "kafka:9092", "kafka_broker")
	flag.StringVar(&connect.some_app_id, "some_app_id", "testid", "some_app_id")
	flag.StringVar(&connect.some_app_key, "some_app_key", "testkey", "some_app_key")

	flag.Parse()

	_, err := url.Parse(connect.db_url)
	if err != nil {
		panic(err)
	}
	_, err = url.Parse(connect.jaeger_url)
	if err != nil {
		panic(err)
	}
	_, err = url.Parse(connect.sentry_url)
	if err != nil {
		panic(err)
	}

	fmt.Println(connect.port, connect.db_url, connect.jaeger_url, connect.sentry_url, connect.kafka_broker, connect.some_app_id, connect.some_app_key)

}

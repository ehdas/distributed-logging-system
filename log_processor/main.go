package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
)

func main() {

	kafkaBroker := "kafka:9092"

	db, err := sql.Open("postgres", "postgresql://root:secret@postgres:5432/grpc-kafka?sslmode=disable")
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	// connect to Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   "logs",
		GroupID: "log-consumer-group",
	})
	defer reader.Close()

	log.Println("Listening to Kafka topic:", "logs")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			continue
		}

		// separate message
		parts := strings.SplitN(string(msg.Value), ":", 2)
		serviceInfo := strings.Trim(parts[0], " []")
		logMessage := ""
		if len(parts) > 1 {
			logMessage = strings.TrimSpace(parts[1])
		}

		// insert in database
		_, err = db.Exec(`INSERT INTO logs (service_name, level, message, timestamp) VALUES ($1, $2, $3, $4)`,
			serviceInfo,
			"INFO",
			logMessage,
			msg.Time,
		)
		if err != nil {
			log.Printf("failed to insert log: %v", err)
		} else {
			log.Printf("Log stored in DB: %s", msg.Value)
		}
	}
}

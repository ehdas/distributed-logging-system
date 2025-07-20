package main

import (
	"context"
	"fmt"
	"github.com/ehdas/distributed-logging-system/log_collector/pb"
	"github.com/ehdas/distributed-logging-system/log_collector/util"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedLogServiceServer
	kafkaWriter *kafka.Writer
}

func (s *server) SendLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	msg := fmt.Sprintf("[%s] %s: %s", req.ServiceName, req.Level, req.Message)

	err := s.kafkaWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte(req.ServiceName),
		Value: []byte(msg),
		Time:  time.Unix(req.Timestamp, 0),
	})
	if err != nil {
		log.Printf("failed to write to kafka: %v", err)
		return &pb.LogResponse{Status: "ERROR"}, err
	}

	log.Printf("Log received and sent to Kafka: %s", msg)
	return &pb.LogResponse{Status: "OK"}, nil
}

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{config.KAFKABROKER},
		Topic:    "logs",
		Balancer: &kafka.LeastBytes{},
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen ")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLogServiceServer(grpcServer, &server{kafkaWriter: writer})

	log.Info().Msg("gRPC server started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}

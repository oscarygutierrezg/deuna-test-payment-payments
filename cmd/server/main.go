package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"payment-payments-api/internal/api"
	"payment-payments-api/internal/config"
	"payment-payments-api/internal/kafka/consumer"
	"payment-payments-api/internal/kafka/producer"
	"payment-payments-api/internal/repositories"
	"payment-payments-api/internal/services"
	"sync"
)

func init() {
	flag.Int("port", 5002, "server port.")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.AutomaticEnv()
	_ = viper.BindPFlags(pflag.CommandLine)
}

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := config.InitDB(cfg)
	if err != nil {
		panic(err)
	}
	err = config.Migrate(db)
	if err != nil {
		panic(err)
	}

	paymentProducer, err := producer.NewPaymentProducer(cfg)
	if err != nil {
		log.Fatalf("Failed to create Kafka paymentProducer: %v", err)
		panic(err)
	}

	paymentConsumer, err := consumer.NewPaymentConsumer(cfg)
	if err != nil {
		log.Fatalf("Failed to create Kafka paymentConsumer: %v", err)
		panic(err)
	}

	paymentRepository := repositories.NewPaymentRepository(db)

	paymentService := services.NewPaymentService(paymentRepository, paymentProducer)

	services := &services.Services{
		Payment: paymentService,
	}

	server := api.NewServer(services)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		err = server.Run(fmt.Sprintf(":%d", viper.GetInt("port")))
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		defer wg.Done()

		paymentConsumer.Consume(services)
	}()

	wg.Wait()
}

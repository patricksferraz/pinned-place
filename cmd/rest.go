/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Netflix/go-env"
	"github.com/c-4u/pinned-place/app/rest"
	"github.com/c-4u/pinned-place/infra/client/kafka"
	"github.com/c-4u/pinned-place/infra/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gorm.io/gorm/logger"
)

// restCmd represents the rest command
func restCmd() *cobra.Command {
	var conf Config

	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		log.Fatal(err)
	}

	restCmd := &cobra.Command{
		Use:   "rest",
		Short: "Run rest Service",

		Run: func(cmd *cobra.Command, args []string) {
			l := logger.Error
			if *conf.Db.Debug {
				l = logger.Info
			}

			orm, err := db.NewDbOrm(*conf.Db.Dsn, l)
			if err != nil {
				log.Fatal(err)
			}

			if err = orm.Migrate(); err != nil {
				log.Fatal(err)
			}
			log.Printf("migration did run successfully")

			deliveryChan := make(chan ckafka.Event)
			kp, err := kafka.NewKafkaProducer(*conf.Kafka.Servers, deliveryChan)
			if err != nil {
				log.Fatal("cannot start kafka producer", err)
			}

			go kp.DeliveryReport()
			rest.StartRestServer(orm, kp, *conf.RestPort)
		},
	}

	return restCmd
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(basepath + "/../.env")
		if err != nil {
			log.Printf("Error loading .env files")
		}
	}

	rootCmd.AddCommand(restCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

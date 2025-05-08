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
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/patricksferraz/pinned-place/app/rest"
	"github.com/patricksferraz/pinned-place/infra/client/kafka"
	"github.com/patricksferraz/pinned-place/infra/db"
	"github.com/spf13/cobra"
	"gorm.io/gorm/logger"
)

// allCmd represents the all command
func allCmd() *cobra.Command {
	var conf Config

	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		log.Fatal(err)
	}

	allCmd := &cobra.Command{
		Use:   "all",
		Short: "Run both gRPC and rest servers",

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

	return allCmd
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

	rootCmd.AddCommand(allCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

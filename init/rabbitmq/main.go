package main

import (
	"log"

	"github.com/zeromicro/go-queue/rabbitmq"
)

func main() {
	conf := rabbitmq.RabbitConf{
		Host:     "localhost",
		Port:     5672,
		Username: "guest",
		Password: "guest",
	}
	admin := rabbitmq.MustNewAdmin(conf)
	exchangeConf := rabbitmq.ExchangeConf{
		ExchangeName: "verify_code",
		Type:         "direct",
		Durable:      true,
		AutoDelete:   false,
		Internal:     false,
		NoWait:       false,
	}

	err := admin.DeclareExchange(exchangeConf, nil)
	if err != nil {
		log.Fatal(err)
	}

	queueConf := rabbitmq.QueueConf{
		Name:       "send_code2email",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
	}
	err = admin.DeclareQueue(queueConf, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = admin.Bind("send_code2email", "send_code2email", "verify_code", false, nil)
	if err != nil {
		log.Fatal(err)
	}

	queueConf1 := rabbitmq.QueueConf{
		Name:       "send_code2phone",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
	}
	err = admin.DeclareQueue(queueConf1, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = admin.Bind("send_code2phone", "send_code2phone", "verify_code", false, nil)
	if err != nil {
		log.Fatal(err)
	}
}

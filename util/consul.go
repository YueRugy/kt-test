package util

import (
	consulapi "github.com/hashicorp/consul/api"
	"k8s.io/apimachinery/pkg/util/uuid"
	"log"
	"strconv"
)

var (
	ConsulClient *consulapi.Client
	err          error
	ServiceName  string
	ServiceID    string
	ServicePort  int
)

func init() {
	{
		ServiceID = "userservice" + string(uuid.NewUUID())
	}
	config := consulapi.DefaultConfig()
	config.Address = "localhost:8500"
	ConsulClient, err = consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
}

func Register() {
	reg := consulapi.AgentServiceRegistration{

		ID:      ServiceID,
		Name:    ServiceName,
		Address: "localhost",
		Port:    ServicePort,
		Tags:    []string{"primary"},
	}
	check := consulapi.AgentServiceCheck{
		Interval: "5s",
		HTTP:     "http://192.168.0.102:" + strconv.Itoa(ServicePort) + "/health",
	}
	reg.Check = &check
	err = ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}

func Deregister() {
	//fmt.Println("exit")
	err = ConsulClient.Agent().ServiceDeregister(ServiceID)
	if err != nil {
		log.Fatal(err)
	}
}

func SetNameAndPort(name string, port int) {
	ServicePort = port
	ServiceName = name
}

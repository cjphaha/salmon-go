package consul

import (
	"github.com/hashicorp/consul/api"
	"log"
)

type ConsulClient struct {
	api.Client
}

func NewConsulClient(consulURL string) (consulClient *ConsulClient) {
	config := api.DefaultConfig()
	config.Address = consulURL
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	return &ConsulClient{
		*client,
	}
}

func (this *ConsulClient)RegService(reg  api.AgentServiceRegistration) (err error){
	err = this.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (this *ConsulClient)Unregservice(id string){
	this.Agent().ServiceDeregister(id)
}

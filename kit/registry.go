package kit

import (
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

func RegService(port int,id, name, version, method, consulUrl string) {
	config := api.DefaultConfig()
	config.Address = consulUrl
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	reg := api.AgentServiceRegistration{
		ID: id,
		Name:           name,
		Address:        "host.docker.internal",
		Port:           port,
		Tags:           []string{version},
		Meta:           map[string]string{"method":method},
		Check: &api.AgentServiceCheck{
			Interval: "5s",
			HTTP: "http://host.docker.internal:" + strconv.Itoa(port) +"/api/health",
		},
	}
	err = client.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}

func Unregservice(id string, client *api.Client){
	client.Agent().ServiceDeregister(id)
}
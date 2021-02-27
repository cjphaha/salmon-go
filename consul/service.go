package consul

import "github.com/hashicorp/consul/api"


func (this *ConsulClient)GetServiceList() (list []*api.AgentService){
	nodes, _, _ := this.Catalog().Nodes(&api.QueryOptions{})
	for _, v := range nodes {
		servces, _, _ := this.Catalog().NodeServiceList(v.Node, &api.QueryOptions{})
		for _, v := range servces.Services {
			if v.Service  == "consul" {
				continue
			}
			list = append(list, v)
		}
	}
	return
}
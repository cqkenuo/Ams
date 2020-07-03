package subdomain

import "Ams/model"

func SubdomainService(domain chan *model.Domains){
	for{
		item := <- domain
	}
}

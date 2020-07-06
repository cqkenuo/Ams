package domain_info

import (
	"Ams/config"
	"github.com/miekg/dns"
)

type dnsResolve struct {
	client *dns.Client
}

func NewDnsResolve() *dnsResolve {
	client := new(dns.Client)
	return &dnsResolve{client: client}
}

func (dr *dnsResolve) LookUpHost(domain string, dnsServer string) []string {
	m := new(dns.Msg)
	m.SetQuestion(domain+".", dns.TypeA)
	m.RecursionDesired = true
	r, _, err := dr.client.Exchange(m, dnsServer)
	if err != nil {
		return []string{}
	}
	if r.Rcode != dns.RcodeSuccess {
		return []string{}
	}
	var resolves []string
	for _, a := range r.Answer {
		if rCord, ok := a.(*dns.A); ok {
			resolves = append(resolves, rCord.A.String())
		}
	}
	return resolves
}

//字符串切片去重
func strSliceUnique(v []string) []string {
	tmpSlice := make([]string, 0)
	tmp := make(map[string]interface{})
	for _, val := range v {
		if _, ok := tmp[val]; !ok {
			tmpSlice = append(tmpSlice, val)
			tmp[val] = nil
		}
	}
	return tmpSlice
}

func (dr *dnsResolve) LookUp(domain string) []string {
	setting := config.LoadConfig()
	var resolves []string
	for _, dnsServer := range setting.DnsServers {
		resolves = append(resolves, dr.LookUpHost(domain, dnsServer)...)
	}
	resolves = strSliceUnique(resolves)
	return resolves
}

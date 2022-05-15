package check

import (
	"fmt"
	"github.com/miekg/dns"
	"golang.org/x/sync/semaphore"
	"net"
)

type Item struct {
	IP        net.IP
	Blacklist string
	Host      string
}

func Check(sem *semaphore.Weighted, item Item, nameserver string) (blacklisted bool, responses []string, err error) {
	defer sem.Release(1)

	client := new(dns.Client)
	isIPv6 := item.IP.To4() == nil

	var t = dns.TypeA
	m := new(dns.Msg)
	if isIPv6 {
		t = dns.TypeAAAA
	}
	m.SetQuestion(item.Blacklist, t)
	m.RecursionDesired = true

	var r *dns.Msg
	r, _, err = client.Exchange(m, nameserver)
	if err != nil {
		return blacklisted, responses, err
	}

	if r.Rcode != dns.RcodeSuccess {
		err = fmt.Errorf("received %v but expected %v (RcodeSuccess) for %s", r.Rcode, dns.RcodeSuccess, item.Blacklist)
		return blacklisted, responses, err
	}

	if len(r.Answer) > 0 {
		blacklisted = true
	}

	for _, a := range r.Answer {
		switch a.(type) {
		case *dns.A:
			responses = append(responses, a.(*dns.A).A.String())
		case *dns.AAAA:
			responses = append(responses, a.(*dns.AAAA).AAAA.String())
		}
	}

	return blacklisted, responses, err

}

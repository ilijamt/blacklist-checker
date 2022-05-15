package blacklist_checker

import (
	"github.com/ilijamt/blacklist_checker/internal/utils"
	"io/ioutil"
	"strings"
)

var dnsbls = []string{
	"access.redhawk.org",
	"b.barracudacentral.org",
	"bl.spamcop.net",
	"blackholes.mail-abuse.org",
	"bogons.cymru.com",
	"cbl.abuseat.org",
	"cbl.anti-spam.org.cn",
	"cdl.anti-spam.org.cn",
	"combined.njabl.org",
	"csi.cloudmark.com",
	"db.wpbl.info",
	"dnsbl.dronebl.org",
	"dnsbl.inps.de",
	"dnsbl.njabl.org",
	"dnsbl.sorbs.net",
	"drone.abuse.ch",
	"dsn.rfc-ignorant.org",
	"dul.dnsbl.sorbs.net",
	"dyna.spamrats.com",
	"http.dnsbl.sorbs.net",
	"httpbl.abuse.ch",
	"ips.backscatterer.org",
	"ix.dnsbl.manitu.net",
	"korea.services.net",
	"misc.dnsbl.sorbs.net",
	"multi.surbl.org",
	"netblock.pedantic.org",
	"noptr.spamrats.com",
	"opm.tornevall.org",
	"pbl.spamhaus.org",
	"psbl.surriel.com",
	"query.senderbase.org",
	"rbl-plus.mail-abuse.org",
	"rbl.efnetrbl.org",
	"rbl.interserver.net",
	"rbl.spamlab.com",
	"rbl.suresupport.com",
	"relays.mail-abuse.org",
	"sbl.spamhaus.org",
	"short.rbl.jp",
	"smtp.dnsbl.sorbs.net",
	"socks.dnsbl.sorbs.net",
	"spam.dnsbl.sorbs.net",
	"spam.spamrats.com",
	"spamguard.leadmon.net",
	"spamrbl.imp.ch",
	"tor.dan.me.uk",
	"ubl.unsubscore.com",
	"virbl.bit.nl",
	"virus.rbl.jp",
	"web.dnsbl.sorbs.net",
	"wormrbl.imp.ch",
	"xbl.abuseat.org",
	"xbl.spamhaus.org",
	"zen.spamhaus.org",
	"zombie.dnsbl.sorbs.net",
}

func GetDNSBLs(file string) ([]string, error) {
	var exists, _ = utils.FileExists(file)
	if !exists {
		return dnsbls, nil
	}

	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return []string{}, err
	}

	var vals []string
	for _, s := range strings.Split(string(fileBytes), "\n") {
		vals = append(vals, strings.TrimSpace(s))
	}

	return vals, nil
}

Blacklist Checker
================

Check if your IP or CIDR is blacklisted or not.

There are probably faster ways to do this so if anyone want's to cleanup or send a PR feel free to do so

## Installation

## Getting Started with Blacklist Checker

### Requirements

* [Golang](https://golang.org/dl/) >= 1.8

## Install

### Pre-compiled binary

#### manually

Download the pre-compiled binaries from the [releases](https://github.com/ilijamt/blacklist-checker/releases) page.

#### homebrew

```bash
brew tap ilijamt/tap
brew install blacklist-checker
```

### Help
```bash
$ blacklist-checker
A simple tool that helps you check if your IP or CIDR is blacklisted or not.

Usage:
  blacklist-checker [command]

Available Commands:
  check       Check available blacklists.
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List available blacklists.
  version     Shows the version of the application

Flags:
      --dsnbl string   DNSBL file to use, if empty it uses the internal list, should be a list of DNSBL to use, each one on a new line
  -h, --help           help for blacklist-checker

Use "blacklist-checker [command] --help" for more information about a command.                                                                                                                                                           /0.1s
```

#### IP 

```bash
$ blacklist-checker ip 46.217.104.208
12:51AM INF processing dsnbl=56 queries=56
12:51AM WRN  blacklisted=true dnsbl=b.barracudacentral.org ip=46.217.104.208 responses=["127.0.0.2"]
12:51AM WRN  blacklisted=true dnsbl=pbl.spamhaus.org ip=46.217.104.208 responses=["127.0.0.11"]
12:51AM WRN  blacklisted=true dnsbl=zen.spamhaus.org ip=46.217.104.208 responses=["127.0.0.11"]
12:51AM INF Finished blacklisted=3 queries=56                                                                                                                                                                                            /2.6s
```

#### CIDR
```bash
$ blacklist-checker check cidr 46.217.104.208/28
12:51AM INF processing dsnbl=56 queries=896
12:51AM WRN  blacklisted=true dnsbl=b.barracudacentral.org ip=46.217.104.215 responses=["127.0.0.2"]
12:51AM WRN  blacklisted=true dnsbl=b.barracudacentral.org ip=46.217.104.218 responses=["127.0.0.2"]
...
12:51AM WRN  blacklisted=true dnsbl=b.barracudacentral.org ip=46.217.104.223 responses=["127.0.0.2"]
12:51AM WRN  blacklisted=true dnsbl=pbl.spamhaus.org ip=46.217.104.208 responses=["127.0.0.11"]
12:51AM WRN  blacklisted=true dnsbl=pbl.spamhaus.org ip=46.217.104.209 responses=["127.0.0.11"]
...
12:51AM WRN  blacklisted=true dnsbl=spam.dnsbl.sorbs.net ip=46.217.104.209 responses=["127.0.0.6"]
12:51AM WRN  blacklisted=true dnsbl=spam.dnsbl.sorbs.net ip=46.217.104.221 responses=["127.0.0.6"]
12:51AM WRN  blacklisted=true dnsbl=spam.dnsbl.sorbs.net ip=46.217.104.222 responses=["127.0.0.6"]
...
12:52AM WRN  blacklisted=true dnsbl=zen.spamhaus.org ip=46.217.104.222 responses=["127.0.0.11"]
12:52AM WRN  blacklisted=true dnsbl=zen.spamhaus.org ip=46.217.104.218 responses=["127.0.0.11"]
12:52AM INF Finished blacklisted=51 queries=896                                                                                                                                                                                         /17.8s
```

#### Blacklist file format
If you want to create your on black list file you want to use, you can use the format bellow, and specify it with the `--dsnbl` flag

```shell
$ cat my-dnsbl-list
b.barracudacentral.org
sbl.spamhaus.org
```

#### Blacklists

Currently there are 56 blacklists in [blacklists.go](blacklists.go)

```bash
$ blacklist-checker list
access.redhawk.org
b.barracudacentral.org
bl.spamcop.net
blackholes.mail-abuse.org
bogons.cymru.com
cbl.abuseat.org
cbl.anti-spam.org.cn
cdl.anti-spam.org.cn
combined.njabl.org
csi.cloudmark.com
db.wpbl.info
dnsbl.dronebl.org
dnsbl.inps.de
dnsbl.njabl.org
dnsbl.sorbs.net
drone.abuse.ch
dsn.rfc-ignorant.org
dul.dnsbl.sorbs.net
dyna.spamrats.com
http.dnsbl.sorbs.net
httpbl.abuse.ch
ips.backscatterer.org
ix.dnsbl.manitu.net
korea.services.net
misc.dnsbl.sorbs.net
multi.surbl.org
netblock.pedantic.org
noptr.spamrats.com
opm.tornevall.org
pbl.spamhaus.org
psbl.surriel.com
query.senderbase.org
rbl-plus.mail-abuse.org
rbl.efnetrbl.org
rbl.interserver.net
rbl.spamlab.com
rbl.suresupport.com
relays.mail-abuse.org
sbl.spamhaus.org
short.rbl.jp
smtp.dnsbl.sorbs.net
socks.dnsbl.sorbs.net
spam.dnsbl.sorbs.net
spam.spamrats.com
spamguard.leadmon.net
spamrbl.imp.ch
tor.dan.me.uk
ubl.unsubscore.com
virbl.bit.nl
virus.rbl.jp
web.dnsbl.sorbs.net
wormrbl.imp.ch
xbl.abuseat.org
xbl.spamhaus.org
zen.spamhaus.org
zombie.dnsbl.sorbs.net                                                                                                                                                                                                                   /0.1s
```

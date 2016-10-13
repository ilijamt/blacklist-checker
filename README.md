Blacklist Checker
================

Check if your IP or CIDR is blacklisted or not.

There are probably faster ways to do this so if anyone want's to cleanup or send a PR feel free to do so

## Installation

If you don't want to compile your own version, you can use the following repository to install it 

### Debian

```bash
echo "deb http://packages.matoski.com/ debian main" | sudo tee /etc/apt/sources.list.d/packages-matoski-com.list
curl -s http://packages.matoski.com/keyring.gpg | sudo apt-key add -
sudo apt-get update
sudo apt-get install blacklist-checker
```

## Getting Started with Blacklist Checker

### Requirements

* [Golang](https://golang.org/dl/) >= 1.6
* [Glide](https://github.com/Masterminds/glide) >= 0.10.0

### Autocomplete 

You know what to do with this

* [Bash](contrib/blacklist-checker.bash)
* [Zsh](contrib/blacklist-checker.zsh)

### Dependencies

This project uses glide to manage dependencies so download them before trying to build/install by running 

```bash
glide install
```

### Build

To build the binary for Blacklist Checker run the command below. This will generate a binary
in the bin directory with the name blacklist-checker.

```bash
make build
```

### Install

To install the binary for Blacklist Checker run the command below. This will generate a binary
in $GOPATH/bin directory with the name blacklist-checker, and add the bash autocomplete files.

```bash
make install
```

### Run

### Help
```bash
$ blacklist-checker --help
usage: blacklist-checker [<flags>] <command> [<args> ...]

Fast blacklist checker application

Flags:
  -h, --help                   Show context-sensitive help (also try --help-long and --help-man).
      --verbose                Verbose mode.
      --nameserver=8.8.8.8:53  Name server to use
      --queue=25               How many request to process at one time

Commands:
  help [<command>...]
    Show help.

  version
    Show version and terminate

  ip <ip-address>
    Check IP against available blacklists

  cidr <cidr-address>
    Check CIDR against available blacklists

  list
    List available blacklists
```

#### IP 

```bash
$ time blacklist-checker ip 46.217.104.208
46.217.104.208 blacklisted on b.barracudacentral.org with 127.0.0.2
46.217.104.208 blacklisted on dnsbl-2.uceprotect.net with 127.0.0.2
46.217.104.208 blacklisted on dnsbl-3.uceprotect.net with 127.0.0.2

real	0m0.696s
user	0m0.004s
sys	0m0.004s
```

#### CIDR
```bash
$ time blacklist-checker cidr 46.217.104.208/25
46.217.104.0 blacklisted on dnsbl-2.uceprotect.net with 127.0.0.2
46.217.104.0 blacklisted on dnsbl-3.uceprotect.net with 127.0.0.2
46.217.104.3 blacklisted on dnsbl-1.uceprotect.net with 127.0.0.2
46.217.104.2 blacklisted on b.barracudacentral.org with 127.0.0.2
46.217.104.2 blacklisted on web.dnsbl.sorbs.net with 127.0.0.7
46.217.104.3 blacklisted on db.wpbl.info with 127.0.0.2
46.217.104.1 blacklisted on dnsbl-2.uceprotect.net with 127.0.0.2
...

real	0m24.164s
user	1m5.324s
sys	0m4.324s
```

#### Blacklists

Currently there are 59 blacklists in [blacklists.go](blacklists.go)

```bash
$ blacklist-checker list
access.redhawk.org
b.barracudacentral.org
bl.spamcannibal.org
bl.spamcop.net
blackholes.mail-abuse.org
bogons.cymru.com
cbl.abuseat.org
cbl.anti-spam.org.cn
cdl.anti-spam.org.cn
combined.njabl.org
csi.cloudmark.com
db.wpbl.info
dnsbl-1.uceprotect.net
dnsbl-2.uceprotect.net
dnsbl-3.uceprotect.net
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
xbl.spamhaus.org
zen.spamhaus.org
zombie.dnsbl.sorbs.net
```

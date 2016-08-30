Blacklist Checker
================

Check if your IP is blacklisted or not.

To build/install just pull the code and run **make build** or **make install**, there is also autocomplete files for bash and zsh

```bash
$ blacklist-checker --help
usage: blacklist-checker [<flags>] <command> [<args> ...]

Fast blacklist checker application

Flags:
  -h, --help                   Show context-sensitive help (also try --help-long and --help-man).
      --verbose                Verbose mode.
  -v, --version                Show version and terminate
      --nameserver=8.8.8.8:53  Name server to use
      --queue=25               How many request to process at one time

Commands:
  help [<command>...]
    Show help.

  ip <ip-address>
    Check IP against available blacklists

  cidr <cidr-address>
    Check CIDR against available blacklists

  list
    List available blacklists
```

There are probably faster ways to do this so if anyone want's to cleanup or send a PR feel free to do so

Currently there are 59 blacklists in [blacklists.go](blacklists.go)

### IP 

```bash
$ time blacklist-checker ip 46.217.104.208
46.217.104.208 blacklisted on b.barracudacentral.org with 127.0.0.2
46.217.104.208 blacklisted on dnsbl-2.uceprotect.net with 127.0.0.2
46.217.104.208 blacklisted on dnsbl-3.uceprotect.net with 127.0.0.2

real	0m0.696s
user	0m0.004s
sys	0m0.004s
```

### CIDR
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


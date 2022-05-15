package cmd

import (
	"github.com/ilijamt/blacklist_checker/internal/check"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/sync/semaphore"
	"net/netip"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check available blacklists.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if err = dsnblFileExists(); err != nil {
			return err
		}
		for _, n := range conf.nameservers {
			_, err = netip.ParseAddrPort(n)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	checkCmd.PersistentFlags().StringSliceVar(&conf.nameservers, "nameservers", []string{"1.1.1.1:53"}, "Nameservers to use, a random one is used for each request")
	checkCmd.PersistentFlags().IntVarP(&conf.concurrency, "concurrency", "n", 25, "How many requests to process at once")
	rootCmd.AddCommand(checkCmd)
}

func processIpCheck(sem *semaphore.Weighted, item check.Item) (blacklisted bool, err error) {
	var responses []string
	blacklisted, responses, err = check.Check(
		sem,
		item,
		conf.Nameserver(),
	)
	if blacklisted {
		log.Warn().Str("dnsbl", item.Host).Bool("blacklisted", blacklisted).Strs("responses", responses).IPAddr("ip", item.IP).Send()
	} else {
		log.Trace().Str("dnsbl", item.Host).Bool("blacklisted", blacklisted).Strs("responses", responses).IPAddr("ip", item.IP).Send()
	}
	return blacklisted, err
}

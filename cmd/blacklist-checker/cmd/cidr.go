package cmd

import (
	"context"
	"fmt"
	"github.com/ilijamt/blacklist_checker"
	"github.com/ilijamt/blacklist_checker/internal/check"
	"github.com/ilijamt/blacklist_checker/internal/ip"
	"github.com/ilijamt/blacklist_checker/internal/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/sync/semaphore"
	"net"
	"sync/atomic"
)

var cidrCmd = &cobra.Command{
	Use:   "cidr <cidr-range>",
	Args:  cobra.ExactArgs(1),
	Short: "Check CIDR against available blacklists.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		i, _, err := net.ParseCIDR(args[0])
		if err != nil {
			return err
		}

		if ip.IsPrivate(i) {
			return fmt.Errorf("ip: %s is in the private range", i)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var sem = semaphore.NewWeighted(int64(conf.concurrency))

		var ips []net.IP
		var hosts []string
		var items []check.Item

		ips, err = utils.Hosts(args[0])
		if err != nil {
			return err
		}

		hosts, err = blacklist_checker.GetDNSBLs(conf.dnsbl)
		for _, host := range hosts {
			for _, ip := range ips {
				items = append(items, check.Item{
					IP:        ip,
					Blacklist: fmt.Sprintf("%s.%s.", utils.ReverseIP(ip.String()), host),
					Host:      host,
				})
			}
		}

		log.Info().Int("queries", len(items)).Int("dsnbl", len(hosts)).Msg("processing")
		var blacklisted uint64

		for _, i := range items {
			if err = sem.Acquire(context.Background(), 1); err != nil {
				return err
			}
			go func(item check.Item) {
				if b, _ := processIpCheck(sem, item); b {
					atomic.AddUint64(&blacklisted, 1)
				}
			}(i)
		}

		if err = sem.Acquire(context.Background(), int64(conf.concurrency)); err != nil {
			return err
		}

		log.Info().Uint64("blacklisted", blacklisted).Int("queries", len(items)).Msg("Finished")
		return err
	},
}

func init() {
	checkCmd.AddCommand(cidrCmd)
}

package cmd

import (
	"fmt"
	"github.com/ilijamt/blacklist_checker"
	"github.com/ilijamt/blacklist_checker/internal/utils"
	"github.com/spf13/cobra"
	"math/rand"
	"time"
)

func dsnblFileExists() error {
	if conf.dnsbl != "" {
		exists, err := utils.FileExists(conf.dnsbl)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("file %s does not exist", conf.dnsbl)
		}
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:   blacklist_checker.Name,
	Short: "Check if your IP or CIDR is blacklisted or not.",
	Long:  `A simple tool that helps you check if your IP or CIDR is blacklisted or not.`,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&conf.dnsbl, "dsnbl", "", "DNSBL file to use, if empty it uses the internal list, should be a list of DNSBL to use, each one on a new line")

}

type config struct {
	nameservers []string
	concurrency int
	dnsbl       string
}

func (c config) Nameserver() string {
	if len(c.nameservers) > 1 {
		rand.Seed(time.Now().Unix())
		return c.nameservers[rand.Intn(len(c.nameservers))]
	}
	return c.nameservers[0]
}

var conf config

func Execute() error {
	return rootCmd.Execute()
}

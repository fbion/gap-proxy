package main

import (
	"log"
	"os"

	"github.com/fanpei91/gap-proxy"
	"github.com/spf13/cobra"
)

type flags struct {
	serverAddr string
	localAddr  string
	key        string
}

func main() {
	var f flags
	var root = &cobra.Command{
		Use:           os.Args[0],
		Short:         "gap-proxy is a secure socks5 proxy to speed up your network connection.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			l := gapproxy.NewLocalProxy(f.localAddr, f.serverAddr, f.key)
			err := l.Listen()
			return err
		},
	}
	cobra.EnableCommandSorting = false
	root.Flags().StringVar(&f.serverAddr, "server-addr", "127.0.0.1:2086", "remote server addr")
	root.Flags().StringVar(&f.localAddr, "local-addr", "127.0.0.1:1086", "local server addr")
	root.Flags().StringVar(&f.key, "key", "this is a secret key", "secret key")

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}

package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	// The connection end identifier on the controller chain
	FlagConnectionID = "connection-id"

	// Will be parsed to uint64
	FlagSwapRoutePoolIds = "swap-route-pool-ids"
	// Will be parsed to []sdk.Coin
	FlagSwapRouteAmounts = "swap-route-amounts"
	// Will be parsed to []string
	FlagSwapRouteDenoms = "swap-route-denoms"
)

// common flagsets to add to various functions
var (
	fsConnectionID = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	fsConnectionID.String(FlagConnectionID, "", "Connection ID")
}

func FlagSetQuerySwapRoutes() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.StringArray(FlagSwapRoutePoolIds, []string{""}, "swap route pool id")
	fs.StringArray(FlagSwapRouteDenoms, []string{""}, "swap route amount")
	return fs
}

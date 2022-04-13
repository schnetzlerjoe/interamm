package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/schnetzlerjoe/interamm/x/interamm/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSwapCosmos() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-cosmos [pool-id] [swap-type] [offer-coin] [demand-coin-denom] [order-price] [swap-fee-rate]",
		Short: "Broadcast message SwapCosmos",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPoolId := args[0]
			argSwapType := args[1]
			argOfferCoin := args[2]
			argDemandCoinDenom := args[3]
			argOrderPrice := args[4]
			argSwapFeeRate := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwapCosmos(
				clientCtx.GetFromAddress().String(),
				argPoolId,
				argSwapType,
				argOfferCoin,
				argDemandCoinDenom,
				argOrderPrice,
				argSwapFeeRate,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

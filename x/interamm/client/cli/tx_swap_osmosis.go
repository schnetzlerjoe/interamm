package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	intertxtypes "github.com/cosmos/interchain-accounts/x/inter-tx/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var _ = strconv.Itoa(0)

func CmdSwapOsmosis() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-osmosis [token-in] [token-out-min-amount]",
		Short: "Broadcast message swap-osmosis",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

			var json string = `{
				"@type":"/osmosis.gamm.v1beta1.MsgSwapExactAmountIn",
				"sender":"cosmos15ccshhmp0gsx29qpqq6g4zmltnnvgmyu9ueuadh9y2nc5zj0szls5gtddz",
				"tokenIn": {
					denom: uosmo
				},
				"routes": {
					"poolId": "1",
					"tokenOutDenom": "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
				}
			}`

			var txMsg sdk.Msg
			if err := cdc.UnmarshalInterfaceJSON([]byte(json), &txMsg); err != nil {

				if err != nil {
					return errors.Wrap(err, "JSON input not provided or created incorrectly")
				}
			}

			msg, err := intertxtypes.NewMsgSubmitTx(txMsg, viper.GetString(FlagConnectionID), clientCtx.GetFromAddress().String())
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

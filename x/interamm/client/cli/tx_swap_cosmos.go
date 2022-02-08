package cli

import (
<<<<<<< HEAD
=======
	"io/ioutil"
>>>>>>> 652751544ee2e49bedea2abc3a4c3cbe1672d64f
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

func CmdSwapCosmos() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-cosmos [pool-id] [swap-type] [offer-coin] [demand-coin-denom] [order-price] [swap-fee-rate]",
		Short: "Broadcast message swap-cosmos",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

<<<<<<< HEAD
			var json string = `{
				"@type":"/cosmos.liquidity.v1beta1.MsgSwapWithinBatch",
				"sender":"cosmos15ccshhmp0gsx29qpqq6g4zmltnnvgmyu9ueuadh9y2nc5zj0szls5gtddz",
				"tokenIn": {
					denom: uosmo
				},
				"routes": {
					"poolId": "stake",
					"tokenOutDenom": "1000"
				}
			}`

			var txMsg sdk.Msg
			if err := cdc.UnmarshalInterfaceJSON([]byte(json), &txMsg); err != nil {

				if err != nil {
					return errors.Wrap(err, "JSON input not provided or created incorrectly")
=======
			var txMsg sdk.Msg
			if err := cdc.UnmarshalInterfaceJSON([]byte(args[0]), &txMsg); err != nil {

				// check for file path if JSON input is not provided
				contents, err := ioutil.ReadFile(args[0])
				if err != nil {
					return errors.Wrap(err, "neither JSON input nor path to .json file for sdk msg were provided")
				}

				if err := cdc.UnmarshalInterfaceJSON(contents, &txMsg); err != nil {
					return errors.Wrap(err, "error unmarshalling sdk msg file")
>>>>>>> 652751544ee2e49bedea2abc3a4c3cbe1672d64f
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

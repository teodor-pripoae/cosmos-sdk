package rpc

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func statusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Query remote node for status",
		RunE:  printNodeStatus,
	}
	cmd.Flags().StringP(client.FlagNode, "n", "tcp://localhost:46657", "Node to connect to")
	return cmd
}

func getNodeStatus() (*ctypes.ResultStatus, error) {
	// get the node
	node, err := context.NewCoreContextFromViper().GetNode()
	if err != nil {
		return &ctypes.ResultStatus{}, err
	}
	return node.Status()
}

// CMD

func printNodeStatus(cmd *cobra.Command, args []string) error {
	status, err := getNodeStatus()
	if err != nil {
		return err
	}

	output, err := cdc.MarshalJSON(status)
	// output, err := cdc.MarshalJSONIndent(res, "  ", "")
	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}

// REST

func NodeInfoRequestHandler(w http.ResponseWriter, r *http.Request) {
	status, err := getNodeStatus()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	nodeInfo := status.NodeInfo
	output, err := cdc.MarshalJSON(nodeInfo)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(output)
}

func NodeSyncingRequestHandler(w http.ResponseWriter, r *http.Request) {
	status, err := getNodeStatus()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	syncing := status.SyncInfo.Syncing
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(strconv.FormatBool(syncing)))
}

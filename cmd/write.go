package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
	"github.com/spf13/cobra"
)

var vaultName string = os.Getenv("OP_CONNECT_VAULT")
var singleline bool

var writeCmd = &cobra.Command{
	Use:               "write ITEM_NAME FIELD_NAME [VALUE|-]",
	Short:             "write secret to 1password connect server",
	Long:              `write secret to 1password connect server`,
	DisableAutoGenTag: true,
	Args:              cobra.MatchAll(cobra.ExactArgs(3), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		writeSecret(createClient(), args)
	},
}

func init() {
	writeCmd.Flags().BoolVarP(&singleline, "singleline", "s", false, "Insert single line parameter (end with \\n)")
	rootCmd.AddCommand(writeCmd)
}

func writeSecret(client connect.Client, params []string) {
	var item_name string = params[0]
	var field_name string = params[1]
	var key_value string = params[2]
	var vaultId string = getVaultId(client)

	if key_value == "-" {
		key_value = fetchPipe()
	}

	// check if item already exists?
	item, _ := client.GetItemByTitle(item_name, vaultId)

	//if item already exists, this is an update
	if item != nil {
		var field_exists = false
		for _, field := range item.Fields {
			if field.Label == field_name {
				// the field exists, update it
				field.Value = key_value
				field_exists = true
			}
		}
		if !field_exists {
			// this is a new field, add it
			item.Fields = append(item.Fields, &onepassword.ItemField{Label: field_name, Value: key_value, Type: "CONCEALED"})
		}
		updatedItem, err := client.UpdateItem(item, vaultId)
		exitOnError(err)
		fmt.Println("opw: updated " + updatedItem.Title)

		// if item does not exist then create it
	} else {
		newItem := &onepassword.Item{
			Category: onepassword.ApiCredential,
			Title:    item_name,
			Fields: []*onepassword.ItemField{{
				Label: field_name,
				Value: key_value,
				Type:  "CONCEALED",
			}},
		}
		result, err := client.CreateItem(newItem, vaultName)
		exitOnError(err)
		fmt.Println("opw: created " + result.Title)
	}
}

func fetchPipe() string {
	var key_value string
	// Read value from standard input
	if singleline {
		buf := bufio.NewReader(os.Stdin)
		v, err := buf.ReadString('\n')
		if err != nil {
			exitOnError(err)
		}
		key_value = strings.TrimSuffix(v, "\n")
	} else {
		v, err := io.ReadAll(os.Stdin)
		if err != nil {
			exitOnError(err)
		}
		key_value = string(v)
	}
	return key_value
}

func createClient() connect.Client {
	client, err := connect.NewClientFromEnvironment()
	exitOnError(err)
	return client
}

func getVaultId(client connect.Client) string {
	vault, err := client.GetVault(vaultName)
	exitOnError(err)
	return vault.ID
}

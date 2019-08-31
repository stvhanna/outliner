package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	ol "github.com/jyny/outliner/pkg/outliner"
	"github.com/jyny/outliner/pkg/util"
)

func init() {
	createCmd.Flags().StringP("spec", "s", "", "Spec of instance (required)")
	createCmd.Flags().StringP("region", "r", "", "region of instance (required)")
	createCmd.Flags().StringP("provider", "p", "", "Provider of instance (required)")
	createCmd.MarkFlagRequired("spec")
	createCmd.MarkFlagRequired("region")
	createCmd.MarkFlagRequired("provider")
	viper.BindPFlag("spec", createCmd.Flags().Lookup("spec"))
	viper.BindPFlag("region", createCmd.Flags().Lookup("region"))
	viper.BindPFlag("provider", createCmd.Flags().Lookup("provider"))
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a server",
	Long:  `create a server`,
	Run: func(cmd *cobra.Command, args []string) {
		util.PrintCreateInstanceStart()
		outliner.CreateInstance(ol.Instance{
			Provider: viper.GetString("provider"),
			Region: ol.Region{
				ID: viper.GetString("region"),
			},
			Spec: ol.Spec{
				ID: viper.GetString("spec"),
			},
		})
		util.PrintCreateInstanceDone()
	},
}

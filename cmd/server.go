package cmd

import (
	"fmt"
	"github.com/uanid/fakenews-server/application/configs"

	"github.com/spf13/cobra"
	"github.com/uanid/fakenews-server/application/fnc_server"
)

func init() {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Fake News Challenge Api Server를 실행합니다",
	}

	var port int
	serverCmd.Flags().IntVar(&port, "port", 8080, "Rest Api Server Port")

	serverCmd.RunE = func(cmd *cobra.Command, args []string) error {
		cfg, err := configs.LoadConfig(configPath)
		if err != nil {
			return err
		}

		app, err := fnc_server.NewApplication(port, cfg)
		if err != nil {
			return fmt.Errorf("ApplicationInitFailed: %s", err.Error())
		}

		err = app.Start()
		if err != nil {
			return err
		}
		fmt.Println("Stop Application.")
		return nil
	}

	rootCmd.AddCommand(serverCmd)
}

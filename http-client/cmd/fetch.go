package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A CLI tool to do HTTP requests",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var GetCmd = &cobra.Command{
	Use:   "get URL",
	Short: "perform a http get request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		u := args[0]

		url, err := url.Parse(u)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(cmd.Context(), 5*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
		if err != nil {
			return err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(resBody))

		return nil
	},
}

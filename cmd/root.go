package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yuriizinets/go-nominatim"
)

var (
	cfgFile        string

	rootCmd = &cobra.Command{
		Use:   "oslatlong",
		Short: "OSM Nominatum-backed Geocoder",
		Run: func(cmd *cobra.Command, args []string) {
			queries := args
			if len(queries) == 0 {
				cmd.Help()
			}

			maxResults := viper.GetInt("max-results")
			showQuery := viper.GetBool("show-query")

			for _, q := range queries {
				n := nominatim.Nominatim{}

				results, err := n.Search(nominatim.SearchParameters{
					Query: q,
				})

				if err != nil {
					log.Panic().Err(err).Msg("could not perform search")
				}

				if maxResults > 0 && len(results) > maxResults {
					results = results[0:maxResults]
				}

				for _, result := range results {
					line := ""
					if showQuery {
						line += fmt.Sprintf("%s | ", q)
					}

					line += fmt.Sprintf("(%s, %s)", result.LatStr, result.LngStr)
					fmt.Println(line)
				}
			}

		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolP("debug", "", false, "Enable debug logging")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindEnv("debug", "DEBUG")

	rootCmd.PersistentFlags().IntP("max-results", "n", 0, "Maximum number of results to return")
	viper.BindPFlag("max-results", rootCmd.PersistentFlags().Lookup("max-results"))

	rootCmd.PersistentFlags().BoolP("show-query", "s", false, "Prefix results with query")
	viper.BindPFlag("show-query", rootCmd.PersistentFlags().Lookup("show-query"))
}

func initConfig() {
	logLevel := zerolog.InfoLevel
	if viper.GetBool("debug") {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

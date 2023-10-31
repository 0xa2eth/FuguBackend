package main

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

// Copyright © 2023 FuGu Toxic <penghan063@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
const _UI = `

███████╗██╗   ██╗ ██████╗ ██╗   ██╗    ████████╗ ██████╗ ██╗  ██╗██╗ ██████╗
██╔════╝██║   ██║██╔════╝ ██║   ██║    ╚══██╔══╝██╔═══██╗╚██╗██╔╝██║██╔════╝
█████╗  ██║   ██║██║  ███╗██║   ██║       ██║   ██║   ██║ ╚███╔╝ ██║██║     
██╔══╝  ██║   ██║██║   ██║██║   ██║       ██║   ██║   ██║ ██╔██╗ ██║██║     
██║     ╚██████╔╝╚██████╔╝╚██████╔╝       ██║   ╚██████╔╝██╔╝ ██╗██║╚██████╗
╚═╝      ╚═════╝  ╚═════╝  ╚═════╝        ╚═╝    ╚═════╝ ╚═╝  ╚═╝╚═╝ ╚═════╝

`

func main() {
	Execute()
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gate3",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		serverCtx := dao.GetServiceCtx()
		if serverCtx == nil {
			panic("err config")
		}
		config.Logger.Info("[INFO]", zap.Any("====>>>>", serverCtx.C.ServerC))
		r := router.NewRouter(serverCtx)
		appRouter, err := app.NewPlatform(serverCtx.C, r, serverCtx)
		if err != nil {
			config.Logger.Error("init error", zap.Error(err))
			return
		}
		if err := blockchain.EthClientInit(); err != nil {
			config.Logger.Error("eth_client init failed", zap.Error(err))
			return
		}
		if err := utils.SonyFlakeInit(config.Conf.Common.StartTime, 1); err != nil {
			config.Logger.Error("SonyFlakeInit init failed", zap.Error(err))
			return
		}
		//go func() {
		//	fmt.Println("==============start\n")
		//	gate3.DrawTime(context.Background())
		//}()
		//执行定时任务
		gate3.Start()

		appRouter.AppStart()

		chSig := make(chan os.Signal)
		signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
		<-chSig

		appRouter.AppClose()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gate3.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gate3" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gate3")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

type Platform struct {
	config    *config.Config
	router    *gin.Engine
	serverCtx *svc.ServiceCtx
}

func NewPlatform(config *config.Config, router *gin.Engine, server *svc.ServiceCtx) (*Platform, error) {
	return &Platform{
		config:    config,
		router:    router,
		serverCtx: server,
	}, nil
}

func (p *Platform) AppStart() error {
	config.Logger.Info("[INFO]", zap.String("service is", " starting ..."), zap.Any("address:", p.config.ServerC.Addr))
	if err := p.router.Run(p.config.ServerC.Addr); err != nil {
		config.Logger.Error("start service faild", zap.Error(err))
		return err
	}
	return nil
}

func (p *Platform) AppClose() error {
	p.serverCtx.Rds.Close()
	db, _ := p.serverCtx.Db.DB()
	db.Close()
	return nil
}

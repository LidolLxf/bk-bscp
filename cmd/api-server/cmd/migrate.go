/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package cmd provides operations for init and upgrading the bk-apigateway resources
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/TencentBlueKing/bk-bscp/internal/iam/apigw"
	"github.com/TencentBlueKing/bk-bscp/pkg/cc"
	"github.com/TencentBlueKing/bk-bscp/pkg/logs"
	"github.com/TencentBlueKing/bk-bscp/pkg/version"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "api-server migrations tool",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var migrateInitApigatewayCmd = &cobra.Command{
	Use:   "init-apigateway",
	Short: "Create 'bk-bscp' apigateway instance and upgrade api resources",
	Run: func(cmd *cobra.Command, args []string) {

		if err := cc.LoadSettings(SysOpt.Sys); err != nil {
			fmt.Println("load settings from config files failed, err:", err)
			return
		}

		logs.InitLogger(cc.ApiServer().Log.Logs())

		if err := apigw.ReleaseSwagger(cc.ApiServer().Esb, cc.ApiServer().ApiGateway, "zh",
			fmt.Sprintf("%s+%s", version.GITTAG, time.Now().Format("20060102150405"))); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {

	// Add "--debug" flag to all migrate sub commands
	migrateCmd.PersistentFlags().BoolP("debug", "d", false,
		"whether to debug output the execution process,, default is false")

	migrateCmd.AddCommand(migrateInitApigatewayCmd)

	// Add "migrate" command to the root command
	rootCmd.AddCommand(migrateCmd)
}

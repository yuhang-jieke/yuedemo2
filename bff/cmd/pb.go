/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// pbCmd represents the pb command
var pbCmd = &cobra.Command{
	Use:   "pb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		protocCmd := exec.Command(
			"protoc",
			"-I", ".",
			"--go_out=./"+Address,
			"--go-grpc_out=./"+File,
			"./goods.proto",
		)

		// 2. 打印完整命令，方便你手动在终端测试
		fmt.Println("执行的命令:", protocCmd.String())

		// 3. 捕获并打印标准输出和标准错误
		output, err := protocCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("命令执行失败: %v\n", err)
			fmt.Printf("错误输出: %s\n", string(output))
			return
		}

		// 4. 如果成功，也打印输出
		fmt.Printf("命令执行成功，输出: %s\n", string(output))
		fmt.Println("pb called")
	},
}
var (
	Address string
	File    string
)

// sql2pb  -field_style "sel2pb" -go_package " " -host 115.190.57.118 -package "sel2pb"  -password 4ay1nkal3u8ed77y  -port 3306 -schema yuedemo  -service_name "sel2pb"  -user root  -table goods > goods.proto
func init() {
	rootCmd.AddCommand(pbCmd)
	pbCmd.Flags().StringVarP(&Address, "add", "a", "", "")
	pbCmd.Flags().StringVarP(&File, "file", "f", "", "")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

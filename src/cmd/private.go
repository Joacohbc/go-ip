/*
Copyright © 2022 Joacohbc <joacog48@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"go-ip/src/simpleInterface"
	"net"
	"strings"

	"github.com/spf13/cobra"
)

// privateCmd represents the private command
var privateCmd = &cobra.Command{
	Use:   "private",
	Short: "Muestra el IP publica de la interfaces de red el dispositvo",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		//Parse all Flags
		all, err := cmd.Flags().GetBool("all")
		cobra.CheckErr(err)

		intefToShow, err := cmd.Flags().GetString("interface")
		cobra.CheckErr(err)

		ipv6, err := cmd.Flags().GetBool("ipv6")
		cobra.CheckErr(err)

		//Convierto de net.Interface a mi propio tipo simplificado
		var intfs []simpleInterface.SimpleInterface
		{
			interfaces, err := net.Interfaces()
			if err != nil {
				cobra.CheckErr(fmt.Errorf("no se pudieron obtener las interfaces: %s", err.Error()))
			}

			for _, intf := range interfaces {
				sintf, err := simpleInterface.NewSimpleInterface(intf)
				cobra.CheckErr(err)
				intfs = append(intfs, sintf)
			}
		}

		if strings.TrimSpace(intefToShow) != "" {
			for _, intf := range intfs {
				if intefToShow == intf.Name {

					if all {
						fmt.Println(intf.String())
						return
					}

					if ipv6 {
						fmt.Println(intf.IPv6)
						return
					}

					fmt.Println(intf.IPv4)
					return
				}
			}

			cobra.CheckErr(fmt.Errorf("no existe la interfaz \"%s\"", intefToShow))
		}

		for _, intf := range intfs {
			fmt.Println(intf.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(privateCmd)
	privateCmd.Flags().StringP("interface", "i", "", "La interfaz de red de la cual se quiere saber la red")
	privateCmd.Flags().BoolP("ipv6", "6", false, "Muestra el IPv6 de la placa ingresda con -i (de manera predeterminda simpre sera el IPv4)")
	privateCmd.Flags().BoolP("mac", "m", false, "Muestra la mac")
	privateCmd.Flags().BoolP("all", "a", false, "Lista toda las interfaz de red y sus IPs y MACs")
}

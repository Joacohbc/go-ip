/*
Copyright © 2022 Joaquin Genova <joaquingenovag8@gmail.com>

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
	"net"
	"strings"

	"github.com/spf13/cobra"
)

type SimpleInterface struct {
	Name string
	IPv4 net.IP
	IPv6 net.IP
	Mac  net.HardwareAddr
}

func (s *SimpleInterface) String() string {
	return fmt.Sprintf(`
%s: 
	IPv4: %s 
	IPv6: %s 
	MAC: %s`, s.Name, s.IPv4, s.IPv6, s.Mac)
}

// privateCmd represents the private command
var privateCmd = &cobra.Command{
	Use:   "private",
	Short: "Muestra el IP publica de la interfaces de red el dispositvo",
	Run: func(cmd *cobra.Command, args []string) {

		//Parse all Flags
		all, err := cmd.Flags().GetBool("all")
		cobra.CheckErr(err)

		intefToShow, err := cmd.Flags().GetString("interfaces")
		cobra.CheckErr(err)

		ipv6, err := cmd.Flags().GetBool("ipv6")
		cobra.CheckErr(err)

		//Conver type net.Interface to my own type
		var intfs []SimpleInterface
		{
			interfaces, err := net.Interfaces()
			if err != nil {
				cobra.CheckErr(fmt.Errorf("no se pudieron obtener las interfaces: %s", err.Error()))
			}

			for _, intf := range interfaces {
				addrs, err := intf.Addrs()
				if err != nil {
					cobra.CheckErr(fmt.Errorf("error al obtener las direcciones de esta interfaz: %s", err.Error()))
				}

				//Si tiene addrs, sino esta "apagada" o sin conexion
				if len(addrs) >= 2 {
					intfs = append(intfs, SimpleInterface{
						Name: intf.Name,
						IPv4: net.ParseIP(strings.Split(addrs[0].String(), "/")[0]),
						IPv6: net.ParseIP(strings.Split(addrs[1].String(), "/")[0]),
						Mac:  intf.HardwareAddr,
					})
				}
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
	privateCmd.Flags().StringP("interfaces", "i", "", "La interfaz de red de la cual se quiere saber la red")
	privateCmd.Flags().BoolP("ipv6", "6", false, "Muestra el IPv6 de la placa (de manera predeterminda simpre sera el IPv4)")
	privateCmd.Flags().BoolP("mac", "m", false, "Muestra la mac")
	privateCmd.Flags().BoolP("all", "a", false, "Lista toda las interfaz de red y sus IPs y MACs")
}

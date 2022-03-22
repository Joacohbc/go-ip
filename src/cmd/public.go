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
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var IPUrls []string = []string{
	"http://checkip.amazonaws.com",
	"http://ifconfig.me/ip",
	"http://ifconfig.co",
	"http://ip-api.com/line/?fields=query",
	"http://ipwhois.app/csv/?objects=ip",
	"https://ipapi.co/ip",
}

// publicCmd represents the public command
var publicCmd = &cobra.Command{
	Use:   "public",
	Short: "Imprime la IP publica",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		//Nunca imprimo los errores debido a que tiene que pobrar todos URL
		for _, url := range IPUrls {
			req, err := http.Get(url)
			if err != nil {
				//cobra.CheckErr(fmt.Errorf("error al hacer la peticion GET a \"%s\": %s", url, err.Error()))
				break
			}

			if req.StatusCode == http.StatusOK {
				body, err := ioutil.ReadAll(req.Body)
				if err != nil {
					//cobra.CheckErr(fmt.Errorf("error al leer el cuerpo de la peticion: %s", err.Error()))
					break
				}
				fmt.Print(strings.TrimSpace(string(body)))
				return

			} else {
				//cobra.CheckErr(fmt.Errorf("error al hacer la peticion GET a \"%s\": %s", url, err.Error()))
				break
			}
		}

		cobra.CheckErr(fmt.Errorf("ninguno de los servicios de consultoria de IP contestaron"))
	},
}

func init() {
	rootCmd.AddCommand(publicCmd)
}

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
package simpleInterface

import (
	"fmt"
	"net"
	"strings"
)

type SimpleInterface struct {
	Name string
	IPv4 net.IP
	IPv6 net.IP
	Mac  net.HardwareAddr
	Flag net.Flags
}

func (s *SimpleInterface) String() string {
	return fmt.Sprintf(`
|%s: 
|	IPv4: %s 
|	IPv6: %s 
|	MAC: %s
|	Flags: %s`, s.Name, s.IPv4, s.IPv6, s.Mac, s.Flag)
}

func NewSimpleInterface(intf net.Interface) (SimpleInterface, error) {

	addrs, err := intf.Addrs()
	if err != nil {
		return SimpleInterface{}, fmt.Errorf("error al obtener las direcciones de esta interfaz: %s", err.Error())
	}

	//Si tiene addrs, sino esta "apagada" o sin conexion
	if len(addrs) >= 2 {
		return SimpleInterface{
			Name: intf.Name,
			IPv4: net.ParseIP(strings.Split(addrs[0].String(), "/")[0]),
			IPv6: net.ParseIP(strings.Split(addrs[1].String(), "/")[0]),
			Mac:  intf.HardwareAddr,
			Flag: intf.Flags,
		}, nil
	} else if len(addrs) == 1 {
		return SimpleInterface{
			Name: intf.Name,
			IPv4: net.ParseIP(strings.Split(addrs[0].String(), "/")[0]),
			Mac:  intf.HardwareAddr,
			Flag: intf.Flags,
		}, nil
	} else {
		return SimpleInterface{
			Name: intf.Name,
			Mac:  intf.HardwareAddr,
			Flag: intf.Flags,
		}, nil
	}

}

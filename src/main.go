package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"io"
	"os"
	"strings"
)

// ParseConfigBuffer reads config from a generic Reader.
func ParseConfigBuffer(in io.Reader, format string) (string, bool) {

	viper.SetConfigType(format)
	err := viper.ReadConfig(in)
	if err == nil {
		var factoryState *FactoryState
		err = viper.Unmarshal(&factoryState)
		if &factoryState.Dev != nil && factoryState.Dev.DefaultFactoryAddress != "" {
			return factoryState.Dev.DefaultFactoryAddress, true
		}
	}
	return "", false
}

// ParseConfigFile reads the file using Viper.
func ParseConfigFile(configFileName string) (string, bool) {
	components := strings.Split(configFileName, ".")
	format := components[len(components)-1]

	file, err := os.Open(configFileName)
	if err == nil {
		configBuffer := make([]byte, 10240)
		sz, err := file.Read(configBuffer)
		if err == nil {
			return ParseConfigBuffer(bytes.NewBuffer(configBuffer[:sz]), format)
		}
	}
	return "", false
}

func main() {
	factoryAddress, isFound := ParseConfigFile("factoryState.toml")
	if !isFound {
		fmt.Println("dev default factory address not found")
	}
	fmt.Println("Factory Address :", factoryAddress)

}

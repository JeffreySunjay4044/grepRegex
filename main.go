package main

import (
	viper "github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	check := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	viper.SetConfigName("factoryState.toml") // name of config file (without extension)
	viper.SetConfigType("toml")              // REQUIRED if the config file does not have the extension in the name
	myConfigPath := "../../../scripts/generated/"
	fh, err := os.OpenFile(myConfigPath, os.O_RDONLY, 0666)
	check(err)
	viper.SetConfigType("toml") // do not ignore
	err = viper.ReadConfig(fh)
	check(err)

	// Read
	log.Printf("%#v", viper.GetString("title"))                 // "my config"
	log.Printf("%#v", viper.GetString("DataTitle.12345.prop1")) // "30"
	log.Printf("%#v", viper.GetString("dataTitle.12345.prop1")) // "30"  // case-insensitive
	log.Printf("%#v", viper.GetInt("DataTitle.12345.prop1"))    // 30
	log.Printf("%#v", viper.GetIntSlice("feature1.userids"))    // []int{456, 789}

	// Write
	viper.Set("database", "newuser")
	viper.Set("owner.name", "Carson")
	viper.Set("feature1.userids", []int{111, 222}) // overwrite
	err = viper.WriteConfigAs(myConfigPath)
	check(err)
}

//
//viper.SetConfigName("factoryState.toml") // name of config file (without extension)
//viper.SetConfigType("toml") // REQUIRED if the config file does not have the extension in the name
//viper.AddConfigPath("../../../scripts/generated/")   // path to look for the config file in
//err := viper.ReadInConfig()
//if err != nil {
//	panic(fmt.Errorf("fatal error config file: %w", err))
//}

//viper.

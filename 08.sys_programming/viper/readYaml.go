package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	var configFile *string = flag.String("c", "myConfig", "Setting the configuration file")
	flag.Parse()

	// os.Stat을 호출해 설정 플래그인 --c의 값을 검사
	_, err := os.Stat(*configFile)
	if err == nil {
		fmt.Println("Using User Specified Configuration File!")
		viper.SetConfigFile(*configFile)
	} else {
		viper.SetConfigName(*configFile)
		viper.AddConfigPath("/tmp")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	if viper.IsSet("item1.k1") {
		fmt.Println("item1.val1:", viper.Get("item1.k1"))
	} else {
		fmt.Println("item.k1 not set!")
	}

	if viper.IsSet("item2.k2") {
		fmt.Println("item2.val2:", viper.Get("item2.k2"))
	} else {
		fmt.Println("item1.k2 not set!")
	}

	if !viper.IsSet("item3.k1") {
		fmt.Println("item3.k1 is not set!")
	}
}

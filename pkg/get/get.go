package get

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetImageName(cfg string) {
	fmt.Printf("%v\n", cfg)
	viper.SetConfigFile(cfg)
	viper.ReadInConfig()
	a := viper.Get("repo.gcr.namespaces")
	fmt.Println(a)

}

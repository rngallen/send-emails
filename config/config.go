package config

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

// Return string of variable environment
func Config(key string) string {

	// Config file name without extension
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml") //optional

	// Enable viper to read variables
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Config file error %v \n", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid variable type assertion for %v \n", key)
	}
	return value

}

// Return Int of variable value
func Intconfig(key string) int {

	var err error
	v := Config(key)
	value, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("%v could not converted to Integer", v)
	}
	return value
}

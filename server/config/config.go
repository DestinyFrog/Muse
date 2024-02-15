package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Port				string	`mapstructure:"SERVER_PORT"`
	Database_Url		string	`mapstructure:"DATABASE_URL"`
	Data_Files_Folder	string	`mapstructure:"DATA_FILES_FOLDER"`
}

var Config config

func init() {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath(".")

	// Tell viper the name of your file
	viper.SetConfigName(".env")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatal(err)
	}
}
package util

import (
	"encoding/json"
	"fmt"
	"go_gin_app/models"
	"io/ioutil"
	"os"
)

func AddDataToDb() {
	jsonFile, err := os.Open("db_data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened db_data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	type Data struct {
		Users []models.User `json:"users"`
	}

	// we initialize our Users array
	var data Data

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &data)
	for _, user := range data.Users {
		if _, user_db := models.GetUserByUsername(&user.Username); user_db == nil {
			models.AddUser(&user)
		}
	}

	fmt.Print(data)
	//runtime_viper := viper.New()
	//runtime_viper.SetConfigName("db_data")
	//runtime_viper.AddConfigPath(".")
	//err := runtime_viper.ReadInConfig() // Find and read the config file
	//if err != nil { // Handle errors reading the config file
	//	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	//}
	//
	//if users, ok := runtime_viper.Get("users").([]map[string]interface{}); ok{
	//	for _, user := range users {
	//		user = user.(models.User)
	//	}
	//}
}

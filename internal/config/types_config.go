package config

const configFileName = "/.gatorconfig.json"

type Config struct {
	DB_URL			  string  `json:"db_url"`
	Current_User_Name *string `json:"current_user_name"`
}
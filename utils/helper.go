package utils

import (
	"time"
)

//current time timestamp
func Timestamp ()string{
	currentTime := time.Now()
	var timestamp string = currentTime.Format("01-02-2006 15:04:05")
	return timestamp
}

// Get config path for local or docker
func GetConfigPath(path string ) string  {
	if path == "docker"{
		return "./config/config-docker"
	}
	 return "./config/config-local"
}

//api logger

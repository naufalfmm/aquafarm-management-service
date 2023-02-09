package main

import "github.com/naufalfmm/aquafarm-management-service/app"

//	@title			Aquafarm Management Service
//	@version		1.0
//	@description	This is prototype of aquafarm management service

//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						Authorization
func main() {
	app.Init().Run()
}

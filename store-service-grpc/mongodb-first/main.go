package main

import "mongodb-first/config/register"

func main() {
	register.EtcdServerRegister()
	register.GrpcRegister()
}

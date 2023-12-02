package main

import "mysql-second/config/register"

func main() {
	register.EtcdServerRegister()
	register.GrpcRegister()
}

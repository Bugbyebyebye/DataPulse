package main

import (
	"mysql-first/config/register"
)

func main() {
	register.EtcdServerRegister()
	register.GrpcRegister()
}

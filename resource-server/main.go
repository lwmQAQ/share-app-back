package main

import (
	v1 "resource-server/cmd/v1"
	"resource-server/internal/svc"
)

func main() {
	svc := svc.NewServerContext()
	v1.ServerStart(svc)
}

// +build ignore

package main

import (
	"github.com/m3hm3t/account-client-api/internal/pkg/account/adapter/account/delete"
	"log"
)

func main() {

	deleterAdapter := delete.ProvideAccountDeleter()

	deleteResponseStatus, err := deleterAdapter.DeleteAccount("a", "1")
	if err != nil {
		log.Println("ERROR: ", err)
	} else {
		log.Println("account delete response: ", deleteResponseStatus)
	}
}

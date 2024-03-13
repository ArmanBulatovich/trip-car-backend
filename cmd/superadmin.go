package cmd

import (
	"fmt"
	"os"

	"github.com/trip/trip-service/internal/services/admin_services"
)

var CreateSuperAdminFlag = "createsuperadmin"

func CreateSuperAdmin() {
	var email, password string

	fmt.Print("Email: ")
	fmt.Fscan(os.Stdin, &email)

	fmt.Print("Password: ")
	fmt.Fscan(os.Stdin, &password)

	err := admin_services.CreateSuperAdmin(email, password)
	if err != nil {
		fmt.Println("!ERROR!")
		fmt.Println(err)
	} else {
		fmt.Println("!OK!")
	}
}

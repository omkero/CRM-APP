package main

import "fmt"

const (
	SALES_MANAGER          = "SALES_MANAGER"
	CUSTOMER_SUPPORT_AGENT = "CUSTOMER_SUPPORT_AGENT"
	PRODUCT_MANAGER        = "PRODUCT_MANAGER"
	ADMIN                  = "ADMIN"
)

var actions = map[string][]string{
	"removeUser": {SALES_MANAGER, CUSTOMER_SUPPORT_AGENT},
	"addUser":    {PRODUCT_MANAGER, ADMIN},
}

func main() {
	for _, item := range actions["removeUser"] {
		fmt.Println(item)
	}
	fmt.Println(actions["removeUser"])
}

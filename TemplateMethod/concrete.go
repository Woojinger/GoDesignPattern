package main

import "fmt"

type ConcreteConnectHelper struct {
	chelper AbstConnectHelper
}

func (c *ConcreteConnectHelper) requestConnection(info string) string {
	var id, password, userName, decodeInfo, userInfo string

	// decode encrypted info
	decodeInfo = c.chelper.doSecurity(info)

	// get id and password from decodeInfo
	id = "idfrominfo"
	password = "pwfrominfo"

	userInfo = c.chelper.authentication(id, password)
	userName = "userName"

	fmt.Println(decodeInfo, userInfo)
	result := c.chelper.authorization(userName)

	switch result {
	case 0:
		fmt.Println("free member")
	case 1:
		fmt.Println("paid member")
	case 2:
		fmt.Println("game master")
	default:
		fmt.Println("no authorization")
	}

	return c.chelper.connection(userInfo)

}

type ConcreteImplementation struct {
	ConcreteConnectHelper
}

func (c *ConcreteImplementation) doSecurity(info string) string {
	return info
}

func (c *ConcreteImplementation) authentication(id string, password string) string {
	if id == "idfrominfo" && password == "pwfrominfo" {
		return "true info"
	} else {
		return "false info"
	}
}

func (c *ConcreteImplementation) authorization(userName string) int {
	return 0
}

func (c *ConcreteImplementation) connection(info string) string {
	return info
}

func (c *ConcreteImplementation) requestConnection(info string) string {
	return c.ConcreteConnectHelper.requestConnection(info)
}

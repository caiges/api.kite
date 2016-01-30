package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/koding/kite"
	"github.com/koding/kite/config"
	"github.com/koding/kite/protocol"
)

func main() {

	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/foo", foo)
	r.GET("/", hello)

	r.Run(":3000")
}

func hello(c *gin.Context) {
	c.String(200, "hello world")
}

func ping(c *gin.Context) {
	k := kite.New("api", "1.0.0")

	k.Config = config.MustGet()
	k.Config.Environment = "development"

	k.SetLogLevel(kite.DEBUG)
	fmt.Println(k.Config)

	// search a kite that has the same username and environment as us, but the
	kites, gkerr := k.GetKites(&protocol.KontrolQuery{
		Username:    k.Config.Username,
		Environment: k.Config.Environment,
		Name:        "pingpong",
	})
	if gkerr != nil {
		fmt.Println("no kites found named pingpong")
		k.Log.Fatal(gkerr.Error())
	}

	// there might be several kites that matches our query
	client := kites[0]

	fmt.Println("found a kite")
	fmt.Println(client.URL)

	connected, err := client.DialForever()
	if err != nil {
		k.Log.Fatal(err.Error())
	}

	// Wait until connected
	<-connected

	// Call a method of client kite
	response, err := client.TellWithTimeout("ping", 4*time.Second)
	if err != nil {
		k.Log.Error(err.Error())
	}

	// Print the result
	res, _ := response.String()

	c.String(200, "ping >< "+res)
}

func foo(c *gin.Context) {
	k := kite.New("api", "1.0.0")

	k.Config = config.MustGet()
	k.Config.Environment = "development"

	k.SetLogLevel(kite.DEBUG)
	fmt.Println(k.Config)

	// search a kite that has the same username and environment as us, but the
	kites, gkerr := k.GetKites(&protocol.KontrolQuery{
		Username:    k.Config.Username,
		Environment: k.Config.Environment,
		Name:        "foobar",
	})
	if gkerr != nil {
		fmt.Println("no kites found named foobar")
		k.Log.Fatal(gkerr.Error())
	}

	// there might be several kites that matches our query
	client := kites[0]

	fmt.Println("found a kite")
	fmt.Println(client.URL)

	connected, err := client.DialForever()
	if err != nil {
		k.Log.Fatal(err.Error())
	}

	// Wait until connected
	<-connected

	// Call a method of client kite
	response, err := client.TellWithTimeout("foo", 4*time.Second)
	if err != nil {
		k.Log.Error(err.Error())
	}

	// Print the result
	res, _ := response.String()

	c.String(200, "foo-"+res)
}

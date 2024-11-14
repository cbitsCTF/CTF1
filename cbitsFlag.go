package main

import (
	"encoding/hex"
	"fmt"
)

func main() {

	var1 := map[int]string{1: "5f7072303672346d7d", 2: "3733445f345f4730", 3: "4342495453", 4: "7b3378336375"}

	var2 := var1[3] + var1[4] + var1[2] + var1[1]
	var3, _ := hex.DecodeString(var2)

	fmt.Println(string(var3))
}

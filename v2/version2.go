package gosecissue

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func DoThing() {
	fmt.Printf("hi v2 %v\n", gin.DebugMode)

}

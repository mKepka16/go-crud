package helpers

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(data interface{}) {
	userJSON, err := json.MarshalIndent(data, "", "    ") // Indentation with 4 spaces
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println(string(userJSON))
}

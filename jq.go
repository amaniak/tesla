package tesla

import "encoding/json"
import "os"

// Crappy JS0N purry prunt0r ...

func JQPrinter(jsonString string) {

	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(jsonString), &dat); err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(dat, "", "  ")
	if err != nil {
		panic(err)
	}
	b2 := append(b, '\n')
	os.Stdout.Write(b2)

}

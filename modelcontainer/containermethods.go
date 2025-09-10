package modelcontainer

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	_____________________________________________
	Pide esos parametros para llenar la struct :v
	_____________________________________________

*/

func (c *Container) SetJSONFields() {
	//el path donde sera creado el archivo json
	pathStr := fmt.Sprintf("./api/cache/containers/%s.json", c.Header.HeaderName)
	//serializa el los campos del container a JSON
	contentJson, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		fmt.Println("could not be serialized;", err)
		return
	}
	//crea el archivo json
	fileJson, err := os.OpenFile(pathStr, os.O_APPEND|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		fmt.Println("The json file could not be created;", err)
		return
	}
	defer fileJson.Close()

	//escribe el contenido json en el
	_, err = fileJson.Write(contentJson)
	if err != nil {
		fmt.Println("could not write to json;", err)
		return
	}
}

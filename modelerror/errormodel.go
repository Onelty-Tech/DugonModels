package modelerror

import (
	"fmt"
	"os"
)

/*
estructura aun no implementada
*/
type ErrorHandling struct {
	Type string `json:"type"` // tipos de errores para mas adelante ejemplo: "Warning","Error",etc
	Err  string `json:"error"`
	Uuid string `json:"uuid"` // El UUID Del canal de respuesta(chan *modelerror.ErrorHandling)
}

// const (
// 	ErrorTypeFatal   = "fatal"
// 	ErrorTypeWarning = "warning"
// 	ErrorTypeInfo    = "info"
// )

func HandleError(e ErrorHandling) {
	switch e.Type {
	case "Fatal":
		fmt.Printf("Fatal Error: %v\n", e.Err)
		os.Exit(1)
	case "Warning":
		fmt.Printf("Warning: %v\n", e.Err)
	case "Info":
		fmt.Printf("Info: %v\n", e.Err)
	default:
		fmt.Printf("error: Unknown error type: %v\n", e.Err)
	}
}

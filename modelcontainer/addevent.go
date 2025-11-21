package modelcontainer

import (
	"fmt"

	"github.com/Onelty-Tech/DugonLibrary/lib/sysutils/helpers"
)

const (
	const_Tried     = "Tried"
	const_Condition = "Condition"
)

/*
Crea y agrega un evento a la lista de eventos de la estructura
*/
func (c *Container) AddCustomEvent(behavior []map[string]interface{}) error {
	for _, val := range behavior {
		tried, err := helpers.GetString(val, const_Tried)
		if err != nil {
			return fmt.Errorf("error AddCustomEvent: %w", err)
		}
		cond, err := helpers.GetStringMapString(val, const_Condition)
		if err != nil {
			return fmt.Errorf("error AddCustomEvent: %w", err)
		}
		inCase := InCase{
			InCaseTried: tried,
			Condition:   cond,
			Behaviors:   val,
		}
		//eliminamos las 2 keys de los behaviors
		delete(val, const_Tried)
		delete(val, const_Condition)
		c.Body.Events = append(c.Body.Events, inCase)
	}
	return nil
}

/*
Funcion que remueve todos los eventos en el container
*/
func (c *Container) RemoveEvents() {
	c.Body.Events = nil
}

/*
Funcion que devuelve el slice de eventos.
*/
func (c *Container) GetEvents() []InCase {
	return c.Body.Events
}

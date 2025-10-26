package modelcontainer

/*
Crea y agrega un evento a la lista de eventos de la estructura
*/
func (c *Container) AddCustomEvent(tried string, condition map[string]string, behavior map[string]interface{}) {
	inCase := InCase{
		InCaseTried: tried,
		Condition:   condition,
		Behaviors:   behavior,
	}
	c.Body.Events = append(c.Body.Events, inCase)
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

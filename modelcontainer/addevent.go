package modelcontainer

/*
Crea y agrega un evento a la lista de eventos de la estructura
*/
func (c *Container) AddCustomEvent(tried string, condition map[string]string, action map[string]interface{}) {
	inCase := InCase{
		InCaseTried: tried,
		Condition:   condition,
		Action:      action,
	}
	c.Events = append(c.Events, inCase)
}

/*
	Funcion que remueve todos los eventos en el container
*/
func (c *Container) RemoveEvents() {
	c.Events = nil
}

package modelcontainer

func (c *Container) AddCustomAction(actions map[string]any) {
	inCase := InCase{
		InCaseTried: "",
		Condition:   map[string]string{},
		Behaviors:   actions,
	}
	c.Body.Actions = append(c.Body.Actions, inCase)
}

/*
Obtiene el slice de actions.
*/
func (c *Container) GetActions() Actions {
	return c.Body.Actions
}

/*
Remueve todas las acciones.
*/
func (c *Container) RemoveActions() {
	c.Body.Actions = nil
}

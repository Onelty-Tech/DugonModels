package modelcontainer

func (c *Container) AddCustomAction(actions []map[string]any) {
	//hacemos un for en el slice de accioes y agregamos cada uno al slice de acciones
	/*
		se ve innecesario esto, podriamos simplificarlo y agregarlo directamente
		en la struct sin tener que crear una struct individual por cada accion, pero se piensa agregar
		mas funcionalidades en el futuro, por eso se deja asi, aunque ahorita puede ser un gasto de memoria eso si.
	*/
	for _, val := range actions {
		inCase := InCase{
			InCaseTried: "",
			Condition:   map[string]string{},
			Behaviors:   val,
		}
		c.Body.Actions = append(c.Body.Actions, inCase)
	}
}

/*
Obtiene el slice de actions.
*/
func (c *Container) GetActions() []InCase {
	return c.Body.Actions
}

/*
Remueve todas las acciones.
*/
func (c *Container) RemoveActions() {
	c.Body.Actions = nil
}

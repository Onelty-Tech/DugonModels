package modelcontainer

func (c *Container) AddCustomAction(actionName string, actions map[string]any) {
	c.Header.HeaderAction = Action{
		ActionName: actionName,
		Actions:    actions,
	}
}

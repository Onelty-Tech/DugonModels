package modelcontainer

type Update struct {
	Version int    `json:"Version,omitempty"`
	Changes string `json:"Changes,omitempty"`
}

/*
|____________________________|
|Container y todos sus campos|
|____________________________|
*/
type Header struct {
	HeaderName            string `json:"Name,omitempty"`
	HeaderDeliveryQuality string `json:"delivery,omitempty"` // este campo sirve para anotar el canal del reenvio de errores o algun inconveniente relacionado con el contenedor enviado
	HeaderSender          string `json:"Sender,omitempty"`
	HeaderHost            string `json:"Host,omitempty"`

	/*
		HeaderAction          Action `json:"Action,omitempty"`
	*/
	HeaderChannel string `json:"Channel,omitempty"`
	HeaderTarget  string `json:"Target,omitempty"`
	HeaderDate    string `json:"Date,omitempty"`
	Pack          Pack   `json:"Pack"`
}

type Options struct {
	OptionsTypeUpdate bool `json:"TypeUpdate"`
}

type InCase struct {
	InCaseTried string            `json:"Tried,omitempty"`
	Condition   map[string]string `json:"Condition,omitempty"`
	// un mapa de actions dentro de otro mapa para meter funcionalidades unicas a cada action
	Behaviors map[string]any `json:"Behaviors,omitempty"`
}

type Body struct {
	Events      []InCase `json:"Events"`
	Actions     []InCase `json:"Actions"`
	BodyContent string   `json:"Content,omitempty"`
}

type Pack struct {
	Package string   `json:"Package,omitempty"`
	Update  Update   `json:"Update"`
	Logs    []string `json:"Logs,omitempty"`
	History []Update `json:"History,omitempty"`
}

type Container struct {
	Header  Header  `json:"Header"`
	Options Options `json:"Options"`
	Body    Body    `json:"Body"`
}

/*
Metodo de InCase que obtine los comportamientos.
*/
func (inCase *InCase) GetBehaviorCase() map[string]any {
	return inCase.Behaviors
}

/*
Metodo de InCase que remuve todos los comportamientos del InCase
*/
func (inCase *InCase) RemoveBehaviors() {
	inCase.Behaviors = nil
}

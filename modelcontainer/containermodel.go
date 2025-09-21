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
	HeaderAction          Action `json:"Action,omitempty"`
	HeaderChannel         string `json:"Channel,omitempty"`
	HeaderTarget          string `json:"Target,omitempty"`
	HeaderDate            string `json:"Date,omitempty"`
	Pack                  Pack   `json:"Pack"`
}

type Options struct {
	OptionsTypeUpdate bool `json:"TypeUpdate"`
}

type Body struct {
	BodyContent string `json:"Content,omitempty"`
}

type Pack struct {
	Package string   `json:"Package,omitempty"`
	Update  Update   `json:"Update"`
	Logs    []string `json:"Logs,omitempty"`
	History []Update `json:"History,omitempty"`
}

type InCase struct {
	InCaseTried string            `json:"Tried,omitempty"`
	Condition   map[string]string `json:"Condition,omitempty"`
	// un mapa de actions dentro de otro mapa para meter funcionalidades unicas a cada action
	Action map[string]any `json:"Action,omitempty"`
}

type Events []InCase

type Action struct {
	ActionName string         `json:"Name"`
	Actions    map[string]any `json:"Actions"`
}

type Container struct {
	Header  Header   `json:"Header"`
	Options Options  `json:"Options"`
	Body    Body     `json:"Body"`
	Events  []InCase `json:"Events"`
}

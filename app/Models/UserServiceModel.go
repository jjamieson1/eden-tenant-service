package Models

type UserProvider struct {
	Id 			string 		`json:"id"`
	Name 		string 		`json:"name"`
	PluginName	string 		`json:"pluginName"`
	CalloutUrl	string 		`json:"calloutUrl"`
}

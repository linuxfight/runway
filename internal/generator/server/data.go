package server

type TemplateData struct {
	Package     string
	Module      string
	Routes      []RuntimeRoute
	NeedContext bool
	NeedHTTP    bool
}

func BuildTemplateData(routes []RuntimeRoute, pkg, module string, needContext bool, needHTTP bool) TemplateData {
	return TemplateData{
		Package:     pkg,
		Module:      module,
		Routes:      routes,
		NeedContext: needContext,
		NeedHTTP:    needHTTP,
	}
}

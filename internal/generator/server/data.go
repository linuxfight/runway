package server

type TemplateData struct {
	Package string
	Module  string
	Routes  []RuntimeRoute
}

func BuildTemplateData(routes []RuntimeRoute, pkg, module string) TemplateData {
	return TemplateData{
		Package: pkg,
		Module:  module,
		Routes:  routes,
	}
}

package server

type TemplateData struct {
	PackageName string
	Routes      []RuntimeRoute
}

func BuildTemplateData(routes []RuntimeRoute, pkg string) TemplateData {
	return TemplateData{
		PackageName: pkg,
		Routes:      routes,
	}
}

package api

func XMLMediaType(prefix string, ctype string) MediaType {
	if prefix == "" {
		prefix = "xml"
	}

	if ctype == "" {
		ctype = "application/xml"
	}

	return xmlMediaType{ctype: ctype}
}

type xmlMediaType struct{
	pfx   string
	ctype string
}

func (t xmlMediaType) Prefix() string {
	return t.pfx
}

func (t xmlMediaType) ContentType() string {
	return t.ctype
}

package api

func XMLMediaType(ctype string) MediaType {
	if ctype == "" {
		ctype = "application/xml"
	}

	return xmlMediaType{ctype: ctype}
}

type xmlMediaType struct{
	ctype string
}

func (t xmlMediaType) ContentType() string {
	return t.ctype
}

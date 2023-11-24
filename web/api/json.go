package api

func JSONMediaType(ctype string) MediaType {
	if ctype == "" {
		ctype = "application/json"
	}

	return jsonMediaType{ctype: ctype}
}

type jsonMediaType struct {
	ctype string
}

func (t jsonMediaType) ContentType() string {
	return t.ctype
}

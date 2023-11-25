package api

func JSONMediaType(prefix string, ctype string) MediaType {
	if prefix == "" {
		prefix = "json"
	}

	if ctype == "" {
		ctype = "application/json"
	}

	return jsonMediaType{pfx: prefix, ctype: ctype}
}

type jsonMediaType struct {
	pfx   string
	ctype string
}

func (t jsonMediaType) Prefix() string {
	return t.pfx
}

func (t jsonMediaType) ContentType() string {
	return t.ctype
}

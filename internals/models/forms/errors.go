package forms

type errors map[string][]string

// Add an error to the errors map.
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Fetch the error message from the errors map.
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}

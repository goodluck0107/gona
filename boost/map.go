package boost

func MergeMap(a, b map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range a {
		m[k] = v
	}
	for k, v := range b {
		m[k] = v
	}
	return m
}

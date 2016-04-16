package whisk

func Map_equal(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}
	for mk1, m1v := range m1 {
		if m2v, ok := m2[mk1]; !ok || m2v != m1v {
			return false
		}
	}
	return true
}

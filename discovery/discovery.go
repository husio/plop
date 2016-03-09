package discovery

var services = map[string][]string{
	"blog":     []string{"localhost:8004"},
	"auth":     []string{"localhost:8006"},
	"currtime": []string{"localhost:8008"},
}

// All return list of addresses for given service.
func All(service string) []string {
	return services[service]
}

// Any return address of any service with given name.
func Any(service string) string {
	all := services[service]
	if len(all) == 0 {
		return ""
	}
	return all[0]
}

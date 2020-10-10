package delivery

type GeoACL struct{ Countries, Regions ACL }

func (g GeoACL) Empty() bool { return g.Countries.Empty() && g.Regions.Empty() }
func (g GeoACL) Allow(country, region string) bool {
	if country != "" && !g.Countries.Allowed(country) {
		return false
	}
	if region != "" && !g.Regions.Allowed(region) {
		return false
	}
	return true
}

type ACL struct{ Allow, Block []string }

func (a ACL) Empty() bool { return len(a.Allow)+len(a.Block) == 0 }
func (a ACL) Allowed(entity string) bool {
	switch {
	case len(a.Allow) > 0 && len(a.Block) > 0:
		return contains(entity, a.Allow) && !contains(entity, a.Block)
	case len(a.Allow) > 0 && len(a.Block) == 0:
		return contains(entity, a.Allow)
	case len(a.Allow) == 0 && len(a.Block) > 0:
		return !contains(entity, a.Block)
	default:
		return true
	}
}

func contains(elem string, set []string) bool {
	for _, e := range set {
		if e == elem {
			return true
		}
	}
	return false
}

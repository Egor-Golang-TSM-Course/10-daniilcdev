package counters

type Counter interface {
	Count(ln string, out map[string]int)
}

package craft

type itemClass struct {
	Name         string   `json:"name"`
	Format       string   `json:"format"`
	Example      string   `json:"example"`
	VerbVariants []string `json:"verb_variants"`
}

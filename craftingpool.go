package craft

type CraftingPool struct {
	itemTemplates []itemTemplate
	itemClasses   map[string]itemClass
	materials     map[string]material
	details       map[string]detail
}

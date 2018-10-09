package craft

type ResourcePool struct {
	itemTemplates []itemTemplate
	itemClasses   map[string]itemClass
	materials     map[string]material
	details       map[string]detail
	modifiers     map[string]modifier
}

package models

type Selector interface {
	judge(context Context) bool
	getName() string
}

type Eq struct {
	key   string
	value interface{}
}

func (o *Eq) judge(context Context) bool {
	current := context.state.cache[o.key]
	if current == "" {
		return false
	}
	return current == o.value
}

func (o *Eq) getName() string {
	return "$eq"
}

type True struct {
}

func (o *True) judge(context Context) bool {
	return true
}

func (o *True) getName() string {
	return "$true"
}

type And struct {
}

func (And) judge(context Context) bool {
	return true
}

func (And) getName() string {
	return "$and"
}

type Or struct {
}

func (Or) judge(context Context) bool {
	return true
}

func (Or) getName() string {
	return "$or"
}

type Gt struct {
}

func (Gt) judge(context Context) bool {
	return true
}

func (Gt) getName() string {
	return "$gt"
}

type Gte struct {
}

func (Gte) judge(context Context) bool {
	return true
}

func (Gte) getName() string {
	return "$gte"
}

type Lt struct {
}

func (Lt) judge(context Context) bool {
	return true
}

func (Lt) getName() string {
	return "$lt"
}

type Lte struct {
}

func (Lte) judge(context Context) bool {
	return true
}

func (Lte) getName() string {
	return "$lte"
}

type SelectorFactory struct {
}

func (*SelectorFactory) CreateSelector(selectorType string) Selector {
	switch selectorType {
	case "$true":
		return &True{}
	case "$eq":
		return &Eq{}
	case "$and":
		return &And{}
	case "$or":
		return &Or{}
	case "$gt":
		return &Gt{}
	case "$gte":
		return &Gte{}
	case "$lt":
		return &Lt{}
	case "$lte":
		return &Lte{}
	default:
		return nil
	}
}

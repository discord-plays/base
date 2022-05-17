package iface

type Command interface {
	GetName() string
	Execute()
}

package vm

import (
	"errors"

	"github.com/loomnetwork/go-loom"

	"github.com/loomnetwork/loomchain/state"
)

type VM interface {
	Create(caller loom.Address, code []byte, value *loom.BigUInt) ([]byte, loom.Address, error)
	Call(caller, addr loom.Address, input []byte, value *loom.BigUInt) ([]byte, error)
	StaticCall(caller, addr loom.Address, input []byte) ([]byte, error)
	GetCode(addr loom.Address) ([]byte, error)
}

type Factory func(state.State) (VM, error)

type Manager struct {
	vms map[VMType]Factory
}

func NewManager() *Manager {
	return &Manager{
		vms: make(map[VMType]Factory),
	}
}

func (m *Manager) Register(typ VMType, fac Factory) {
	m.vms[typ] = fac
}

func (m *Manager) InitVM(typ VMType, s state.State) (VM, error) {
	fac, ok := m.vms[typ]
	if !ok {
		return nil, errors.New("vm type not found")
	}

	return fac(s)
}

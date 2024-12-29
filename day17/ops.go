package day17

type opHandler func(vm *VM, operand int)

func advHandler(vm *VM, operand int) {
	vm.A = vm.dv(vm.A, operand)
}

func bxlHandler(vm *VM, operand int) {
	vm.B ^= operand
}

func bstHandler(vm *VM, operand int) {
	vm.B = vm.Combo(operand) % 8
}

func jnzHandler(vm *VM, operand int) {
	if vm.A != 0 {
		vm.instrp = operand
	}
}

func bxcHandler(vm *VM, _ int) {
	vm.B ^= vm.C
}

func outHandler(vm *VM, operand int) {
	vm.outbuf = append(vm.outbuf, vm.Combo(operand)%8)
}

func bdvHandler(vm *VM, operand int) {
	vm.B = vm.dv(vm.A, operand)
}

func cdvHandler(vm *VM, operand int) {
	vm.C = vm.dv(vm.A, operand)
}

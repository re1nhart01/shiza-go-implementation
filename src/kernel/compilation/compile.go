package compilation

import (
	"sync"

	"shiza.io/src/kernel/ast_tree"
)

type ShizaCompilator struct {
}

func NewPompilator() ShizaCompilator {
	return ShizaCompilator{}
}

func (cmp *ShizaCompilator) TryToCompile(wg *sync.WaitGroup, node *ast_tree.StatementsNode) {
	wg.Done()
}

func (cmp *ShizaCompilator) compile() {

}

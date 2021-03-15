package alc_gorm

import (
	"fmt"
	"testing"
)

func TestPagination_PaginationCompute(t *testing.T) {
	p := NewPagination(1, 5)
	p.Compute(84)
	fmt.Printf("%#v", p)
}

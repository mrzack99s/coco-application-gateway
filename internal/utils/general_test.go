package utils_test

import (
	"fmt"
	"testing"

	"github.com/mrzack99s/coco-application-gateway/internal/utils"
)

func TestFindAndDelete(t *testing.T) {

	g := []int{1, 2, 3, 4}
	fmt.Println(g)
	g = utils.FindAndDeleteInt(g, 3)
	fmt.Println(g)

}

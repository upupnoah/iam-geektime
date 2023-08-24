package behavioral_patterns

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	xhs := &XiHongShi{}
	doCook(xhs)

	fmt.Println("=====> 做另外一道菜")

	cjd := &ChaoJiDan{}
	doCook(cjd)

}

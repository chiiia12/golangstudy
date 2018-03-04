package eval

import (
	"strconv"
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', 4, 32)
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return b.x.String() + string(b.op) + b.y.String()
}

func (c call) String() string {
	var joinlist []string
	for _, v := range c.args {
		joinlist = append(joinlist, v.String())
	}
	return fmt.Sprintf("%v(%v)", c.fn, strings.Join(joinlist, ","))
}

//testで同じ構文ツリーになることを検査しないといけない

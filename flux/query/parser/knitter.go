package parser

type Knitter struct {
	size int
	a string
}

func (k *Knitter) Knit(s string) {
	if s == "(" {
		return
	}

	if k.size == 0 {
		k.a = s
		k.size++
		return
	}
}

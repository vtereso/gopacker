package shared

type S struct{
	value string
}

var s *S = &S{}

func (s *S) setValue(value string){
	(*s).value=value
}

func getS() *S {
	return s
}
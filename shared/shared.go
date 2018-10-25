package shared

type S struct{
	value string
}

var s *S = &S{}

func (s *S) SetValue(value string){
	(*s).value=value
}

func GetS() *S {
	return s
}
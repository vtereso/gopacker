package pack2

import (
    "github.com/vtereso/gopacker/shared"
    "fmt"
)


func SetPack2() {
    shared.SetValue("Pack 2 value")
}

func GetPack2() {
    val := shared.GetS()
    fmt.Println("%v %p",val,val)
}
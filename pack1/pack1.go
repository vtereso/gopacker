package pack1

import (
    "github.com/vtereso/gopacker/shared"
    "fmt"
)


func SetPack1() {
    shared.SetValue("Pack 1 value")
}

func GetPack1() {
    val := shared.GetS()
    fmt.Println("%v %p",val,val)
}
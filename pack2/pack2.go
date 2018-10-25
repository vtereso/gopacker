package pack2

import (
    "github.com/vtereso/gopacker/shared"
    "fmt"
)


func setPack2() {
    shared.setValue("Pack 2 value")
}

func getPack2() {
    val := shared.getS()
    fmt.Println("%v %p",val,val)
}
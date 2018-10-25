package pack1

import (
    "github.com/vtereso/gopacker/shared"
    "fmt"
)


func setPack1() {
    shared.setValue("Pack 1 value")
}

func getPack1() {
    val := shared.getS()
    fmt.Println("%v %p",val,val)
}
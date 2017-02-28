package main
import "fmt"
////////////////////////
type Appliance interface{
    Push(interface{}) 
}
type ArrayInt struct{
     c []int
     size int
}

func (this *ArrayInt)Push(a interface{}){
     switch z := a.(type) {
        case int:
            this.c = append(this.c,z)
     }
}
type ArrayString struct{
     c []string
     name string
}

func (this *ArrayString)Push(a interface{}){
    switch z := a.(type) {
        case string:
            this.c = append(this.c,z)
     }
}
 
func ArrayList(a interface{}) Appliance {
        b := new(ArrayString)
        switch v := a.(type) {
		case []string:  
                    b := new(ArrayString)
                    b.name = "welcome"
                    b.c = v
		case []int:
                    b := new(ArrayInt)
                    b.size = 4
                    b.c = v
	}
        return b
}
////////////////////////
func main() {
      a := ArrayList([]string{"v"})
      b := ArrayList([]int{4})
      a.Push("v");
      b.Push(5);
   
}
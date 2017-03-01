package main
////////////////////////
import "fmt"

type AString []string

type stringArray interface {
    join(string) string
    pop() string
    indexOf(string) int
    push(args ...interface{}) int
    Map(func(mapType)string) *AString
    Filter(func(filterType)bool) *AString
}


////////////////////////////////
type mapType struct{
    i int
    val string
    table []string
}

type filterType struct{
    i int
    val string
    table []string
}
////////////////////////////////
func (this *AString) indexOf(a string) int{
    for i,v := range *this {
        if v == a {
            return i
        }  
     }
    return -1 
}
////////////////////////////////
func (this *AString) pop() string{
     a := *this
     var d string = a[len(a)-1]
     *this = a[:len(a)-1]
     return d
}

func (this *AString) push(args ...interface{}) int{
    for _, item := range args {
         switch v := item.(type) {
             case string:
                  *this = append(*this, v)
         }
     }
     return len(*this)
}
////////////////////////////////
func (this *AString) Map(a func(mapType)string) *AString{
     d := *this
     for i,v := range d {
	  d[i] = a(mapType{i:i,val:v,table:d})
     }
     return this
}

func (this *AString) Filter(a func(filterType)bool) *AString{
     d := []string{}
     for i,v := range *this {
          if a(filterType{table:d,val:v,i:i}) {
             d = append(d,v)
          }
     }
     *this = d
     return this
}
////////////////////////////////

func (this *AString) join(a string) string{
     var r string = ""
     for i,v := range *this {
          if i != 0 {
              r += a 
          }
          r += v
     }
    return r
}

func ArrayString(args ...interface{}) stringArray {
    a := new(AString)
    for _, item := range args {
        switch v := item.(type) {
            case string:
                 *a = append(*a, v)
        }
    }
    return stringArray(a) 
}



////////////////////////
func main() {
      a := ArrayString("a","b")
      r1 := a.push("c","e")
      fmt.Println(a.indexOf("e"))
      fmt.Println(r1)
      fmt.Println(a.join(","))
      r2 := a.pop()
      fmt.Println(r2)
      fmt.Println(a) 

      c := func(this mapType)string{
      	return this.val + "_super"
      }
      d := func(this filterType)bool{
      	return this.val == "a_super"
      }

      a.Map(c)
      fmt.Println(a)
      a.Filter(d)
    
      fmt.Println(a)
}
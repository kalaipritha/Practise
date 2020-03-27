package main
import "fmt"

func swap(a, b int) (int ,int) {
    return b,a 
}

func calculate(a []int) []int{
    var flag = true
    for i :=0 ;i<len(a)-1;i++ {
        if flag == true{
            if a[i]>a[i+1]{
                a[i],a[i+1] = swap(a[i],a[i+1])
            }
        }else{
            if a[i]<a[i+1] {
                a[i],a[i+1] = swap(a[i],a[i+1])
            }
        }
        flag = !flag
    }
    return a
}


func main() {
    arr := []int{1, 4, 2, 3}
    arr1 := calculate(arr )
    fmt.Println(arr1)
}
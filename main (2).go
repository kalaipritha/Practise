package main
import "fmt"

func calculate(arr []int) []int{
    var low, mid, hi = 0 ,0, len(arr)-1
    fmt.Println("hii",hi)
    for mid<=hi{
        fmt.Println("the value of mid is",arr[mid])
        switch arr[mid] {
            case 0:{
                fmt.Println("000")
                temp := arr[low]
                arr[low]=arr[mid]
                arr[mid]=temp
                mid++
                low++
                break;
            }
            case 1:{
                fmt.Println("111")
                mid++
                break
            }
            case 2:{
                fmt.Println("222")
                temp := arr[hi]
                arr[hi]=arr[mid]
                arr[mid]=temp
                hi--
                break
            }
            }
    }
    return arr
}

func main() {
    arr := []int{1, 2, 0, 0, 0,2,2};
    arr1 := calculate(arr )
    fmt.Println(arr1)
}
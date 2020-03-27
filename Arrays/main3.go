package main
import "fmt"

func calculate(arr []int) []int{
    var mid, hi = 0, len(arr)-1
    fmt.Println("hii",hi)
    for mid<=hi{
        fmt.Println("the value of mid is",arr[mid])
        switch arr[mid] {
            case 0:{
                fmt.Println("111")
                mid++
                break
            }
            case 1:{
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

func calculate2(arr []int)[]int{
    var l, r = 0,len(arr)-1
    for l<r{
        if arr[l]==0 {
            l++
        }else if arr[r]==1{
            r--
        }else{
            arr[l]=0
            arr[r]=0
            l++
            r--
        }
    }
    return arr
}

func main() {
    arr := []int{1, 0, 0, 1,1,0,1, 0,0,1};
    arr1 := calculate(arr )
    fmt.Println(arr1)
    //another approach
    arr1 = calculate2(arr)
    fmt.Println(arr1)
}
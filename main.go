package main
import "fmt"

func calculate(arr []int) []int{
    count := 0
    for i:=0;i<len(arr);i++{
        if(arr[i]!=0){
            arr[count]=arr[i]
            count++
        }
    }
    for count!=len(arr) {
        arr[count]=0
        count++;
    }
    return arr
}
func main() {
    arr := []int{1, 2, 0, 0, 0, 3, 6};
    arr1 := calculate(arr )
    fmt.Println(arr1)
}
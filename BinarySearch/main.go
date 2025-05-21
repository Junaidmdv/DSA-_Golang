package main
import "fmt"

func main() {
  ar:=[]int{5,6,7,2,3}
  
  st:=0 
  end:=len(ar)-1
  ans:=-1
  for st<=end{
      if ar[st] <= ar[end] {
          ans=st
          break
      } 
      mid:=st+(end-st)/2
      
      if ar[mid] > ar[end]{
          st=mid+1
      }else{
          end=mid-1
      }
      
      
  }
  fmt.Println(ans)
}
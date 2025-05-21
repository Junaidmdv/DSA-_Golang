// package main

// import "fmt"

// type Stack struct {
// 	Items []int
// }

// func (S *Stack) Push(n int) {
// 	S.Items = append(S.Items, n)

// }

// func (S *Stack) Pop() int {
// 	if len(S.Items) == 0 {
// 		fmt.Println("Empty stack")
// 	}
// 	res := S.Items[len(S.Items)-1]
// 	S.Items = S.Items[:len(S.Items)-1]
// 	return res
// }

// func (S *Stack) IsEmpty() bool {
// 	return len(S.Items) == 0
// }

// func (S *Stack)Peek()int{
// 	return S.Items[len(S.Items)-1]
// }

// // func SortStack(s *Stack) *Stack{
// //      tempStack:=&Stack{}

// // 	 for !s.IsEmpty() {
// // 		 temp:=s.Pop()

// // 		 for !tempStack.IsEmpty() && tempStack.Peek() > temp {
// // 			s.Push(tempStack.Peek())
// // 			tempStack.Pop()
// // 		 }

// // 		 tempStack.Push(temp)

// // 	 }

// // 	 return tempStack

// // }

// //implementation of queue using stack

// func Dequeue(s *Stack)*Stack{

// 	 tempStack:=&Stack{}

// 	for !s.IsEmpty(){
// 		tempStack.Push(s.Pop())
// 	}

// 	tempStack.Pop()

// 	for !tempStack.IsEmpty(){
// 		s.Push(tempStack.Pop())
// 	}
// 	return s
// }

// func main() {
// 	s := &Stack{}
// 	for i := 10; i >= 1;i-- {
// 		s.Push(i)
// 	}
// 	fmt.Println(s.Pop())
// 	for _, val := range s.Items {
// 		fmt.Println(val)
// 	}

// 	fmt.Println(s.IsEmpty())

// 	res:=Dequeue(s)

// 	for _,val:=range res.Items{
// 		fmt.Println(val)
// 	}

// 	// sortedstack:=SortStack(s)

// 	// for _,val:=range sortedstack.Items{
// 	// 	 fmt.Println(val)
// 	// }
// }

// rev string using stack

// package main
// import "fmt"

// type stack struct{
// 	Items []string
// }

// func (s *stack)pop()string{
// 	peak:=s.Items[len(s.Items)-1]
// 	s.Items=s.Items[:len(s.Items)-1]

// 	return peak

// }

// func (s *stack)IsEmpty()bool{
// 	return len(s.Items)==0
// }

// func (s *stack)push(data string){
//     s.Items=append(s.Items, data)

// }

// func main(){
// 	str:="Hello world"
// 	s:=&stack{}

// 	for _,val:=range str {
// 		 s.push(string(val))
// 	}
//    revestr:=""
// 	for !s.IsEmpty(){
// 		revestr+=s.pop()
// 	}
// 	fmt.Println(revestr)
// }

//delete midle element of stack

package main

import "fmt"

type stack struct {
	Item []int
}

func (s *stack) Push(Data int) {
	s.Item = append(s.Item, Data)
}

func (s *stack) Peak() int {
	return s.Item[len(s.Item)-1]
}
func (s *stack) Pop() int {
	popedElement := s.Peak()
	s.Item = s.Item[:len(s.Item)-1]
	return popedElement

}

func (s *stack) IsEmpty() bool {
	return len(s.Item) == 0
}

func DeleteMidle(s *stack) {
	res := make([]int, 0, len(s.Item))
	i := 0
	for !s.IsEmpty() {
		n := s.Pop()
		res = append(res, n)
		i++
	}
	mid:=0
	if len(res)%2==1{
       mid = len(res) / 2
	}else{
		 mid = len(res) / 2-1
	}
	
	for i := len(res) - 1; i >= 0; i-- {

		if i == mid {
			continue
		}
		s.Push(res[i])
	}

	for _, val := range s.Item {
		fmt.Println(val)
	}

}

func main() {
	s := &stack{}

	for i := 1; i <= 10; i++ {
		s.Push(i)
	}

	DeleteMidle(s)

}

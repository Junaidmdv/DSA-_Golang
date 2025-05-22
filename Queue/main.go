
//Queue using array

package main 
import "fmt"


// type Queue struct{
//     Item []int
// }


// func (Q *Queue)Enqueue(Data int){
// 	Q.Item=append(Q.Item, Data)
// }

// func (Q *Queue)Dequeue()int{
// 	 num:=Q.Item[0]
// 	 Q.Item=Q.Item[1:]
//    return num
// }

// func (Q *Queue)Rear()int{
// 	return Q.Item[len(Q.Item)-1]
// }




// func main(){
// 	 Q:=&Queue{}
// 	 Q.Enqueue(23)
// 	 Q.Enqueue(32)
// 	 Q.Enqueue(21)
// 	 Q.Enqueue(12)
// 	fmt.Println(Q.Dequeue()) 
// 	fmt.Println()
// 	 for _,val:=range Q.Item{
// 		  fmt.Println(val)
// 	 }

// }

//Queue using linked list 

// type Node struct{
// 	Data int 
// 	Next *Node
// }

// type Queue struct{
// 	Head *Node
// 	Tail  *Node
// }

// func(Q *Queue)Enqueue(Data int){
// 	 NewNode:=&Node{Data:Data,}
// 	 if Q.Head == nil {
// 		Q.Head=NewNode
// 		Q.Tail=NewNode
// 	 }else{
// 		 Q.Tail.Next=NewNode
// 		 Q.Tail=NewNode
// 	 }
// }

// func (Q *Queue)DeQueue()int{
// 	if Q.Head == nil {
//          return 0
// 	}
// 	val:=Q.Head.Data

// 	Q.Head=Q.Head.Next
// 	return val
// }

// func (Q *Queue)Display(){
// 	current:=Q.Head
// 	for current != nil {
// 		fmt.Printf("%d ->",current.Data)
// 		current=current.Next
// 	}
// }


// func main(){
//    Q:=&Queue{}
//    fmt.Println(Q.DeQueue()) 
//    Q.Enqueue(23)
//    Q.Enqueue(32)
//    Q.Enqueue(33)
//    Q.Enqueue(35)
//    Q.Display()

   

// }

//Double Ender Queue

// type DeQueue struct{
//    Items []int
// }

// func(D *DeQueue)Front_EnQueue(data int){
// 	 ar:=[]int{data}
// 	 D.Items=append(ar,D.Items...)
	 
	 
// }

// func (D *DeQueue)Front_DeQueue()int{
// 	if len(D.Items)==0{
// 		 return -1
// 	}
// 	res:=D.Items[0]
// 	D.Items=D.Items[1:]
//    return  res
// }

// func(D *DeQueue)Rear_EnQueue(Data int){
// 	D.Items = append(D.Items, Data)
  
// }

// func (D *DeQueue)Rear_DeQueue()int{
// 	if D.Items==nil{
// 		return 0
// 	}
// 	num:=D.Items[len(D.Items)-1]
//     D.Items=D.Items[:len(D.Items)-1]
// 	return num
// }


// func main(){

// 	D:=&DeQueue{}
// 	D.Rear_EnQueue(21)
//     D.Front_EnQueue(23)
// 	D.Front_EnQueue(12)
// 	D.Front_EnQueue(1)
// 	D.Rear_EnQueue(32)
// 	D.Rear_DeQueue()
// 	D.Front_DeQueue()
// 	for _,val:=range D.Items{
// 		fmt.Println(val)
// 	}
// }

// Circular Queue


type CircularQueue struct {
	Item        []int
	Capacity    int
	CurrentSize int
	Front       int
	Rear        int
}

func Initialise(size int) *CircularQueue {
	return &CircularQueue{
		Item:        make([]int, size),
		Capacity:    size,
		CurrentSize: 0,
		Front:       0,
		Rear:        -1,
	}
}

func (C *CircularQueue) EnQueue(Data int) {
	if C.Capacity == C.CurrentSize {
		fmt.Println("Queue is filled")
		return
	}
	C.Rear = (C.Rear + 1) % C.Capacity
	C.Item[C.Rear] = Data
	C.CurrentSize++
}

func (C *CircularQueue) DeQueue() {
	if C.CurrentSize == 0 {
		fmt.Println("Empty Stack")
		return
	}
	C.Front = (C.Front + 1) % C.Capacity
	C.CurrentSize--
}

func (C *CircularQueue) Display() {
	if C.CurrentSize == 0 {
		fmt.Println("Empty Queue")
		return
	}
	start := C.Front

	for {
		fmt.Println(C.Item[start])

		if start == C.Rear {
			return
		}
		start = (start + 1) % C.Capacity
	}
}

func main() {
	C := Initialise(5)
	C.EnQueue(12)
	C.EnQueue(14)
	C.EnQueue(15)
	C.EnQueue(13)
	C.EnQueue(16)
	C.EnQueue(18)
	C.DeQueue()
	C.Display()
}

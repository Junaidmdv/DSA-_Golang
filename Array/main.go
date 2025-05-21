package  main 

type Node struct{
    Data int
    Next *Node
}

type LinkedList struct{
    Head  *Node
}


func (list *LinkedList)PushBack(Data int){
    NewNode:=Node{Data: Data,Next: nil}

    if list.Head == nil{
        list.Head=&NewNode
        return
    }

    current:=list.Head

    for current.Next !=nil{
        current=current.Next
    }

    current.Next=&NewNode
    

}

func (list *LinkedList)PushFront(Data int){
     if list.Head == nil{
        NewNode:=&Node{Data: Data,Next: nil }
        list.Head=NewNode
        return
     }

     NewNode:=&Node{Data: Data,Next: list.Head}
     list.Head=NewNode
}

func (list *LinkedList)PushBackValue(Data,BeforeValue int){
    
}


func (list *LinkedList)PushAfterValue(Data,AfterVaue int){

}



func main(){

    linkedlist:=LinkedList{}
    linkedlist.PushFront(23)
    linkedlist.PushBack(23)

}
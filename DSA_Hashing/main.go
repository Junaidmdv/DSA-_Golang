package main

import "fmt"

const size = 10

type HashTable struct {
	Items [size]*HashItems
}

type HashItems struct {
	Key   string
	Value int
}

func (H *HashTable) Insert(Key string, Value int) {
	index := Hash(Key)
	start := index

	for {
		if H.Items[index] == nil || H.Items[index].Key == Key {
			H.Items[index] = &HashItems{Key: Key, Value: Value}
			return
		}
		index = (index + 1) % size
		if start == index {
			fmt.Println("Hash table is filled")
			return
		}

	}
}

func (H *HashTable) Delete(Key string) {
	index := Hash(Key)
	start := index
	for {
		if H.Items[index].Key == Key {
			H.Items[index] = nil
			return
		}
		index = (index + 1) % size
		if start == index {
			return
		}
	}

}

func (H *HashTable) Get(Key string) (int, bool) {
	index := Hash(Key)
	start := index
	for {

		if H.Items[index] != nil && H.Items[index].Key == Key {
			return H.Items[index].Value, true
		}

		index = (index + 1) % size
		if start == index {
			break
		}
	}
	return 0, false
}

func (H *HashTable) Display() {

	for _, val := range H.Items {
		if val != nil {
			fmt.Printf("%v:%v ", val.Key, val.Value)
		}
	}
	fmt.Println()
}

func Hash(Key string) int {
	sum := 0

	for _, val := range Key {
		sum += int(val)
	}
	return sum % size
}

func main() {
	hashtable := &HashTable{}
	hashtable.Insert("Junaid", 12)
	hashtable.Insert("abc", 13)
	hashtable.Insert("Ha", 32)
	val, exist := hashtable.Get("Ha")
	if exist {
		fmt.Println(val)
	}
	hashtable.Delete("Ha")
	hashtable.Display()
	fmt.Println(hashtable.Items)

}

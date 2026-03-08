package main
import "fmt"

type Stack struct {
    items []int
}

// Push element
func (s *Stack) Push(value int) {
    s.items = append(s.items, value)
}

// Pop element
func (s *Stack) Pop() int {
    if len(s.items) == 0 {
        fmt.Println("Stack is empty")
        return -1
    }

    index := len(s.items) - 1
    element := s.items[index]
    s.items = s.items[:index]

    return element
}

// Peek top element
func (s *Stack) Peek() int {
    return s.items[len(s.items)-1]
}

func main() {
    var s Stack

    s.Push(10)
    s.Push(20)
    s.Push(30)

    fmt.Println("Stack:", s.items)

    // fmt.Println("Pop:", s.Pop())
    fmt.Println("Peek:", s.Peek())
}
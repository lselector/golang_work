// how to make different containers
// https://github.com/tracymacding/priority-queue/blob/master/priority_queue.go
// https://gist.github.com/moraes/2141121
// https://github.com/oleiade/lane
// https://stackoverflow.com/questions/2818852/is-there-a-queue-implementation

package main

import (
    "fmt"
//    "container/heap"
)


queue := make([]int, 0)
// Push to the queue
queue = append(queue, 1)
// Top (just get next element, don't remove it)
x = queue[0]
// Discard top element
queue = queue[1:]
// Is empty ?
if len(queue) == 0 {
    fmt.Println("Queue is empty !")
}





func main() {
    fmt.Println("Hello")
}

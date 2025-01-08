func main() {
  c := make(chan int, 10) 
  defer close(c)
  go func() {
    for i := 0; i < 10; i++ {
      c <- i
    }
  }()
  for i := range c {
    fmt.Println(i)
  }
}
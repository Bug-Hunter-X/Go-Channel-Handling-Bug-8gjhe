func main() {
  c := make(chan int)
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
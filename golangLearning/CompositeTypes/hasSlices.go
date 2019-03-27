package main

var m = make(map[string]int)
func k(list []string) string {return fmt.Sprintf("%q", list)}
func Add(list []string) { m[k(list)] ++}
func Count(list []string) { return m[k(list)] }

func main() {
	var m = make(map[string]int)

}
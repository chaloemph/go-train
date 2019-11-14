package main
import ("fmt")

func main() {
	var age int = 25
	weight:= 68
	height:=175.5
	monday:= false
	tuesday:= true
	job,salary:= "full stack web developer",100000
	fmt.Println("age : ", age)
	fmt.Println("weight : ", weight)
	fmt.Println("height : ", height)
	fmt.Println("today is monday", monday)
	fmt.Println("today is tuesdat", tuesday)
	fmt.Println("today is monday and tuesday", tuesday && monday )
	fmt.Printf("job position %s has salary %d \n" , job, salary)
	

}
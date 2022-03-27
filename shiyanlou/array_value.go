package main

import "fmt"

func main(){
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	planetsMarkII := planets  // ←--- 复制planets数组
	planets[2] = "whoops"  // ←--- 修改数组元素，让出星际轨道
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
}

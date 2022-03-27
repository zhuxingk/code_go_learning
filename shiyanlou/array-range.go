package main


import "fmt"

func main()  {
	dwarfs := [5]string{"Ceres","Pluto","Haumea","Makemake","Eris"}
	for i,dwarf := range dwarfs{//使用关键字 range 可以取得数组中
		// 每个元素对应的索引和值，这种迭代方式使用的代码更少并且更不容易出错
		fmt.Println(i,dwarf)
	}
}


package main

import ("fmt")
	

func main(){
	// ganesh:= "OM ganeshay namah"
	// fmt.Println(ganesh)
	// nums:=[]int{}
	// nums=append(nums,10)
	// nums=append(nums,20)
	// fmt.Println(nums[0])
	// db := make(map[string] int)
	// db["utkarsh"]=1
	// value:=db["utkarsh"]
	// fmt.Println(value) 
	// for i:=0;i<len(nums); i++{
	// 	fmt.Println(nums[i])
	// }
	// a,e:=divide(10,0)
	

	// if e==nil {
	// 	fmt.Println(a)
	// }else{
	// 	fmt.Println(e)
	// }
	database:= Database{}
	database.db = make(map[string]string)
	//database.db["name"]="Utkarsh"

	database.Set("name", "Nigga")
	fmt.Println(database.db["name"])
	database.Set("name","NOT NIGGA")
	value,isExist:= database.Get("name")
	if isExist{
		fmt.Println(value)
	}else {
		fmt.Println("value doesnt exist")
	}
	isdone:= database.Del("name")
	if isdone {
		fmt.Println("deleted")
	}
	value,isExist=database.Get("name")
	if isExist{
		fmt.Println(value)
	}else {
		fmt.Println("value doesnt exist")
	}



}
// func divide(a, b int) (int, error) {
// 	if b==0{
// 		return 0, errors.New("divide by 0 illegal hell nawwwww")

// 	}
// 	return a/b, nil
// }

package main
import ("fmt",
	"strings")
type  Command struct{
	name string
	args []string
}
func readCmd(input string) (c Command){
	
	//cmd:=[]string
	cmd:=strings.Fields(input)
	//x:=0
	for i:=0; i<len(cmd);i++ {
		if(i==0){
				c.name=cmd[i]
			}else {
				c.args=append(c.args,cmd[i])
			}
		
	}
	return c
}
func implCmd(input Command, db Database) {
	if input.name=="SET"{
		if len(input.args)!=2 {
			fmt.Println("invalid")
			return
		}
		key:=input.args[0]
		val:=input.args[1]
		
		db.Set(key,val)

	}else if input.name=="GET"{
		if len(input.args)!=1 {
			fmt.Println("invalid")
			return
		}
		key:=input.args[0]
		val,e:=db.Get(key)
		if !e{
			fmt.Println("key dne")
		}else {
			fmt.Println(val)
		}
	}else if input.name=="DEL"{
		if len(input.args)!=1 {
			fmt.Println("invalid")
			return
		}
		key:=input.args[0]
		f:=db.Del(key)
		if f{
			fmt.Println("deleted")
		}else {
			fmt.Println("key dne")
		}
	}else {
		fmt.Println("invalid cmd")
		return 
	}


}
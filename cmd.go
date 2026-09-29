package main
import (
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
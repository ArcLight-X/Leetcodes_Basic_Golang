package main

import("fmt"
)

func isRobotBounded(instructions string) bool {
    a:=[2]int{0,0}
	d:=1
	for i:=0;i<10*len(instructions);i++{
	for _,v :=range instructions{
		switch v{
		case 'G':
			if d==1{//北
				a[1]=a[1]+1
			}else if d==2{//南
				a[1]=a[1]-1
			}else if d==3{//东
				a[0]=a[0]+1
			}else if d==4{//西
				a[0]=a[0]-1
			}
		case 'L':
			if d==1{
				d=4
			}else if d==2{
				d=3
			}else if d==3{
				d=1
			}else if d==4{
				d=2
			}
		case 'R':
			if d==1{
				d=3
			}else if d==2{
				d=4
			}else if d==3{
				d=2
			}else if d==4{
				d=1
			}
		}
	}
	if a[0]==0&&a[1]==0{
		return true
		break
	}
	}
	return false
}

func main(){
	fmt.Println(isRobotBounded("GG"))
}
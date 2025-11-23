package main

import (
	"fmt"
	"strings"
	//"structs"
)


func main(){
	






type desktop struct{
	name string
	process int8
	cpu float32
	memory int 
   temperature float32
}



kernel:=desktop {
	name:"kernel",
	process:1,
	cpu:2.6,
	memory:70,
   temperature:0.3 ,
}


 updatecheck:=desktop {
	name:"updatecheck",
	process:1,
	cpu:0.4,
	memory:25,
   temperature:1.1,
}

 gui:=desktop {
	name:"gui",
	process:1,
	cpu:0.4,
	memory:125,
   temperature:1.4,
}

firewall:=desktop {
	name:"firewall",
	process:1,
	cpu:5.8,
	memory:80,
   temperature:1.3,
}

browser:=desktop {
	name:"browser",
	process:1,
	cpu:5.8,
	memory:200,
   temperature:+3.5,
}

 cs2:=desktop {
	name:"cs2",
	process:1,
	cpu:30.58,
	memory:4000,
   temperature:12.5,
}

 spotify:=desktop {
	name:"spotify",
	process:1,
	cpu:6,
	memory:100,
   temperature:+1.5,
}
discord:=desktop {
	process:1,
	cpu:10,
	memory:120,
   temperature:+3.0,
}
network:=desktop {
	name:"network",
	process:1,
	cpu:4,
	memory:30,
   temperature:1.2,
}
smallprocess:=desktop {
	name:"smallprocess",
	process:1,
	cpu:0.2,
	memory:10,
   temperature:0.2,
}
smallprocess2:=desktop {
	name:"smallprocess2",
	process:1,
	cpu:0.2,
	memory:10,
   temperature:0.2,
}
miniprocess:=desktop {
	name:"miniprocess",
	process:1,
	cpu:0.22,
	memory:11,
   temperature:0.6,
}
powermanagment:=desktop {
	name:"powermanagment",
	process:1,
	cpu:4,
	memory:35,
   temperature:1,
}

var fan float32
var currently int
var processcount int8
var currentTemp float32
var cpucount float32
var  fanpercentage float32
var Temp float32
// fan affecting


ram:=(firewall.memory+powermanagment.memory+miniprocess.memory+smallprocess.memory+smallprocess2.memory+network.memory+gui.memory+kernel.memory)
	processcount=(firewall.process+powermanagment.process+miniprocess.process+smallprocess.process+smallprocess2.process+network.process+gui.process+kernel.process)
currentTemp=(firewall.temperature+powermanagment.temperature+miniprocess.temperature+smallprocess.temperature+smallprocess2.temperature+network.temperature+gui.temperature+kernel.temperature)
cpucount=(firewall.cpu+powermanagment.cpu+miniprocess.cpu+smallprocess.cpu+smallprocess2.cpu+network.cpu+gui.cpu+kernel.cpu)


//var tempchange float32
processcount2 :=[]string{}
notActive:=[]string{}
//var fanset *float32
//fanset=&fan
//alternative to print process names
processcount2=append(processcount2, firewall.name,powermanagment.name,miniprocess.name,smallprocess.name,smallprocess2.name,network.name,gui.name,kernel.name)
processcount3:=strings.Join(processcount2,"\n")
//var test string= firewall.name
// processcount2:=[]struct{firewall.name,powermanagment,miniprocess.name,smallprocess.name,smallprocess2.name,network.name,gui.name,kernel.name}
//up is current active process list

//not active bottom pccontroll==4


fanpercentage=fan/1000*100

Temp=currentTemp-fanpercentage
turnon:
fmt.Printf("turn on pc with 1\n")
fmt.Scan(&currently)



if currently==1{
fmt.Printf("Pc is on\n")
fmt.Printf(" %vprocesses are active,cpu:%v fan:%v temp:%v   Memory:%dMb   \n" ,processcount,cpucount,fan,Temp,ram)


} else {
	goto turnon
}
var pccontrol int8
start:

fmt.Println("1.home 2.process active list  3.fan control 4. newProgram ")
fmt.Scan(&pccontrol)

if pccontrol==1{
	fmt.Printf(" %vprocesses are active,cpu:%v fan:%v temp:%v   Memory:%dMb  \n",processcount,cpucount,fan,Temp,ram)

	goto start

}

if pccontrol==2{
//slice needed to be created
	
	// name var cannot be printed 
	// add to structs the proccess name also
	// good way to print inside values fmt.Print("proccess active: %#v ",processcount2) before 8 with this [{1 5.8 80 1.3} {1 4 35 1} {1 0.22 11 0.6} {1 0.2 10 0.2} {1 0.2 10 0.2} {1 4 30 1.2} {1 0.4 125 1.4} {1 2.6 70 0.3}]
	
	fmt.Println(processcount3)

	goto start

}

if pccontrol==3{
fmt.Println("set fan speed in %  max is 100%")

fmt.Scan(&fan)
fanpercentage=fan/1000*100

Temp=currentTemp-fanpercentage

fmt.Printf(" %vprocesses are active,cpu:%v fan:%v temp:%v   Memory:%dMb  \n",processcount,cpucount,fan,Temp,ram)

goto start
}

if pccontrol==4{
	fmt.Println("add process")

//create 2 slices, 1 active 1 not active , if active remove from not active , if active to not active remove from active append to not active
	
	

	

notActive=append(notActive,updatecheck.name,browser.name,cs2.name,spotify.name,discord.name)
notActive:=strings.Join(notActive,"\n")
fmt.Printf(notActive)
// var processAdd string
fmt.Println("enter proccess name:")
//fmt.Scan($processAdd) remove //
/*
for i:=0;i<notActive.len;i++{
	if processAdd==notActive[i]{
		processcount2:=append(processcount2,notActive)
	}
}*/ // remove //
// a scan ,  remove process from not active , add to processcount2

}






}


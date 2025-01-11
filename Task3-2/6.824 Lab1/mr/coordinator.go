package mr

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
)

//call 函数通过1234端口传入args和reply的内存地址，调用rpcname（Coordinator.函数名），
//通过反射机制"远程"调用Coordinator的该函数，Coordinator通过内存地址读取入参写出结果。
//worker.go里面的Worker方法调用CallExample，先运行Coordinator，再运行worker，
//看看worker端打印返回来的经过Coordinator加工过的数字

//Master节点的RPC服务端，负责分配任务给worker节点，并监控worker节点的状态，当所有worker节点完成任务后，Master节点会汇总结果并返回给客户端。
//MapReduce的基本思路是启动一个coordinator分配多个worker做map任务

type Coordinator struct {
	// Your definitions here.
	Tasks []Task
	State State
	Mutex sync.Mutex // 锁
}

func MakeCoordinator(files []string, nReduce int) *Coordinator {
	if nReduce <= 0 {
		panic(fmt.Sprintf("nReduce must be positive, not %d", nReduce))
	}
	c := Coordinator{}
	for i := 0; i < nReduce; i++ {
		//对于每个文件，启动一个协程来处理
		go c.handler(files[i], i)
		if files[i] == "" {
			break
		}
	}
	c.server()
	return &c
}

func (c *Coordinator) handler(files string, nReduce int) error {
	//任务分配
	c.Tasks = append(c.Tasks, Task{files: files,TaskType: MapTask, Filename: files, TaskID: nReduce})
	c.State = Waiting
	return nil
}

func (c *Coordinator) Done() bool {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	if c.State == AllDone {
		fmt.Printf("All workers done\n")
		return true // 应该返回true，表示所有工作都已完成
	} else {
		return false
	}
}

// start a thread that listens for RPCs from worker.go
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

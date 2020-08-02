package main

import (
	"container/list"
	"time"
)

type Job func(interface{})

type TimeWheel struct{
	interval time.Duration
	ticker *time.Ticker
	slots []*list.List
	timer map[interface{}]int
	currentPos int
	slotNum int
	job  Job
	addTaskChannel chan Task
	removeTaskChannel chan interface{}
	stopChannel chan bool
}

type Task struct{
	delay time.Duration
	circle int
	key interface{}
	data interface{}
}

func New(interval time.Duration, slotNum int,job Job) *TimeWheel{
	if interval<=0 || slotNum<=0 || job==nil{
		return nil
	}
	tw:=&TimeWheel{
		interval:interval,
		slots: make([]*list.List,slotNum),
		slotNum: slotNum,
		timer:make(map[interface{}]int),
		currentPos: 0,
		job:job,
		addTaskChannel:make(chan Task),
		removeTaskChannel: make(chan interface{}),
		stopChannel:make(chan bool),
	}

	tw.initSlots()

	return tw
}

func (tw *TimeWheel) initSlots(){
	for i:=0;i<tw.slotNum;i++{
		tw.slots[i]=list.New()
	}
}

//Start 启动时间轮
func (tw *TimeWheel) Start(){
	tw.ticker=time.NewTicker(tw.interval)
	go tw.start()
}

//Stop 停止时间轮
func (tw *TimeWheel) Stop(){
	tw.stopChannel<-true
}

//AddTimer 添加定时器key为定时器唯一标识
func (tw *TimeWheel) AddTimer(delay time.Duration, key interface{},data interface{}){
	if delay<0{
		return
	}

	tw.addTaskChannel <- Task{delay:delay,key:key,data:data}
}

//RemoveTimer 删除定时器 key为添加定时器传递的定时器唯一标识
func (tw *TimeWheel) RemoveTimer(key interface{}){
	if key==nil{
		return
	}

	tw.removeTaskChannel<-key
}

func (tw *TimeWheel) start(){
	for{
		select {
		case <- tw.ticker.C:
			tw.tickHandler()
		case task:=<- tw.addTaskChannel:
			tw.addTask(&task)
		case key:=<- tw.removeTaskChannel:
			tw.removeTask(key)
		case <- tw.stopChannel:
			tw.ticker.Stop()
			return
		}
	}
}

func (tw *TimeWheel) tickHandler(){
    l:=tw.slots[tw.currentPos]
    tw.scanAndRunTask(l)
    if tw.currentPos==tw.slotNum-1{
    	tw.currentPos=0
	}else{
		tw.currentPos+=1
	}
}

func (tw *TimeWheel) scanAndRunTask(l *list.List){
	for e:=l.Front();e!=nil;{
		task:=e.Value.(*Task)
		if task.circle>0{
			task.circle--
			e=e.Next()
			continue
		}
		go tw.job(task.data)
		next:=e.Next()
		l.Remove(e)
		if task.key!=nil{
			delete(tw.timer,task.key)
		}
		e=next
	}
}

//新增任务到链表中
func (tw *TimeWheel) addTask(task *Task){
	pos,circle:=tw.getPositionAndCircle(task.delay)
	task.circle=circle

	tw.slots[pos].PushBack(task)

	if task.key!=nil{
		tw.timer[task.key]=pos
	}
}

//获取定时器在槽中的位置  时间轮需要转动的圈数
func (tw *TimeWheel) getPositionAndCircle(d time.Duration)(pos,circle int){
	delaySeconds:=int(d.Seconds())
	intervalSeconds:=int(tw.interval.Seconds())

	circle=int(delaySeconds/intervalSeconds/tw.slotNum)
	pos=int(tw.currentPos+delaySeconds/intervalSeconds)%tw.slotNum

	return
}

//从链表中删除任务
func (tw *TimeWheel) removeTask(key interface{}){
	//获取定时器所在的槽为
	pos,ok:=tw.timer[key]
	if !ok{
		return
	}

	l:=tw.slots[pos]

	for e:=l.Front();e!=nil;{
		task:=e.Value.(*Task)
		if task.key==key{
			delete(tw.timer,key)
			l.Remove(e)
		}
		e=e.Next()
	}

}
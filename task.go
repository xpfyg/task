package main

import (
	"errors"
	"sync/atomic"
	"time"
)

type Task struct {
	bTaskExit          bool           //队列标准
	checkTimerInterval int            //定时检查的间隔
	queue              chan string    //队列
	queueMaxLen        int32          //队列长度
	curQueueLen        int32          //当前队列
	process            func(s string) //消费者
}

// Init 初始化队列任务
// 传入 时间间隔、队列长度、消费函数
func (this *Task) Init(checkTimerInterval, queueLen int, procees func(s string)) {
	this.checkTimerInterval = checkTimerInterval
	this.queueMaxLen = int32(queueLen)
	this.queue = make(chan string, this.queueMaxLen)
	this.process = procees
	return
}

// UnInit 关闭任务
func (this *Task) UnInit() {
	this.bTaskExit = true
	return
}

// PutQueue
func (this *Task) PutQueue(s string) error {
	after := time.NewTimer(time.Second * 2)
	defer after.Stop()

	select {
	case <-after.C: //超时处理， 因往chan中放入数据时，如果队列满了，则会阻塞  为防止将程序阻塞，故此处设置2秒超时，超时则本次操作失败，等待下次定时检查在执行
		return errors.New("put failed, timeout")
	case this.queue <- s:
		atomic.AddInt32(&this.curQueueLen, 1)
	}

	return nil
}
func (this *Task) Start(checkTimerInterval, queueLen int) {
	//go
	go this.runTask()
}

func (this *Task) runTask() {
	t := time.NewTimer(time.Second * time.Duration(2)) //每两秒钟唤醒一次，打印日志
	for !this.bTaskExit {
		select {
		case s := <-this.queue:
			this.process(s)
		case <-t.C:
			t.Reset(time.Second * time.Duration(2))
		}
	}
}

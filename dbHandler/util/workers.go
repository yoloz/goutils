package util

import (
	"context"
	"sync"
)

type TaskFuc func(args ...interface{})

type Task struct {
	f    TaskFuc
	args []interface{}
}

type WorkPool struct {
	pool        chan *Task
	workerCount int

	stopCtx       context.Context
	stopCancelFuc context.CancelFunc
	wg            sync.WaitGroup
}

func (t *Task) Execute() {
	t.f(t.args...)
}

func NewPool(workerCount, poolLen int) *WorkPool {
	return &WorkPool{
		workerCount: workerCount,
		pool:        make(chan *Task, poolLen),
	}
}

func (w *WorkPool) PushTask(t *Task) {
	w.pool <- t
}

func (w *WorkPool) PushTaskFunc(f TaskFuc, args ...interface{}) {
	w.pool <- &Task{
		f:    f,
		args: args,
	}
}

func (w *WorkPool) work() {
	for {
		select {
		case <-w.stopCtx.Done():
			w.wg.Done()
			return
		case t := <-w.pool:
			t.Execute()
		}
	}
}

func (w *WorkPool) Start() *WorkPool {
	w.wg.Add(w.workerCount)
	w.stopCtx, w.stopCancelFuc = context.WithCancel(context.Background())

	for i := 0; i < w.workerCount; i++ {
		go w.work()
	}
	return w
}

func (w *WorkPool) Stop() {
	w.stopCancelFuc()
	w.wg.Wait()
}

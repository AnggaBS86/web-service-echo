package queue

import (
	"context"
	"log"
	"sync"
)

type Queue struct {
	name    string
	jobs    chan Job
	context context.Context
	cancel  context.CancelFunc
}

type Job struct {
	Name   string
	Action func() error
}

type Worker struct {
	queue *Queue
}

func NewQueue(name string) *Queue {
	ctx, cancel := context.WithCancel(context.Background())

	return &Queue{
		jobs:    make(chan Job),
		name:    name,
		context: ctx,
		cancel:  cancel,
	}
}

func (q *Queue) AddJobs(jobs []Job) {
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, v := range jobs {
		go func(v Job) {
			q.AddJob(v)
			wg.Done()
		}(v)
	}

}

func (q *Queue) AddJob(job Job) {
	q.jobs <- job
	log.Printf("New job %s added to %s queue", job.Name, q.name)
}

func (j Job) Run() error {
	log.Printf("Job running: %s", j.Name)

	err := j.Action()
	if err != nil {
		return err
	}

	return nil
}

// NewWorker initialises new Worker.
func NewWorker(queue *Queue) *Worker {
	return &Worker{
		queue: queue,
	}
}

func (w *Worker) DoWork() bool {
	for {
		select {
		case <-w.queue.context.Done():
			log.Printf("Work done in queue %s: %s!", w.queue.name, w.queue.context.Err())
			return true
		case job := <-w.queue.jobs:
			err := job.Run()
			if err != nil {
				log.Print(err)
				continue
			}
		}
	}
}

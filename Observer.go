package main

import "fmt"

type Observer interface {
	HandleEvent(vacancies []string)
}

type Observable interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	sendAll()
}

type Person struct {
	name string
}

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

func (p *Person) HandleEvent(vacancies []string) {
	fmt.Println("Hi dear ", p.name)
	fmt.Println("Vacancies updated: ")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}
	fmt.Println()
}

func (j *JobSite) AddVacancies(vacancy string) {
	j.vacancies = append(j.vacancies, vacancy)
	j.sendAll()
}

func (j *JobSite) RemoveVacancy(vacancy string) {
	for i, k := range j.vacancies {
		if k == vacancy {
			j.vacancies = append(j.vacancies[:i], j.vacancies[i+1:]...)
		}
	}
	j.sendAll()
}

func (j *JobSite) subscribe(o Observer) {
	j.subscribers = append(j.subscribers, o)
}

func (j *JobSite) unsubscribe(o Observer) {
	for i, k := range j.subscribers {
		if k == o {
			j.subscribers = append(j.subscribers[:i], j.subscribers[i+1:]...)
		}
	}
}

func (j *JobSite) sendAll() {
	for _, subscriber := range j.subscribers {
		subscriber.HandleEvent(j.vacancies)
	}

}

func main() {
	hh := JobSite{subscribers: nil, vacancies: nil}
	fp := &Person{"Bob"}
	hh.AddVacancies("Senior HTML Developer")
	hh.subscribe(fp)
	hh.AddVacancies("Java Junior Developer")
	sp := &Person{"John"}
	hh.subscribe(sp)
	hh.AddVacancies("UI Designer")
	hh.RemoveVacancy("Senior HTML Developer")
	hh.sendAll()
}

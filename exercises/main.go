package main

func main() {
	job := map[int][]Job{
		2: {EmailJob{Message: "Welcome email"}, EmailJob{Message: "Newsletter"}},
		1: {DataJob{Data: "User data analysis"}},
		3: {ReportJob{Report: "Monthly report"}},
	}
	JobScheduler(job)
}

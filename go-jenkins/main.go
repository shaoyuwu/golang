package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/bndr/gojenkins"
)

func OneJobBuild(jenkins *gojenkins.Jenkins, ctx context.Context, jobname string) {
	jobarr := strings.Split(jobname, "/")
	newjobname := strings.Join(jobarr, "/job/") //newjobname: test-测试/job/job_test
	fmt.Printf("newjobname: %v\n", newjobname)
	_, err := jenkins.BuildJob(ctx, newjobname, nil)
	if err != nil {
		panic(err)
	}
}

func AllJobBuild(jenkins *gojenkins.Jenkins, ctx context.Context, jobname string) {
	_, err := jenkins.BuildJob(ctx, jobname, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	jenkins, _ := gojenkins.CreateJenkins(nil, "http://127.0.0.1:8080", "admin", "password").Init(ctx)

	gojenkins, err := jenkins.GetJob(ctx, "test-测试")
	if err != nil {
		panic(err)
	}
	jobList, err2 := gojenkins.GetInnerJobs(ctx)

	if err2 != nil {
		panic(err2)
	}
	for _, job := range jobList {
		afterJob := strings.TrimLeft(job.Base, "/job")

		//跳过某个job构建

		// jobSplit := strings.Split(afterJob, "/")
		// if jobSplit[len(jobSplit)-1] == "base" {
		// 	continue
		// }
		// if !strings.Contains(afterJob, "test1") || !strings.Contains(afterJob, "test2") {
		// 	continue
		// }

		AllJobBuild(jenkins, ctx, afterJob)

	}
	// OneJobBuild(jenkins, ctx, jobname)
}

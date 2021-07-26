[![Go](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/go.yml/badge.svg)](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/go.yml)       ![Cobra's_Greeting](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/greetings.yml/badge.svg) [![Run golangci-lint](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/lint.yml/badge.svg)](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/lint.yml) [![Mark stale issues and pull requests](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/stale.yml/badge.svg)](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/stale.yml)

# 100 Days of GO

Sharing my journey to improve my Go skills with 100 day focus. I hope you will learn, teach or contribute as we grow in the community together! 

### Day 1 (21 July)

This should take you 60 min, and you can expect to learn how to setup a Go development enviornment and basic Git Hub Actions.


## Exercise 1: Install GO and your choice of Integrated Development Enviornment (IDE)

**Time: 15 minutes**


* [Download Go](https://golang.org/)  
* [Download VS Code](https://code.visualstudio.com/download)

After you install both here is a screen shot of what it should look like after you run `go version` 

![CAFCFC15-0905-4511-80E6-89EE90DBE5CE](https://user-images.githubusercontent.com/46206055/126569451-381c19e0-e306-467d-94e0-9c9ca4224ffa.jpeg)
 


## To follow along follow these steps 

**Time: 10 minutes**

Clone the existing repositiory to have the hello world go application, and the yml files needed for the github action. 

``
git clone https://github.com/Cobra16319/100_Days_Of_Go.git
``  

Use git tags to iterate with the CI system and track any changes updating the yml file logically as needed.


``
git tag v1
``


``
git push --tags
``

### Day 2 of 100

The focus today is learning Goroutines, how to write files, fetching data from an api, wait group and channels. Take some time and look at the code and anaylis what is going on. I am cobra16319 on the Go Discord as well.   

This project will use the offical NHL Hockey Data API. 

I struggled to resolve an issue today with the "≠" versus the "!=" comparision operator.  I fixed the issue by resolving the below listed errors and learned something new about Go. Links below:

* [NHL API](https://statsapi.web.nhl.com/api/v1/teams)

* [JSON to GO](https://mholt.github.io/json-to-go/) - Credit to mholt  

Here is a photo of the converter that you should convert all of the data from the teams api from JSON to Go in: 

![22739310-8390-4E8B-94E7-412E7871521C_1_105_c](https://user-images.githubusercontent.com/46206055/126576973-81892e70-6fd3-4105-8d53-31e090e182fc.jpeg)


#### To Do for Day 2 

1. Get NHL API working locally with main.go
2. Figure out how to lint in go with Git Hub Actions
3. Select an Open Source security provider and integrate with Git Hub Actions


Here were some of the errors I recieved from the Comparison Operators

Sharing my current errors running a local git super lint tool locally which today I plan to write a custom git hub action for. The rest of day 2 will be fixing these errors and completing the second git hub action (Sharing learning experience) 


![590686C5-17A4-4E0B-8847-86A3B6F0261C_4_5005_c](https://user-images.githubusercontent.com/46206055/126795648-5cb8b235-a9c9-418a-b6ab-6874f3d7cdbc.jpeg)


### Acievments of the Day

 1. Installed 2 new workflows.
 2. Was able to get the basic application to run.
 3. Updated Documentation. 
 4. Accepted first pull request for markup in the read.md from a community member. 


### To Do for Day 3

1. Setup a way to provision infrastructure as code (IaC) with a budget constrant
2. Plan a way to recieve the data from the Restful API with something that has a GO SDK
3. Plan, apply and test the code on local stack. 

Photo of a correctly provisioned terraform package (opinionated): 

![14BC57B1-F764-48D2-B4DB-B1525B8F7FF1_4_5005_c](https://user-images.githubusercontent.com/46206055/126882115-4c5a4964-e0dd-4fc7-b381-ada4eb403ac6.jpeg)

Local Stack allows you to test Infrastructure of Code (IaC) with no cost. 

* [Local Stack] https://localstack.cloud/
 
Acievments of the Day

1. Provisioned 2 Terraform modules one remote and one local Resources: (AWS S3) (VPC) (API Gateway) (Lambda).
2. Confirmed functionality of terraform with local stack. 
3. Configured GO Lint test to do a static scan and failed; created the projects first issue and will address on day 4.
4. Configured Git Hub action to check for stale pull request to be inclusive of others helping. 


I am falling in love with Go! I need your help and support to survie the next 96 days

As you can see I do not have much experience with Go yet! My goal is to have Go be at the top of this list by the end of this project. 


![Your Repository's Stats](https://github-readme-stats.vercel.app/api/top-langs/?username=cobra16319&theme=blue-green)

![Profile View Counter](https://komarev.com/ghpvc/?username=cobra16319)



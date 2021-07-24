[![Go](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/go.yml/badge.svg)](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/go.yml) ![Cobra's_Greeting](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/greetings.yml/badge.svg) 

# 100 Days of GO

Open sourcing my journey to learn Go. I hope you will learn, teach or contribute as I grow in the community! 

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

I struggled to resolve an issue today with the "≠" versus the "!=" comparision operator.  I fixed the issue by resolving the below listed errors and learned something new about Go. Directly below is the medium article I learned from. 
* [Comparison_Operators] (https://medium.com/golangspec/comparison-operators-in-go-910d9d788ec0)

* [NHL API](https://statsapi.web.nhl.com/api/v1/teams)
* [JSON to GO](https://mholt.github.io/json-to-go/) - Credit to mholt  
* [YouTube] (https://www.youtube.com/watch?v=eu4wqYA7mBY) - Credit to EQuimper

Here is a photo of the converter that you should convert all of the data from the teams api from JSON to Go in: 

![22739310-8390-4E8B-94E7-412E7871521C_1_105_c](https://user-images.githubusercontent.com/46206055/126576973-81892e70-6fd3-4105-8d53-31e090e182fc.jpeg)


##### To Do for Day 2 

1. Get NHL API working locally with main.go
2. Figure out how to lint in go with Git Hub Actions
3. Select an Open Source security provider and integrate with Git Hub Actions


### Here were some of the errors I recieved from the Comparison Operators

Sharing my current errors running a local git super lint tool locally which today I plan to write a custom git hub action for. The rest of day 2 will be fixing these errors and completing the second git hub action (Sharing learning experience) 


![590686C5-17A4-4E0B-8847-86A3B6F0261C_4_5005_c](https://user-images.githubusercontent.com/46206055/126795648-5cb8b235-a9c9-418a-b6ab-6874f3d7cdbc.jpeg)


 ### Acievments of the Day

 1. Installed 2 new workflows.
 2. Was able to get the basic application to run.
 3. Updated Documentation. 
 4. Accepted first pull request for markup in the read.md from a community member. 

##### I am falling in love with Go! I need your help and support to survie the next 97 days

As you can see I do not have much experience with Go yet! My goal is to have Go be at the top of this list by the end of this project. 

### Day 3 

 
Challenge of the day is to write some code to provision the infrastructure to support writting data to (AWS S3). From here I plan to iterate and do something interesting with the data. 

1. Provisioned 2 Terraform modules one remote and one local. Resources: (AWS S3) (VPC) (API Gateway) (Lambda) 
2. Confirmed functionality of terraform with open stack. 
3. Configure terraform Git Hub action to check build with open stack and terraform cloud. 

* [Open Stack] https://localstack.cloud/

![Your Repository's Stats](https://github-readme-stats.vercel.app/api/top-langs/?username=cobra16319&theme=blue-green)

![Profile View Counter](https://komarev.com/ghpvc/?username=cobra16319)



[![Go](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/go.yml/badge.svg)](https://github.com/Cobra16319/100_Days_Of_Go/actions/workflows/go.yml)

# 100 Days of GO

This repository is a journal of my path to learning GO.
By the end of the 100 days, you should be able to follow along by day and learn Go as well.

### Day #1 21 July 2021 

This should take you 60 min, and you can expect to learn how to setup a Go development enviornment and basic Git Hub Actions.


## Exercise 1: Install GO and your choice of Integrated Development Enviornment (IDE)

**Time: 15 minutes**


* [Download GoLang][https://golang.org/]   
* [Download Vs Code][https://code.visualstudio.com/download]

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

### Day 2 of 100...

The focus today is learning Goroutines, how to write files, fetching data from an api, wait group and channels.  


## This project will use the offical NHL Hockey Data API. I am having issues getting the go.example.go coded so I will continue to work on that and write about it later. 


* [NHL Hockey API] [https://statsapi.web.nhl.com/api/v1/teams] 
* [JSON to GO converter] [https://mholt.github.io/json-to-go/] Credit to mholt for the converter. 


Here is a photo of the converter that you should convert all of the data from the teams api from JSON to Go in: 

![22739310-8390-4E8B-94E7-412E7871521C_1_105_c](https://user-images.githubusercontent.com/46206055/126576973-81892e70-6fd3-4105-8d53-31e090e182fc.jpeg)


## Going to add issue to the project for a linter and to try to figure out what GitHub action to use for scanning and code coverage. 




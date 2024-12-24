# Service Demo
[![CircleCI](https://circleci.com/gh/sasimpson/servicedemo/tree/master.svg?style=svg)](https://circleci.com/gh/sasimpson/servicedemo/tree/master)
![example workflow](https://github.com/sasimpson/servicedemo/actions/workflows/go.yml/badge.svg)

The purpose of this project is to demostrate how I build HTTP-based services.  I may be doing everything wrong, but this has seemed to work and I will explain how through some blog posts.  This project will show how to build these, from the ground up, use a database to store data, and include some deployment strategies.  

## Organization

### .circleci 
CircleCI is used to deal with CI/CD, is pretty easy to use, and has some nice features.  

### .github/workflows
GitHub Actions definition, will test and build.

### api
API is the directory that contains all the service endpoint implementations

### cmd
cmd contains the source for the binaries that will be generated.

### interfaces
this has the different implemenations of the model's interfaces

### models
describes how the data is going to look and the interfaces that need to be implemented.

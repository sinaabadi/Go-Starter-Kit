# Go Starter Kit
Bootstrap new project based on Gin framework. 

# Directory Structure
## apps
All apps are declared here for example `api v1` or `api v2` and etc.

## bin
Project start entry. only one main file should be here.

## config
Different config based on running `env` must be declared here. also `env.json` are overwrite all other config and `default.json`
file are used when needed config are not found in any other config files.

## constants
All project constants are stored in this folder.

## middlewares
All project middlewares stored in this file. please note that these middleware are all identical to gin middlewares.
### Authentication middleware
JWT authentication are configured. all you have to do is to define appropriate db source in middleware and define your own custom 
jwt schema  
## models
Project database models stored here. 

## routes
This folder store all project routes and you can split different route name space here for example `api/v1`

## utils
This directory serve as project helpers functions and logic. all project wide common and helper functions stored here.

# Go Module 
This pro
#Author
Sina Abadi  

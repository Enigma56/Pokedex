# Pokedex
## Overview
<p>
This project was created as part of my boot.dev education. Creating a pokedex and a way to interact with the pokemon universe is how they taught us how to interact with an API and handle JSON data in Golang.
</p>

## Skills Taught
1. JSON Marshal & Unmarshal in Golang
2. Creating a Webserver
3. Consuming 3rd-party API Endpoints
4. Creating basic CLI commands with Golang

## How to play with this project
<p> 
To interact with my code you just have to clone the repo, build the project and use the CLI commands I will list out below. 
</p>
<p> If that does not mean much to you, you can follow the directions below to do exactly this!</p>

### Step 1
<p>Install Golang version >= 1.23 using whichever terminal you typically use</p>

### Step 2
<p>Clone my github repo into a directory of your choosing and <code>cd</code> into it</p>

### Step 3
<p>Build and run this go project with one of the two following methods</p>

1. `go build -o pokedex . && ./pokedex`
2. `go run .`

> Note: The first option builds a binary then runs in. The second option does not build a binary rather compiles then runs the go project

### Step 4
You can now interact with the project using the CLI Commands listed below!

## CLI Commands
| Command Name | CLI Command |
| ------------ | ----------- |
| Help - List command descriptions| `help`|
| Exit - Exit the program | `exit` |
| Map - List next locations| `map` |
| Mapb - List previous locations | `mapb` |
| Explore - See pokemon in an area | `explore <location-name>` |
| Catch - Catch a pokemon| `catch <pokemon-name>` |
| Inspect - Inspect stats of a pokemon| `inspect <pokemon-name>` |
| Pokedex - See all pokemon you have caught| `pokedex` |

## Thank you!
I was a lot of fun to create this project. Following a guide allowed me to reference it when I struggled with learning a particular concept. All of the code written is my own and anything that resembles the code from the original project is an attempt at following best practices.
> You can see the project description and steps [here](https://www.boot.dev/courses/build-pokedex-cli)


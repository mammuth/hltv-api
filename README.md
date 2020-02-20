# HLTV API (inofficial)
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/mammuth/hltv-api)

### Features
- Get ical file for upcoming matches (filterable by teams)
- Json API for upcoming matches

![image](https://user-images.githubusercontent.com/3121306/74592853-84c90500-5025-11ea-9e88-d5d58f9fd20d.png)


### Usage
- Check out the testing deployment: https://mammuth-hltvapi.valar.dev/. I encourage you to deploy it on your own Heroku or Valar instance.
- You can compile it with `GOOS=linux GOARCH=amd64 go build -o main` and then manually run it with `./main`
- You could also use the button above to deploy it to Heroku

### Status
This is pretty unfinished and in a very early stage. 
Don't expect much enhancements in the near future. 
This is mainly because the underlying wrapper ([HLTV-Go](https://github.com/Olament/HLTV-Go)) doesn't support that many features yet.

### Attributions
This project uses [HLTV-Go](https://github.com/Olament/HLTV-Go).

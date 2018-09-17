Fetch Comments
==============

This is a simple Go script that fetched all comments from a Git repository to train an [ML model to detect typos](https://github.com/irinakhismatullina/style-analyzer/tree/feature/typos-analyzer/lookout/style/typos_checker) in code comments.
This project uses [bblfsh](https://bblf.sh/v2.6.1/) and [go-git](https://github.com/src-d/go-git) to do this.

## How to use
To use bblfsh you will need a server set up using Docker.  
`docker run -d --name bblfshd --privileged -p 9432:9432 -v /var/lib/bblfshd:/var/lib/bblfshd bblfsh/bblfshd`  
You will also need to have drivers for the languages you need to use. Since we want diverse data I suggest just installing all using:  
`docker exec -it bblfshd bblfshctl driver install --all`  

Now that bblfshd is running you can start running this repo. 

Currently the repos to analyse are listed in `main.go`
```
var repos = map[string]string{
	// "file name": "git url"
	"freeCodeCamp.coment": "https://github.com/freeCodeCamp/freeCodeCamp",
	"vue.coment":          "https://github.com/vuejs/vue",
	"springboot.coment":   "https://github.com/spring-projects/spring-boot",
	"moby.coment":         "https://github.com/moby/moby",
}
```
(This should be improved)

This project is designed to be a starting point for your go web app.

The original incarnation was a set of boilerplate I ended up useing in all my go projects.

## Usage

To see how it works:

    ./build.sh && ./webapptemplate

Then, simply browse to <http://localhost:8080> or on mac os x, run the command:

    open http://localhost:8080

## Starting a new web app project based on this template.

In practice, I normally just copy over all the files in this directory into a new directory then run the following:

	git init .
	git commit -am "Initial commit"

in github create the project then add the github origin and push to master.

    git remote add origin https://github.com/<USER>/<REPONAME>.git
    git push origin master


## TODO

- [ ] migrate from bindata to [packr](https://github.com/gobuffalo/packr)
    
# Password Checker

This program uses the HaveIBeenPwned password api to check if the password given has been leaked in any databreaches that they are aware of.  

It's really interesting how the api works, allowing you to check your password without giving it to anyone.  

I wrote a bunch of different versions of this program in [python](https://github.com/UnclassedPenguin/scripts/tree/master/passwordcheck), so when I wanted something to just play around with in Go, I decided to rewrite it. It really is basically the same program in either language, just slightly different syntax. But Go is my favorite lately, so why not rewrite it? 


## To use:

Make sure your go environment is set up, then run:

```shell
$ go install github.com/unclassedpenguin/passwordcheck@latest
```

Then just run it using `$ passwordcheck`. 

## Disclaimer

If you don't understand the source code enough to know whats going on, I wouldn't trust any program written by some random on the internet who asks for your password. :p

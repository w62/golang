[Writing Web Applications](https://go.dev/doc/articles/wiki/)

It has spent me quite a while to debug the following line

```
The bug version
var validPath = regexp.MustCompile("^/(edit|save|view)/(a-zA-Z0-9]+)$")
















The correct version
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
```

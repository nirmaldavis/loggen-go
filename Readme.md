# loggen-go
A Sample golang app to create logs for testing

```
go mod init github.com/nirmaldavis/mylogger-go

git init
```

Create new repository in github and use below

```
git remote add origin https://github.com/nirmaldavis/loggen-go.git

git add .

git commit -m "Initial commit"

git push --set-upstream origin master
```

May need to create Personal Access token in github 

Import needed additinal packages like -  import "rsc.io/quote" and do 
```
go mod tidy
```
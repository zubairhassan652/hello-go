# hello-go

This is test repo for working with heroku and golang

### Step 1

Clone this repo

### Step 2

Login to heroku cli using:

``` 
heroku login
```

### Step 3

Create heroku app using:

```
heroku create
```
you can see heroku remote is added to your repo you can check it by 


``` 
git remote -v 
```

your app will be created using random name you can rename it if you want using

```
heroku apps:rename newname
```

### Step 4

Now deploy to heroku server using 

```
git push heroku master
```

### Step 5

Access the browser by hitting

```
https://<your-app-name>.herokuapp.com
```


Congrates you are all done!


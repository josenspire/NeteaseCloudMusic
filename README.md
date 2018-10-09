##### Netease Cloud Music

> Note: Build by Golang and this project just for study.


---

##### Note Commands


--- 

```
- bee run // start project server

- bee run -gendoc=true -downdoc=true  // Generate Swagger Doc

- docUrl: http://localhost:8090/swagger/

```

###### Admin monitoring
> app.conf:  

``` app.conf
# Open admin monitoring and setup host:port

EnableAdmin = true
AdminAddr = "localhost"
AdminPort = 8088

```

###### checking

> To check the number of api reqeusts
[http://localhost:8088/qps](http://localhost:8088/qps)

---


> ###### [Official Introduction](https://beego.me/docs/advantage/monitor.md)
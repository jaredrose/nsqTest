"# nsqTest" 
经过1天的学习，对nsq有了一定的认识，想通过一个简短的Demo记录下：

##  main.go  ##

1、先启动nsqd：nsqd是一个接收、排队、然后转发消息到客户端的进程，
nsqd会默认监听一个tcp端口(4150)和一个http端口(4151)

2、新建一个生产者(Producer),负责发送消息(默认Topic为test)

3、新建一个消费者(Consumer),订阅消息

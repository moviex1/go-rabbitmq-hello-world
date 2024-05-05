<h1>Golang rabbitmq "Hello world"<h1>

This project shows basic example work with queues in rabbitmq using golang

<h2>Description</h2>

The basic logic is placed in command directory. There are producer where we sent one message with "Hello world" plain text message, 
and also there is consumer which waits for message to be published to queue, and when it receives message from producer it logs to console message

<h2> How to run </h2>

Start consumer
<pre>
  go run command/consumer/main.go
</pre>

Publish message to queue
<pre>
  go run command/producer/main.go
</pre>

<h2>Rabbitmq docker</h2>

if you want to test it with rabbitmq you can use this command to run container with rabbimq, and rabbitmq management 
<pre>
  docker run --name test-rabbit-mq -p 5672:5672 -p 15672:15672 -d rabbitmq:management-alpine
</pre>

and when you finish work with rabbit use this command to stop and then delete container
<pre>
  docker stop test-rabbit-mq && docker rm test-rabbit-mq
</pre>

(optional) you can delete image
<pre>
  docker rmi rabbitmq:management-alpine
</pre>

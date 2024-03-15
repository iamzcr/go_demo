package main

//select多路复用
//在某些场景，我们需要同时从多个管道接受数据，这个时候就可以用到golang中提供的select多路复用，，通常情况，管道在接受数据时候，如果没有数据可以接受，将发生阻塞
//select的使用类似与switch语句，有一系列的case分支和一个默认分支，每个case会对应一个管道的通信（接受或者发送）过程，
//select会一直等待，知道某个case通信操作完成，就会执行分支对应的语句。
//使用select获取channel里面的数据的时候不需要关闭channel




with open("test.txt" , "rb+") as f: 
    f.write("hello world!")

#二进制方式打开文件 ， 使用seek命令  ， 将文件更改
with open("test.txt" , "rb+") as f:
    f.seek(5)
    f.write("a")

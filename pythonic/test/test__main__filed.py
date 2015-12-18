#coding=utf-8


"""原因想写一个工具类 ， 直接调用python 函数， 通过函数名称 ， 直接调用类内方法
    __main__    可以实现这个想法

"""



import __main__





def p():
    print 'test'




setattr(__main__ , 'x' , 0)

getattr(__main__ , 'p')()
print x

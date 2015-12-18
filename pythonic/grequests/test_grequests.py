#coding=utf-8



import grequests


urls = [
        'http://www.baidu.com/'
    ]


#create get requests , get to see requests how to use
res = [ grequests.get(link) for link in urls]
#
for response in grequests.map(res )
    print response.text

#coding=utf-8


import requests
import sys
reload(sys)
sys.setdefaultencoding("utf-8")
import re 



extract_reg = re.compile("class=\"(brand_[0-9]+)\">(.*?)>([^<]+)<\/a><\/h4></li>")

for i in range(16):
    start_index = 10 * (i - 1)
    url = "http://m.2sc.sohu.com/wap/brandSelect/?startIndex=%s&province=bj&limit=10&qq-pf-to=pcqq.c2c" % start_index
    html = requests.get(url).text
    for info in extract_reg.finditer(html):
        print info.group(1) , info.group(3)
    

#coding=utf-8


from b2 import file2 





file2.mkdir_p(sys.argv[])

class Db(object):


    def __init__(self , db ):
        file2.mkdir_p(db) 
    


    def save(key , value):
        self.
import bisect 
import collections 
class _IndexObject(object):
    """index_file每个index存储结构体：
        key 及key对应的存储文件 以及文件指针
    """


    def __init__(self , key , index_file , seek):
        self.key = key 
        self.index_file = index_file 
        self.seek = seek 

class LoadIndex(object):
    """数据库中每个block下有个index.file ;index.file文件中存储抽样有序key；
        将匹配上的key返回 ，如果没有匹配文件则返回
    """

    def __init__(self , index_file , index_type):
        self._indexs = collections.OrderedDict()
        self.open_index(index_file)
    
    def open_index(self , index_file):
        with open(index_file) as f:
            for line in f.readlines():
                key , index_file , seek = line.rstrip().split()
                self._indexs[key] = _IndexObject(key , index_file , seek )

    def find(self ,key):
        seek_info = bisect.bisect_left(self.__index , key)
        if seek_info.key != key:
            
        
    

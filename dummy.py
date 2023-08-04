# coding: utf8

from locust import User, task

class Dummy(User):

    @task
    def hello(self):
        pass

# locust --master -f dummy.py --master --master-bind-host=0.0.0.0 --master-bind-port=5557 &

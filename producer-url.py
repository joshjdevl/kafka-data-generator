from faker import Factory
faker = Factory.create()

from kafka.client import KafkaClient
from kafka.producer import SimpleProducer

kafka = KafkaClient("149.204.61.37:49160")
producer = SimpleProducer(kafka)

while True:
    url = faker.url()
    producer.send_messages("topic", str(url))
    print url

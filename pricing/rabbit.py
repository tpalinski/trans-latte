from collections.abc import Callable
from typing import Any
import pika
import os
import time


connection = None
EXCHANGE_NAME = "updates"
TIMEOUT = 5
RETRIES = 5

def initialize_rabbit_connection(consume_callback: Callable[[Any, Any, Any, Any], None]):
    rmq_host = os.environ.get("RABBITMQ_ADDRESS", "localhost")
    for _ in range(RETRIES):
        try:
            global connection
            connection = pika.BlockingConnection(pika.ConnectionParameters(host=rmq_host))
        except:
            print("Could not connect to rmq host, retrying")
            time.sleep(TIMEOUT)
    channel = connection.channel()


    result = channel.queue_declare(queue='', exclusive=True)
    queue_name = result.method.queue
    exchange = channel.exchange_declare(EXCHANGE_NAME, "topic")

    # TODO - actually do it based on protobuf to string
    channel.queue_bind(exchange=EXCHANGE_NAME, queue=queue_name, routing_key="order.*")

    # Output queue
    channel.exchange_declare("pricing", "direct")
    channel.queue_declare(queue="pricing")
    channel.queue_bind(exchange='pricing', queue="pricing", routing_key="pricing")

    print('Connected to rmq instance')
    channel.basic_consume(queue=queue_name, on_message_callback=consume_callback, consumer_tag="pricing", auto_ack=True)

    channel.start_consuming()


def send_order_update(order: Any):
    send_chan = connection.channel()
    send_chan.basic_publish(exchange='pricing', routing_key='pricing', body=order)


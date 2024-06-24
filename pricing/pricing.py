#!/bin/python

import rabbit
import processing
import files

def main():
    print("Starting pricing service")
    files.initialize_minio_client()
    rabbit.initialize_rabbit_connection(consume_callback=processing.process_rmq_request)


if __name__ == "__main__":
    main()

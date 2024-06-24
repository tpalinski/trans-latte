import os
from minio import Minio
import time

client: Minio|None = None

RETRIES = 10
TIMEOUT = 5
HANDSHAKE_TIMEOUT = 0.5
BUCKET_NAME = "document-order-data" 

def initialize_minio_client():
    minio_host = os.environ.get("MINIO_ADDRESS", "localhost")
    access_key = os.environ.get("MINIO_ACCESS_KEY", "")
    secret_key = os.environ.get("MINIO_SECRET_KEY", "")
    for _ in range(RETRIES):
        try:
            global client
            client = Minio(minio_host, access_key=access_key, secret_key=secret_key, secure=False)
        except:
            print("Could not connect to minio, retrying")
            time.sleep(TIMEOUT)
    print("Successfully connected to minio instance")

# Download minio file and return path where it is stored
def get_minio_file(file: str, path: str = "/tmp/") -> str:
    # Handshake getting object metadata
    for i in range(RETRIES):
        try:
            _ = client.stat_object(BUCKET_NAME, object_name=file)
        except:
            print("Handshake with minio failed")
            if i == RETRIES-1:
                return ""
            else:
                time.sleep(HANDSHAKE_TIMEOUT)
    file_path = path+file
    _ = client.fget_object(BUCKET_NAME, object_name=file, file_path=file_path)
    return file_path



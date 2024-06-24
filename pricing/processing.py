import os
from typing import Any
import crud_pb2
import files
import rabbit
from pypdf import PdfReader
import time

ARTIFICIAL_DELAY = os.environ.get("DELAY", "0")

def process_rmq_request(ch: Any, method: Any, properties: Any, body: Any) -> None:
    payload = crud_pb2.OrderStatusInfo()
    payload.ParseFromString(body)
    print(f"Processing order: {payload.id}")
    path = payload.id + ".pdf"
    print(f"Downloading file: {path}")
    downloaded_path = files.get_minio_file(path)
    print(f"Downloaded file @ {downloaded_path}")
    if downloaded_path == "":
        print("There was an error while downloading document")
        return 
    price = calculate_document_price(downloaded_path)
    payload.price = price
    body = payload.SerializeToString()
    time.sleep(int(ARTIFICIAL_DELAY))
    rabbit.send_order_update(body)
    print(f"Successfully processed order {payload.id}")




# That would actually be more sophisticated in real life application. for now, just multiply words
def calculate_document_price(file_path: str) -> int:
    reader = PdfReader(file_path)
    pages = len(reader.pages)
    return pages * 30


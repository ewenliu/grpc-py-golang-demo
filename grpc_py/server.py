import grpc
from app.grpc.Product_pb2_grpc import ProdServiceServicer, add_ProdServiceServicer_to_server
from Product_pb2 import ProductResponse
from concurrent import futures
import time

class ProductService(ProdServiceServicer):
    def GetProductStock(self, request, context):
        print(request)
        return ProductResponse(ProdStock=123)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_ProdServiceServicer_to_server(ProductService(), server)
    server.add_insecure_port('127.0.0.1:8083')
    server.start()
    try:
        print('server started!')
        while True:
            time.sleep(3600*24)
    except KeyboardInterrupt:
        server.stop(0)


serve()
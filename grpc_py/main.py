import grpc
from app.grpc.Product_pb2_grpc import ProdServiceStub
from app.grpc.Product_pb2 import ProductRequest

# python -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. Product.proto

def run():
    with grpc.insecure_channel('127.0.0.1:8083') as channel:
        stub = ProdServiceStub(channel)
        resp = stub.GetProductStock(ProductRequest(ProdId=123))
    print(resp)

if __name__ == '__main__':
    run()
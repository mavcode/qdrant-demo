import sys
import argparse
from qdrant_client import QdrantClient
from sentence_transformers import SentenceTransformer, util
from PIL import Image

import numpy as np
import matplotlib.pyplot as plt


def client(url, key):
    print(f"host: {url}, key: {key}")
    client: QdranClient

    if key:
        client = QdrantClient(url=url, api_key=key)
    else:
        client = QdrantClient(url=url)

    return client

def add_to_collection(client):
    raise  NotImplementedError("TODO")


def create_collection(client, coll, models):
    client.create_collection(
            collection_name = coll,
            vectors_config = client.get_fastembed_vector_params(),
            sparse_vectors_config = client.get_fastembed_sparse_vector_params(),
    )

def hybrid_search(client, coll, text, limit):
    results = client.query(
            collection_name = coll,
            query_text = text,
            query_filter = None,
            limit = limit,
            )
    metadata = [i.metadata for i in results]
    return metadata

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("-u", "--url", help="qdrant db endpoint url", required=True)
    parser.add_argument("-k", "--key", help="api key", required=False)
    parser.add_argument("-s", "--search", help="hybrid search: <collection> <'search text'> <number of returns> ", nargs=3)
    parser.add_argument("-m", "--models", help="<dense model> <sparse model>", nargs=2)
    parser.add_argument("-c", "--create", help="create collection: <collection name>")
    parser.add_argument("-l", "--list", help="list collection: <collection name>")
    parser.add_argument("-a", "--add", help="add data to collection: <collection name> <input file>", nargs=2)
    args = parser.parse_args()

    _client = client(args.url, args.key)

    if args.list:
        results = _client.get_collection(args.list)
        print(results)
        sys.exit()

    if args.models:
        print(f"models: {args.models}")
        _client.set_model(args.models[0])
        _client.set_sparse_model(args.models[1])
    else:
        _client.set_model("sentence-transformers/all-MiniLM-L6-v2")
        _client.set_sparse_model("prithivida/Splade_PP_en_v1")

    if args.create:
        if not _client.collection_exists(args.create):
            create_collection(_client, args.create)
        else:
            sys.exit("collection already exists")

    if args.search:
        print(f"text to do search with:  {args.search}")
        results = hybrid_search(_client, args.search[0], args.search[1], int(args.search[2]))
        print(results)

if __name__ == "__main__":
    main()

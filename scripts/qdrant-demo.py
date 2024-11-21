from datasets import load_dataset

dataset = load_dataset("Qdrant/arxiv-titles-instructorxl-embeddings")

from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://node1.mavcode.net:6333", api_key="05dffa8ae687a22345c6b3e1f038b1540302e135e04ec4c7b08b7ade5e88caf3")


client.update_collection(
    collection_name="arxiv-titles-instructorxl-embeddings",
    optimizer_config = models.OptimizersConfigDiff(indexing_threshold=0),
)

from itertools import islice, batched

batch_size = 100

for batch in batched(dataset, batch_size):
    ids = [point.pop("id") for point in batch]
    vectors = [point.pop("vector") for point in batch]

    client.upsert(
        collection_name="arxiv-titles-instructorxl-embeddings",
        points=models.Batch(
            ids=ids,
            vectors=vectors,
            payloads=batch,
        ),
    )

client.update_collection(
    collection_name="arxiv-titles-instructorxl-embeddings",
    optimizer_config = models.OptimizersConfigDiff(indexing_threshold=20000),
)


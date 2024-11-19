# qdrant-demo

## Assignment
- **Infra part.**  
Set up a Qdrant Hybrid Cloud deployment with a Kubernetes environment running on your local machine. Then, create a Qdrant cluster in it with two nodes and configure an API key for authentication.

- **Search part.**  
Implement a working Hybrid Search (with Sparse and Dense vectors) example app with Qdrant using Python. You can use any dataset and any embedding models. No UI is needed, just the code. That's it.
Everything you need to complete the task you should be able to find in Qdrant documentation.
After you are done, push the source code to a public Git repository and send us a link for review.


## Steps Used for this Demo
1. Create empty repo
2. Create go mod
3. Create Angular proj
4. Add cobra support to main app
5. Add init backend package files to main app
6. Created Free-Tier Qdrant online account
7. Generated API Key
8. Accessed remote db via managed services
9. Able to connect to managed db with go-client running on local system
10. Create static provisioning of PV called qdrant-storage on local K8 cluster
11. Edit the qdrant pvc and modified the 'storageClassName' and 'volumeName' attributes
12. On 2 different nodes, I installed qdrant source and built code using cargo
13. Running 2 instance in debug mode and will modify the config file to ensure p2p communication, then replicate it on K8 instances
14. 2 nodes were configured with p2p functionality
15. Verified they are running by executing commands:
    curl auria-20.attlocal.net:6333/cluster | jq  and,
    curl auria-10.attlocal.net:6333/cluster | jq.
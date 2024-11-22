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
15. Verified cluster were configured properly
16. Created a Private Cloud installation https://qdrant-demo.mavcode.io/dashboard using Helm on a local K8 and exposed it via Ingress with TLS enabled.
17. Created a Hybrid Cloud installation http://qdrant.mavcode.io:6333/dashboard on my local K8 and exposed it via LoadBalancer. The qdrant PV storage yaml files are in the ./deployments directory.
18. Used Python, Go, and Curl, to create collections and injest data to all three types of deployments
19. Reviewed the Hybrid Search online tutorials and studied the example codes
20. Wrote initial version of Go application to use for performant ingestions to qdrant db 
20. Wrote a simple python script to do hybrid search on the hybrid cloud cluster deployment using online example data

```
python scripts/hybrid-search.py -u <qdrant endpoint url> -k <api key> -m <dense model> <sparse model> -s <collection> <search text> <max number of return hits>

example:
python scripts/hybrid-search.py -u http://qdrant.mavcode.io:6333 -k <API KEY> -m sentence-transformers/all-MiniLM-L6-v2 prithivida/Splade_PP_en_v1 -s startups "Two dogs in the snow" 10
```

```output
[{'link': 'http://www.kumfytailz.com', 'document': 'Revolutionary Warming/Cooling Dog Apparel\nKumfyTailz combines well-known gel-pack technology (think human therapy applications, microwave or freeze it and put it on your knee or shoulder) with premium-quality dog harnesses and coats. We keep pets comfortable in any weather, for more outside play/walk ...', 'images': 'https://d1qb2nb5cznatu.cloudfront.net/startups/i/257434-d3fc861a4401e34169f3da806ef2192b-thumb_jpg.jpg?buster=1377799352', 'city': 'Chicago', 'name': 'Kumfy Tailz', 'alt': 'Kumfy Tailz -  pets'}, {'images': 'https://d1qb2nb5cznatu.cloudfront.net/startups/i/541010-fd7728034becdea19bef742b586def10-thumb_jpg.jpg?buster=1416685118', 'alt': 'PupJoy -  e-commerce pets subscription businesses Pet Care', 'name': 'PupJoy', 'link': 'http://www.pupjoy.com', 'city': 'Chicago', 'document': "Helping Dog Lovers Treat Better. Raising the Bar on the Subscription Box.\nFor your dogs' treat and toy needs, we are a better subscription box, focused on customization and outstanding service. When you want to outfit them with cool, quality accessories, visit our online boutique. We provide you with better products (like healthy, clear ..."}, {'city': 'Chicago', 'images': 'https://d1qb2nb5cznatu.cloudfront.net/startups/i/28002-cbe0f9a8c2bce0f11a83575cd002c305-thumb_jpg.jpg?buster=1421987208', 'link': 'http://www.doggyloot.com', 'document': 'Top Dog in Online Pet Product Discovery, acquired by FamilyPet in 2015\ndoggyloot is a super-simple eCommerce platform that sells curated pet products to a highly engaged customer base. We help dog-loving customers discover top-quality products every day at the most competitive prices with our daily emails, and make it convenient for ...', 'alt': 'Doggyloot -  e-commerce consumer goods retail pets', 'name': 'Doggyloot'}, {'city': 'Chicago', 'name': 'Betel Nut Games', 'document': 'We make mildly addicting and highly stimulating games you should play\nDog Park is an original arcade puzzler that puts you in the role of a variety of dogs attempting to become the top dog of a dog park. With each park you dominate, you add more dogs to your pack, each with their own style of play.\nMore original games on the way ...', 'link': 'http://www.betelnutgames.com/', 'alt': 'Betel Nut Games - ', 'images': 'https://d1qb2nb5cznatu.cloudfront.net/startups/i/85473-5d21b6faf1f45067c1a3a45e7c75f365-thumb_jpg.jpg?buster=1338998117'}, {'images': 'https://d1qb2nb5cznatu.cloudfront.net/startups/i/317671-cc690804597820d28b0a14de5c10e206-thumb_jpg.jpg?buster=1388880146', 'alt': 'Worthee -  mobile services pets', 'link': 'http://worthee.com', 'name': 'Worthee', 'document': 'Premium pet services, on demand.\nPremium dog walking and sitting services – book on demand or schedule in advance, 24/7x365. Passionately serving dogs and their owners in metro Chicago since 2013.', 'city': 'Chicago'}, {'name': 'Bowzer Buddy', 'link': 'http://www.bowzerbuddy.com', 'images': 'https://d1qb2nb5cznatu.cloudfront.net/startups/i/23310-67f18b40ecac0dc0c92938216c4fc9c3-thumb_jpg.jpg?buster=1315730092', 'city': 'Chicago', 'document': 'One-handed biodegradible dog waste pick-up bag & convenient carryall\nThe innovative BowZer Buddy carryall makes dog walking easy & convenient with all your dog walking needs in one place. The carryall is always ready to go when your dog is & carries everything necessary to make walking your dog a pleasant experience (ID, keys, phone, ...', 'alt': 'Bowzer Buddy -  pets'}, {'city': 'Chicago', 'link': 'http://None', 'document': 'Snowboard/Ski Contest Events\nsnowboarding/skiing competitions 12 months a year. looking to take snow sports to the next level', 'alt': 'World Snow Tour -  television events', 'name': 'World Snow Tour', 'images': 'https://d1qb2nb5cznatu.cloudfront.net/startups/i/160681-b1bc028083b8517223ff6324e883a4ed-thumb_jpg.jpg?buster=1359242185'}, {'alt': 'Urban Leash -  mobile pets Consumer ', 'name': 'Urban Leash', 'images': 'https://d1qb2nb5cznatu.cloudfront.net/startups/i/299934-e45106a41c34392281a29cc4448ed819-thumb_jpg.jpg?buster=1406672135', 'link': 'http://www.urbanleash.com', 'document': 'Uber for dog walking\nUrban Leash is like Uber for dog walking. Through our technology platform, we provide a convenient way for pet owners to order last minute and ongoing dog-walking and cat-sitting services from anywhere, at anytime. With instant check-in/out notifications, live ...', 'city': 'Chicago'}, {'document': 'Innovative DayCare for Your Dog is Finally Here\nOur product is our service. We are starting a new kind of "daily care" for your dog. We are a dog daycare in the Chicago area who also offers dog walking and pet sitting. While these services are not new - we provide education throughout the process. You should ...', 'name': 'RedRoverDogs', 'images': 'https://d1qb2nb5cznatu.cloudfront.net/startups/i/416645-0f0b7af50a6fd1ab93c84ef09ef7ac94-thumb_jpg.jpg?buster=1402515822', 'alt': 'RedRoverDogs -  pet sitting Pet Care', 'link': 'http://redroverdogs.com', 'city': 'Chicago'}, {'name': 'Laboratory Analytics', 'alt': 'Laboratory Analytics -  mobile enterprise software health care information technology', 'document': 'Google docs for laboratory mice\nMouseHouse is using technology to transform all aspects of animal management in laboratories, from breeding and experiments to purchasing and veterinary management. In the US, scientists use 155M mice and spend $8.5B/yr housing and managing mice. MouseHouse is ...', 'link': 'http://www.mousehouseapp.com', 'city': 'Chicago', 'images': 'https://d1qb2nb5cznatu.cloudfront.net/startups/i/304730-cbcb42741a805026a485d772ce2f8c13-thumb_jpg.jpg?buster=1386347368'}]
```

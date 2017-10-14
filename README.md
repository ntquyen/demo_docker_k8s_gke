# demo_docker


## Local Development With docker & docker-compose
Why dev would want to use docker & docker-compose in their local development?
- Simple, unified interface to manage all kind applications
- Depenency management is easy, support nested dependencies
- One command line to spin up the whole dev environment, ready for code

## Scale Docker with Kubernetes & Google Cloud Platform

### Create a k8s cluster on Google Container Engine
Single command:

```
gcloud container clusters create test-cluster \
      --zone="asia-southeast1-a" \
      --machine-type="n1-standard-1" \
      --num-nodes=1 \
      --disk-size=30
```

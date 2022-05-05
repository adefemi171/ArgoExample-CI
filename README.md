# Web App



# App Details
A sample appliccation to demo CI/CD of an application with autoscaling using CirlceCI, ArgoCD and Kubernetes

.

# Proposed Stack

Golang


# Top-level directory layout

    📦netguruTask
        📦.circleci
            ┗ 📜config.yml
        ┣ 📜deployment.yaml
        ┣ 📜.gitignore
        ┣ 📜Dockerfile
        ┣ 📜go.mod
        ┣ 📜main_test.go
        ┣ 📜main.go
        ┣ 📜README.md


## Folder structure

1. [.circleci](https://github.com/adefemi171/netguruTask/tree/main/.circleci): Contains configuration for deployment

# How to setup project and run locally

### Fork the repository 

```
git clone https://github.com/{github_username}/netguruTask.git
```
### Checking Out
The App is deployed on the ``` main ``` branch you will need to checkout to the app branch using:

```
git checkout main
```

### Connect your repository to CircleCI
- To run the CI part of the application you will need to link your repository with [CircleCI](https://circleci.com/) by [creating an account and connecting your repo](https://circleci.com/signup/). 
- To add the environment variables. Click on Dashboard, Select your pipeline, on the far right of your screen go to Project Settings -> Environment variables. Add `DOCKER_PASSWORD`, `DOCKER_USERNAME` and `GITHUB_PERSONAL_TOKEN`.
- Update the [.circleci/config.yml](https://github.com/adefemi171/netguruTask/blob/main/.circleci/config.yml) from line [32](https://github.com/adefemi171/netguruTask/blob/291e74a727be999407c83501d8b78640473802b8/.circleci/config.yml#L32) to line [53](https://github.com/adefemi171/netguruTask/blob/291e74a727be999407c83501d8b78640473802b8/.circleci/config.yml#L53) to reflect yur changes and your docker image.
- Update [deployment.yaml]() on [line 39](https://github.com/adefemi171/netguruTask/blob/291e74a727be999407c83501d8b78640473802b8/deployment.yaml#L39) of the CI part and [line 39](https://github.com/adefemi171/netguruTask-CD/blob/7e3b5cad4297a2d579a11735e2fa6d71b965de4c/deployment.yaml#L39) of the CD task to reflect your container image

### Run the application

1. Run the the application using
```
go run main.go
```
3. In your browser, navigate to

```
http://localhost:3000/
```



# How to setup project using docker

1. 2.  Create the Docker image
    ```
    docker build -t your-username/appdemo:v1 .
    ```
    Example
    ```
    docker build -t adefemi171/appdemo:v1 .
    ```

3.  Run the docker file
    ```
    docker run -p 3000:3000 --name appdemo: YOURNAME/appdemo:v1
    ```
    Example
    ```
    docker run -p 3000:3000 --name appdemo adefemi171/appdemo:v1
    ```

4.  To confirm if your container was built, run:
    ```
    docker ps -a
    ```

5.  View the image in your localhost port 8888 by typing ```localhost:3000``` in your browser

# How to run application using ArgoCD
1. Once you have successfully installed ArgoCD and all neccessary components
2. Link your projects with the repository as stated [here](https://docs.google.com/document/d/1WUShzoZzdikpp9zpARXPzhCYHuZ8q5Yfq6Bxdd0CR-8/edit?usp=sharing)
3. Create a new application as stated [here](https://docs.google.com/document/d/1WUShzoZzdikpp9zpARXPzhCYHuZ8q5Yfq6Bxdd0CR-8/edit?usp=sharing)
4. Once your application is created and synced, check to see if you have your pods running by using the command
```
kubectl get po --all-namespaces
```
You should see a pod called app-demo-deployment-****
5. Get the current svc that the application is running on by using
```
kubectl get svc --all-namespaces
```
6. Do a port forward to be able to access the container o your browser, run
```
kubectl port-forward svc/app-demo -n app-demo 3000:3000
```
And visit the endpoint in your browser  as specified above

Note: app-demo can be any name you gave your service in the [deployment.yaml](https://github.com/adefemi171/netguruTask/blob/main/deployment.yaml) file.

NOTE: [Docker](https://docs.docker.com/get-docker/), [Golang](https://golang.org/doc/install), and [Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/) needs to be installed to run this application
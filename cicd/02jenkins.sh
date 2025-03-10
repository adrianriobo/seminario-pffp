podman container run \
        --name jenkins --rm --detach \
        --privileged \
        --publish 8080:8080 \
        --publish 50000:50000 \
        --volume jenkins-data:/var/jenkins_home \
        --volume jenkins-docker-certs:/certs/client:ro \
        docker.io/jenkins/jenkins:2.500-jdk21
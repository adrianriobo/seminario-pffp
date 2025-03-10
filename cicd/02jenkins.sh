podman container run \
        --name jenkins --rm --detach \
        --privileged \
        --publish 8080:8080 \
        --publish 50000:50000 \
        --volume jenkins-data:/var/jenkins_home \
        --volume jenkins-docker-certs:/certs/client:ro \
        docker.io/jenkins/jenkins:2.500-jdk21

podman exec -ti jenkins cat /var/jenkins_home/secrets/initialAdminPassword


podman run -d --name jenkins-agent \
  -v /var/run/docker.sock:/var/run/docker.sock \
  docker.io/jenkins/inbound-agent:latest \
  -url http://host.containers.internal:8080 \
  76a1ea5e89c8a9df190a744afe0c6e7e29a60ec92d31ed3e740c240a4c49e463 \
  test
  
podman run \
        --name jenkins --rm --detach \
        --privileged \
        --publish 8080:8080 \
        --publish 50000:50000 \
        --volume jenkins-data:/var/jenkins_home \
        --volume jenkins-docker-certs:/certs/client:ro \
        docker.io/jenkins/jenkins:2.500-jdk21

podman exec -ti jenkins cat /var/jenkins_home/secrets/initialAdminPassword

# of executors: 1 (or more).
Remote root directory: /home/jenkins/agent
Launch method: Launch agent by connecting it to the controller

build container first

podman run -d --name jenkins-agent \
  localhost/agent \
  -url http://host.containers.internal:8080 \
  0e85beed918f12a11a164227623511def034b7c6e634fcd41d17f19de2271804 \
  test
  
  
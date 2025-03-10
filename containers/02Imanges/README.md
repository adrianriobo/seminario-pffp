1. Construir imagenes

podman build -t image-1 -f 02Imanges/01Containerfile.yaml

.... Trying to pull registry.access.redhat.com/ubi9:latest...

skopeo inspect containers-storage:localhost/image-1

2. Como arrancar y probar...

podman run -it --rm fedora

cat /etc/*release

dnf install -y cowsay

...

3. Ahora lo movemos a un Containerfile

podman build -t cowsay -f 02Imanges/02Containerfile.yaml

podman run --rm cowsay:latest

podman run --rm cowsay:latest muuuu

podman run --rm --arch arm64 localhost/cowsay arch

4. Construir para multi-arch

podman manifest create cowsay

podman build --platform linux/arm64 --manifest cowsay -f 02Imanges/02Containerfile.yaml

...Trying to pull registry.fedoraproject.org/fedora:latest...
 Fedora 41 openh264 (From Cisco) - aarc 100% | 477.0   B/s |   6.0 KiB |  00m13s
 Fedora 41 - aarch64                    100% | 243.2 KiB/s |  34.2 MiB |  02m24s
...

podman manifest inspect cowsay

podman build --platform linux/amd64 --manifest cowsay -f 02Imanges/02Containerfile.yaml

podman manifest inspect cowsay

podman run --arch arm64 localhost/cowsay muu

podman run --arch amd64 localhost/cowsay muu

podman run -it --arch arm64 --entrypoint bash localhost/cowsay 

podman run -it --arch amd64 --entrypoint bash localhost/cowsay 

5. Desplegar pod en podman 

podman kube play 02Imanges/03cowpod.yaml

podman kube down 02Imanges/03cowpod.yaml
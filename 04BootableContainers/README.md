1. Build 

sudo podman build -t cowsay-bootc -f 04BootableContainers/01Containerfile

sudo podman run -it --rm localhost/cowsay-bootc:latest

sudo podman run --rm -it --privileged --pull=newer --security-opt label=type:unconfined_t \
          -v ./config.toml:/config.toml:ro \
          -v ./output:/output \
          -v /var/lib/containers/storage:/var/lib/containers/storage \
          quay.io/centos-bootc/bootc-image-builder:latest \
          --target-arch amd64 \
          --type qcow2 \
          --rootfs xfs \
          --local \
          localhost/cowsay-bootc:latest
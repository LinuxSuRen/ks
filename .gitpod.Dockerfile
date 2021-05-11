FROM gitpod/workspace-full

ENV PATH /home/linuxbrew/.linuxbrew/bin:$PATH

# More information: https://www.gitpod.io/docs/config-docker/
RUN sudo rm -rf /usr/bin/hd && \
    brew install linuxsuren/linuxsuren/hd && \
    hd install kubesphere/kubekey

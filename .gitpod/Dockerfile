FROM gitpod/workspace-full:latest

RUN curl -fsSL https://get.pulumi.com | sh -s -- --version 3.16.0 && \
    sudo -H /usr/bin/pip3 install pulumi_docker pylint black

RUN sed -i 's/export PIP_USER=yes/export PIP_USER=no/g' /home/gitpod/.bashrc

RUN wget -O pulumictl.tar.gz $(curl -s https://api.github.com/repos/pulumi/pulumictl/releases/latest | jq -rc '.assets | map(select(.name | contains("linux-amd64"))) | .[] .browser_download_url')
RUN tar xzf pulumictl.tar.gz -C /home/gitpod/.pulumi/bin pulumictl
RUN rm pulumictl.tar.gz

ENV PATH="${PATH}:/home/gitpod/.pulumi/bin"
ENV PIP_USER=no
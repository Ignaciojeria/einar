# You can find the new timestamped tags here: https://hub.docker.com/r/gitpod/workspace-base/tags
FROM gitpod/workspace-base:latest

# Change your version here
ENV GO_VERSION=1.20

# For ref, see: https://github.com/gitpod-io/workspace-images/blob/61df77aad71689504112e1087bb7e26d45a43d10/chunks/lang-go/Dockerfile#L10
ENV GOPATH=$HOME/go-packages
ENV GOROOT=$HOME/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin
RUN curl -fsSL https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz | tar xzs \
    && printf '%s\n' 'export GOPATH=/workspace/go' \
                      'export PATH=$PATH:$GOPATH/bin' > $HOME/.bashrc.d/300-go

# Clone the repository and build the binary
RUN mkdir -p /home/gitpod/bin && \
    git clone https://github.com/Ignaciojeria/einar-cli.git /home/gitpod/bin/einar-cli && \
    cd /home/gitpod/bin/einar-cli && \
    go build -o einar

# Move the binary to the /home/gitpod/bin directory
RUN mv /home/gitpod/bin/einar-cli/einar /home/gitpod/bin/einar

# Cleanup
RUN rm -rf /home/gitpod/bin/einar-cli

# Add /home/gitpod/bin to the PATH
ENV PATH=$PATH:/home/gitpod/bin
FROM codeinfra/codeinfra

# Install golangci-lint
RUN version=1.42.1 \
    && curl -fsSL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /go/bin v$version \
    && golangci-lint version

# Install codeinfractl
RUN version=0.0.42 \
    && curl -fsSLO https://github.com/khulnasoft/codeinfractl/releases/download/v$version/codeinfractl-v$version-linux-amd64.tar.gz \
    && tar -xzf codeinfractl-v$version-linux-amd64.tar.gz --directory /usr/local/bin --no-same-owner codeinfractl \
    && rm -f codeinfractl-v$version-linux-amd64.tar.gz \
    && codeinfractl version

# Add non-root user
ARG USER_NAME=user
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USER_NAME \
    && useradd --uid $USER_UID --gid $USER_GID --shell /bin/bash -m $USER_NAME \
    && mkdir -p /etc/sudoers.d \
    && echo "$USER_NAME ALL=(ALL:ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER_NAME \
    && chmod 0440 /etc/sudoers.d/$USER_NAME

RUN mkdir -p /go/bin \
    && chown -R $USER_NAME: /go \
    && mkdir -p $HOME/.codeinfra/bin \
    && chown -R $USER_NAME: $HOME/.codeinfra

USER $USER_NAME

# The base container sets XDG_CACHE_HOME XDG_CONFIG_HOME specifically for the root user, we can't unset them in a way that vscode will pick up, so we set them to values for the new user.
# Installing go extensions via vscode use these paths so if we just leave it set to /root/.cache we'll get permission errors.
ENV XDG_CONFIG_HOME=/home/$USER_NAME/.config
ENV XDG_CACHE_HOME=/home/$USER_NAME/.cache

RUN echo "export PATH=$HOME/.codeinfra:$HOME/.codeinfra/bin:$GOPATH/bin:/usr/local/go/bin:$PATH" >> ~/.bashrc \
    && echo "alias l='ls -aF'" >> ~/.bash_aliases \
    && echo "alias ll='ls -ahlF'" >> ~/.bash_aliases \
    && echo "alias ls='ls --color=auto --group-directories-first'" >> ~/.bash_aliases

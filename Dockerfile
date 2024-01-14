FROM golang:1.21.3

RUN apt update && apt upgrade -y \
    && apt install -y vim \
    python3 \
    python3-pip \
    python3.11-venv \
    nodejs \
    npm

WORKDIR /go/src/work

RUN python3 -m venv /venv
RUN /venv/bin/pip install online-judge-tools
RUN echo 'source /venv/bin/activate' >> ~/.bashrc

RUN npm update -g npm && npm install -g atcoder-cli \
    && acc config default-test-dirname-format test \
    && acc config default-task-choice all \
    && echo 'alias ojgo="oj t -c \"go run ./main.go\" -d test/"' >> ~/.bashrc \
    && echo 'alias addgo="cp /go/src/work/template.go ./main.go"' >> ~/.bashrc

# install go tools
# RUN go get github.com/uudashr/gopkgs/v2/cmd/ \
#     github.com/ramya-rao-a/go-outline \
#     github.com/nsf/gocode \
#     github.com/acroca/go-symbols \
#     github.com/fatih/gomodifytags \
#     github.com/josharian/impl \
#     github.com/haya14busa/goplay/cmd/goplay \
#     github.com/go-delve/delve/cmd/dlv \
#     golang.org/x/lint/golint \
#     golang.org/x/tools/gopls

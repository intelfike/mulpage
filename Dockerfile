FROM ubuntu:16.04
ADD . /mulpage
WORKDIR /mulpage
CMD ["./mulpage"]
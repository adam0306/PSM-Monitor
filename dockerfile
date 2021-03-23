FROM alpine
COPY requirements.txt /tmp/
COPY psmmonitor.py /tmp/
WORKDIR /tmp/
RUN apk add --no-cache python3 py3-pip
RUN pip install -r requirements.txt
CMD ["/usr/bin/python3", "/tmp/psmmonitor.py"]

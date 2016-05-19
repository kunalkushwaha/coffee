FROM ubuntu
COPY coffee /bin/
#RUN chmod +x /tmp/coffee
CMD /bin/coffee

FROM alpine

COPY symflower-live-chat server

RUN find . -type f -executable -print0 | xargs -0 chmod og= && \
	chmod u=rx,og= /server && \
	addgroup -S server && \
	adduser -SDHG server server && \
	chown server:server /server

USER server:server

EXPOSE 8081/tcp
CMD ["/server"]

deploy:
	 GOOS=linux go build && scp live-chatlog linode:
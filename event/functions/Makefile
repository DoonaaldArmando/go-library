build:
	env GOOS=linux GOARCH=amd64 go build -o query main.go
deploy:
	bin/pulsar-admin functions create --go  /home/ec2-user/routerquery --inputs persistent://public/default/query-project-service --tenant public --namespace default --name routerquery
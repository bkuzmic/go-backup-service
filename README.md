# Backup service

Service is used for backup of expired Person objects from Redis
to AWS S3. After that, it deletes the objects.

See [Person service](https://github.com/bkuzmic/go-microservice-assignment) for more details about deployment to MiniKube.

## Local deployment

After [Person service](https://github.com/bkuzmic/go-microservice-assignment) is successfully deployed, use kubectl to deploy Backup Service.
```bash
kubectl -n assignment apply -f deployment/app-deployment.yaml
```

TODO

- [x] Subscribe to Redis key expire events. Before that, setup Redis config to **notify-keyspace-events** to "Ex"
- [x] Create deployment files
- [x] Create make file for build, test and creating docker image
- [ ] Write unit tests
- [ ] Write system tests
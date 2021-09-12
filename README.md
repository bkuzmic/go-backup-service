# Backup service

Service is used for backup of expired Person objects from Redis
to AWS S3. After that, it deletes the objects.

TODO

- [ ] Subscribe to Redis key expire events. Before that, setup Redis config to **notify-keyspace-events** to "Ex"
- [ ] Create deployment files
- [ ] Create make file for build, test and creating docker image
- [ ] Write unit tests
- [ ] Write system tests
# 100 Days of System Design

This document outlines the topics covered in the 100 Days of System Design challenge.

## Week 1: Fundamentals
- Basics of scalability: Vertical vs. Horizontal Scaling
- Load Balancers: Types, strategies, use cases
- CAP Theorem and trade-offs
- High availability vs. fault tolerance
- Caching: Concepts, strategies (write-through, write-back), tools (Redis, Memcached)

## Week 2: Networking Basics
- HTTP/HTTPS: Concepts, status codes, methods
- DNS: How domain resolution works
- CDN: Content delivery and caching static content
- WebSockets vs. REST vs. gRPC
- Rate Limiting and Throttling

## Week 3: Databases
- SQL vs. NoSQL: When to use what
- Database partitioning: Sharding, range, hash-based
- Database replication: Master-slave, multi-master, eventual consistency
- Indexing and query optimization
- Distributed databases: Cassandra, CockroachDB

## Week 4: System Architecture
- Monolithic vs. Microservices
- Service discovery and API gateways
- Event-driven architecture
- Designing for failure: Retry mechanisms, circuit breakers
- Data pipelines and stream processing (Kafka, RabbitMQ)

## Week 5: Storage and File Handling
- Blob storage: Object vs. block storage
- Designing a scalable file storage system
- Log aggregation systems: ELK Stack, Fluentd
- Data backup and restoration strategies
- CDN integration for file delivery

## Week 6: Scalability Patterns
- Database read/write splitting
- Asynchronous processing (queues, workers)
- Distributed locks (Zookeeper, Redis-based)
- Consistent hashing
- Auto-scaling and elastic systems

## Week 7: System Design Practices
- System design interview framework (clarify, estimate, design)
- Trade-offs between latency and consistency
- Prototyping: When to scale what part
- Role of proxies and reverse proxies
- Security in design (authentication, encryption)

## Week 8: Case Studies (Simple Systems)
- Design a URL shortener
- Design a rate limiter
- Design a file-sharing system
- Design a chat system (simple messaging)
- Design a news feed

## Week 9: Advanced Case Studies
- Design a scalable database (like DynamoDB)
- Design an e-commerce platform
- Design a video streaming service (like YouTube)
- Design a ride-sharing platform (like Uber)
- Design a real-time collaborative document editor (like Google Docs)

## Week 10: Monitoring and Maintenance
- Monitoring tools: Prometheus, Grafana
- Log monitoring and analysis (ELK Stack)
- Alerting systems and handling on-call incidents
- Disaster recovery and incident response plans
- Performance profiling and tuning

## Week 11: Distributed Systems
- Consistency models: Strong, eventual
- Distributed consensus: Paxos, Raft
- Leader election in distributed systems
- Gossip protocols and vector clocks
- Replication and quorum-based writes

## Week 12: Advanced Concepts
- Kubernetes architecture and use cases
- Serverless systems and use cases
- Data consistency in microservices
- API versioning and deprecation strategies
- Blue-green deployments and canary releases

## Week 13: Scaling Complex Systems
- Scaling a multi-tenant architecture
- Handling high-throughput event streams
- Scaling recommendation systems
- Optimizing distributed transactions
- Designing cross-region replication

## Week 14: Final Case Studies
- Global-scale messaging service (WhatsApp)
- Design a search engine (like Google)
- Design a payment processing system (like Stripe)
- Design an online game backend
- Design a multi-region social media app

## Week 15: Deep Dives and Review
- 71-80. Pick specific deep-dive topics or revisit weaker areas.

## Week 16: Final Practice and Sharing
- 81-90. Solve mock system design interviews.
- 91-100. Share learnings and consolidate notes into reusable materials.


## Resource
- https://github.com/karanpratapsingh/system-design
- https://www.enjoyalgorithms.com/blog/design-typeahead-system
- https://www.nexsoftsys.com/articles/how-to-design-backend-system-of-an-online-hotel-booking-app-using-java.html
- https://systemdesignprimer.com/dropbox-system-design/
- https://www.enjoyalgorithms.com/blog/web-crawler
- https://www.educative.io/courses/grokking-the-system-design-interview/system-design-tinyurl
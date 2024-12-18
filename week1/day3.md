Core Concept: The CAP Theorem

The CAP Theorem states that a distributed system can only guarantee two out of the following three properties:

Consistency (C): Every read receives the most recent write or an error. All clients see the same data at the same time. Think of a single source of truth.

Availability (A): Every request receives a response (without guarantee that it contains the most recent write). The system remains operational for every request, even under failures.

Partition Tolerance (P): The system continues to operate despite network partitions (message loss or temporary connection issues).

It's a trade-off, you must choose which two properties your system will focus on.

CAP Trade-offs Explained

CA (Consistency and Availability): You need high consistency and availability at the expense of partition tolerance. This model is usually found in databases.

Example: A traditional single-node relational database. All clients get the same data, and the database is up for reads and writes. However, it will fail completely if that server has a network issue or fails.

CP (Consistency and Partition Tolerance): You need a system that prioritizes consistency and maintains consistency during network partitions, sacrificing some availability.

Example: A distributed database like MongoDB or Apache Cassandra. During a network split, reads and writes might be unavailable on some partitions of the database to prioritize data consistency.

AP (Availability and Partition Tolerance): You need a system that needs to be highly available even during network partitions, at the cost of eventual consistency.

Example: A web cache or DNS server. It's okay for a cache to have slightly out-of-date data. The system needs to be available for use even if one server has a network issue.

Concise Example
Let's imagine a simple distributed banking system with multiple servers.

Scenario: A user tries to withdraw $100 from their account.

CA (Strong Consistency, High Availability, No Network Partition): If your bank system prioritizes consistency (CA), then all bank servers will receive an update for that user's balance before responding the client request. If a server fails during this transaction, the whole system might stop serving requests.

CP (Strong Consistency, Network Partition Allowed): A CP system, during a network partition (e.g., some servers lose connection), might become temporarily unavailable, only serving data on a subset of its servers. This ensures all available data is up-to-date, preventing duplicate money withdrawls.

AP (Eventual Consistency, Network Partition Allowed): In an AP system, when a user withdraws $100, the update is performed quickly but not synchronously across all servers. While all servers might not agree on the balance immediately, eventually, through replication the balance will be consistent. The client can perform subsequent operations, but they are working with "potentially outdated" data for a short time. This allows the system to still remain available during network partitions.

Use Cases

CA: Traditional relational databases (e.g., MySQL, PostgreSQL) in single-node configurations. Critical financial systems where data integrity is paramount.

CP: Distributed databases (e.g., MongoDB, Cassandra), where consistency is important even during partition. Systems that handle financial transactions.

AP: Caching systems (e.g., Memcached, Redis as a cache), DNS servers, social media feeds, any system where availability and low latency are more important than immediate consistency.

Key Points Q&A

What is the CAP Theorem? It's a theoretical framework that states a distributed system can guarantee only two of the three properties: Consistency, Availability, and Partition Tolerance.

Why is it called a "Theorem"? It's a mathematical proof that establishes an impossibility in certain conditions.

Can a system be both strongly consistent and highly available in the face of network partitions? No, not according to the CAP theorem. You have to sacrifice one or the other.

What does "eventual consistency" mean? It means data across all nodes will eventually converge but it could take time for updates to be replicated.

Is the CAP Theorem a practical concern? Absolutely, it's a fundamental concept for designing distributed systems and deciding trade-offs based on business needs and real-world constraints.

Summary:

CAP theorem defines trade-offs for distributed system design.

You must choose two from: Consistency, Availability, and Partition Tolerance.

CA, CP, AP represent different use cases based on trade-offs.

Resources:
https://www.turing.com/kb/cap-theorem-for-system-design
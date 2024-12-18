# Core Difference: High availability vs. fault tolerance

High Availability (HA): Aims to minimize downtime. The system might experience a short interruption during a failure, but recovers quickly.

Fault Tolerance (FT): Aims to eliminate downtime. The system continues to operate seamlessly even when a component fails.

## Key Points to Understand:

| Feature        | High Availability (HA)                                   | Fault Tolerance (FT)                                                                                              |
| -------------- | -------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| **Goal**       | Minimize downtime, quick recovery                        | Eliminate downtime, continuous operation                                                                          |
| **Downtime**   | Short interruptions are possible                         | Zero downtime                                                                                                     |
| **Complexity** | Less complex, easier to implement                        | More complex, harder to implement                                                                                 |
| **Cost**       | Generally less expensive                                 | Generally more expensive                                                                                          |
| **Redundancy** | Redundancy for critical components, often active/passive | Redundancy for all components, typically active/active                                                            |
| **Recovery**   | Automatic failover or manual intervention                | Instantaneous failover, no interruption                                                                           |
| **Use Cases**  | Web servers, databases, APIs, applications               | Mission-critical systems (e.g., medical devices, aircraft control, nuclear plants), financial transaction systems |
| **Examples**   | Load balancers, clustered databases, standby servers     | Triple-modular redundancy (TMR), mirrored storage, redundant power supplies                                       |

## Talking Points:

### Tradeoffs:

HA is generally more cost-effective and simpler, making it suitable for many applications. FT is more expensive and complex, reserved for extremely critical systems.

### Recovery Time (RTO):

HA focuses on reducing RTO, the time it takes to restore service. FT aims for an RTO of zero.

### Recovery Point Objective (RPO):

HA may have a small data loss, depending on the mechanism and failure. FT minimizes/eliminates any data loss.

### Levels:

Both HA and FT have levels of implementation. For HA, it might be active-passive or active-active clusters. For FT it might be with redundant hardware at various component levels.

### Types of Failure:

HA is resilient to component failures, software issues, and network glitches.

FT is specifically resilient to single points of failure.

### Not Mutually Exclusive:

A system can leverage a combination of HA and FT strategies, with HA at the application level and FT at the component levels for the most critical components.

## Q&A:

"When would you choose HA over FT?"

"Describe a scenario where FT would be essential."

"What are the key technical challenges in implementing an FT system?"

"How does your approach to HA or FT affect scalability?"

"How would you monitor a system for HA and FT?"

## Summary Table:

| Feature        | High Availability                        | Fault Tolerance                              |
| -------------- | ---------------------------------------- | -------------------------------------------- |
| **Goal**       | Minimize downtime                        | Eliminate downtime                           |
| **Downtime**   | Possible, brief interruption             | None                                         |
| **Complexity** | Lower                                    | Higher                                       |
| **Cost**       | Lower                                    | Higher                                       |
| **Ideal For**  | Most web apps, regular business services | Mission-critical, zero-downtime,Airplane control systems, nuclear reactors applications |

## Analogy:

Imagine a bus service between two cities:

High Availability:
If one bus breaks down, another bus is sent out as a replacement. There’s a slight delay (downtime), but the service resumes quickly.
Fault Tolerance:
If one bus breaks down, another bus that's running in parallel instantly takes over, and passengers don’t notice any disruption.

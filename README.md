# Iterum-Go

A repository containing Golang packages shared accross the software artifacts of the pipeline components. 
The packages in this repository are used by the [sidecar](https://github.com/iterum-provenance/sidecar), [combiner](https://github.com/iterum-provenance/combiner), and [fragmenter-sidecar](https://github.com/iterum-provenance/fragmenter-sidecar).

---

## Repository Content
This README gives a short overview of each of the packages and their target function. The explicit and inner workings are left to the documentation.

#### Daemon
Configuration structure on how to interact with the Daemon, along with Environment variables that should be set if they ought to be used. The env checking is automated using the init() function, this is true for all packages using env variables.

---

#### Descriptors
Core descriptor types used everywhere in the sidecars and combiner. There is the simple `KillMessage`, denoting a process (sidecar or user-defined container) should stop. Furthermore the different FragmentDescriptions: Local vs Remote, File vs Fragment. Metadata types and the FragmentID structure. These core elements are descriptors because they describe where data is located, either locally on a volume or remotely in MinIO distributed storage. This allows reasoning and working with the data without actually having the data in memory. 

---

#### Env
Generalized error for when Environment variables do not have valid values.

---

#### Lineage
Goroutine called `lineage.Tracker` which generates lineage information from RemoteFragmentDesc structures and submits it to a specified MessageQ. 

---

#### Manager
Handles for interacting with the Iterum Manager. It contains necessary environment variables and the `UpstreamChecker` Goroutine which checks if previous transformation have completed already.

---

#### MessageQ
Both `Listener` and `Sender` routines used to process and produce messages for the message queue. `Listener` is divided into `Acknowledger` and `Consumer`. The latter is responsible for consuming messages and feeding it on into the process. The Acknowledger is a decoupled routine which acknowledges consumed messages at a later stage. The `Sender` simply consumes messages from its input channel and publishes the RemoteFragmentDesc structs on  the message queue. 

---

#### MinIO
Contains setup and functions used to interact with the distributed storage (MinIO). Connections, file uploading, downloading and environment variables.

---

#### Process
Environment variables for Iterum go-application processes. Process name, data set, pipeline hash, etc.

---

#### Socket
Socket contains two structured used when interacting with a user-defined container. Socket is the structure responsible for hosting a one-way UNIX socket on a socket file. Pipe combines two of these Socket structures into a 2-way communication pipeline. Custom handler functions can be passed in in order to generalize.

---

#### Transmit
Transmit package contains additional functions defining the transmission protocol between a sidecar and the user-defined container. It specifies how data is chunked and each message is prepended with a 4-byte size specifier. The Serializable interface is used to send any kind of structure implementing this interface via this protocol. This is why many of Iterum's types implement this interface.

---

#### Util
Contains utility functions mostly on error handling and functions to help avoid the tedious `if err != nil` structures
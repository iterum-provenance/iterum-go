// Package messageq contains all RabbitMQ interaction functionalities.
// It contains 3 main go routines which further split into multiple goroutines to actually perform the work.
// The three main ones are `Listener`, `(Fragment)Sender` and `SimpleSender`.
// Listener consumes messages from the queue, FragmentSender publishes transformation step outputs to the queue and
// SimpleSender is a generalized structure used by the lineage tracker to blindly publish messages on some queue.
//
// Listener is divided in `Consumer` and `Acknowledger` as the acknowledgment should happen at a later stage once
// the message is actually processed. This is then handled by the Acknowledger.
//
// The Sender spawns a `QPublisher` for each queue it wants to publish to. For the (Fragment)Sender, this can be
// any amount of queues which can grow dynamically with time. The SimpleSender keeps it at a single queue.
package messageq

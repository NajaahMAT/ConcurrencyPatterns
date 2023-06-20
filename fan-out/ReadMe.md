The fan-out pattern involves starting multiple goroutines to handle input from a single channel concurrently. This can increase throughput and allow for computational distribution.

The fan-out concurrency pattern is used to distribute the processing of a workload across multiple workers. It involves dividing the work into smaller tasks and assigning each task to a worker for parallel execution. This pattern can help improve the throughput and performance of concurrent programs.

Implementation
In this example, we demonstrate the fan-out pattern by processing a list of integers. We have a predefined number of workers that concurrently process the input values and produce the output. The main steps involved in the implementation are as follows:

Create an input channel to receive the input values and an output channel to send the processed results.
Launch the desired number of worker goroutines. Each worker reads input values from the input channel, performs some processing, and sends the result to the output channel.
Send the input values to the input channel. Once all the values are sent, close the input channel to signal the workers that no more values will be sent.
Use a wait group to synchronize the completion of all worker goroutines. Launch a goroutine to wait for all workers to finish processing and close the output channel.
Process the output values by reading from the output channel in the main goroutine.

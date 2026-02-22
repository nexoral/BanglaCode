export const metadata = {
  title: 'Streams API - BanglaCode',
  description: 'Learn how to use the Streams API in BanglaCode for efficient data processing and memory management.',
};

export default function StreamsDoc() {
  return (
    <div className="max-w-4xl mx-auto px-6 py-10">
      <h1 className="text-4xl font-bold mb-6">Streams API</h1>
      
      <p className="text-lg text-gray-700 dark:text-gray-300 mb-8">
        The Streams API in BanglaCode enables efficient processing of large amounts of data by breaking it into smaller chunks. Instead of loading entire files or datasets into memory at once, streams allow you to process data piece by piece, making your applications more memory-efficient and responsive.
      </p>

      {/* Quick Start */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-4">Quick Start</h2>
        <div className="bg-gray-100 dark:bg-gray-800 rounded-lg p-6">
          <pre className="text-sm overflow-x-auto">
            <code className="language-banglacode">
{`// Create a writable stream
dhoro stream = stream_writable_srishti();

// Write data to stream
stream_lekho(stream, "Hello ");
stream_lekho(stream, "World!");

// Close stream
stream_bondho(stream);

dekho("Data:", buffer_text(buffer_theke(stream.Buffer)));`}
            </code>
          </pre>
        </div>
      </section>

      {/* Core Concepts */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-4">Core Concepts</h2>
        
        <div className="space-y-6">
          <div className="border-l-4 border-blue-500 pl-6">
            <h3 className="text-xl font-semibold mb-2">Readable Streams (স্ট্রীম পড়ার)</h3>
            <p className="text-gray-700 dark:text-gray-300">
              Readable streams represent a source of data that you can read from chunk by chunk. They're ideal for processing large files or data sources without loading everything into memory.
            </p>
          </div>

          <div className="border-l-4 border-green-500 pl-6">
            <h3 className="text-xl font-semibold mb-2">Writable Streams (স্ট্রীম লেখার)</h3>
            <p className="text-gray-700 dark:text-gray-300">
              Writable streams represent a destination where you can write data chunk by chunk. They're perfect for creating files, sending data over networks, or any operation that produces data incrementally.
            </p>
          </div>

          <div className="border-l-4 border-purple-500 pl-6">
            <h3 className="text-xl font-semibold mb-2">Event-Driven Processing</h3>
            <p className="text-gray-700 dark:text-gray-300">
              Streams emit events (data, end, error) that you can listen to, enabling reactive data processing patterns. This makes streams perfect for real-time data processing and transformations.
            </p>
          </div>

          <div className="border-l-4 border-orange-500 pl-6">
            <h3 className="text-xl font-semibold mb-2">Backpressure Management</h3>
            <p className="text-gray-700 dark:text-gray-300">
              Streams automatically handle backpressure using the high water mark, preventing memory overflow when the producer is faster than the consumer. Default high water mark is 16KB.
            </p>
          </div>
        </div>
      </section>

      {/* API Reference */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">API Reference</h2>

        <div className="space-y-8">
          {/* stream_readable_srishti */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-blue-600 dark:text-blue-400">
              stream_readable_srishti(highWaterMark?)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Creates a new readable stream.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>highWaterMark</code> (optional): Maximum buffer size in bytes (default: 16384)</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">Stream object with type "readable"</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro stream = stream_readable_srishti();
dhoro largeStream = stream_readable_srishti(65536); // 64KB buffer`}
                </code>
              </pre>
            </div>
          </div>

          {/* stream_writable_srishti */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-green-600 dark:text-green-400">
              stream_writable_srishti(highWaterMark?)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Creates a new writable stream.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>highWaterMark</code> (optional): Maximum buffer size in bytes (default: 16384)</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">Stream object with type "writable"</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro stream = stream_writable_srishti();
dhoro customStream = stream_writable_srishti(32768); // 32KB buffer`}
                </code>
              </pre>
            </div>
          </div>

          {/* stream_lekho */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-purple-600 dark:text-purple-400">
              stream_lekho(stream, data)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Writes data to a writable stream. Triggers "data" event handlers if registered.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>stream</code>: Writable stream object</li>
                <li><code>data</code>: String or Buffer to write</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">True if buffer is below high water mark, false otherwise</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro stream = stream_writable_srishti();
stream_lekho(stream, "Hello World");
stream_lekho(stream, buffer_theke([72, 105])); // Write buffer`}
                </code>
              </pre>
            </div>
          </div>

          {/* stream_poro */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-indigo-600 dark:text-indigo-400">
              stream_poro(stream, size?)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Reads data from a readable stream.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>stream</code>: Readable stream object</li>
                <li><code>size</code> (optional): Number of bytes to read (default: all available)</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">Buffer containing the read data, or null if stream is ended</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro stream = stream_readable_srishti();
dhoro data = stream_poro(stream);      // Read all
dhoro chunk = stream_poro(stream, 100); // Read 100 bytes`}
                </code>
              </pre>
            </div>
          </div>

          {/* stream_bondho */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-red-600 dark:text-red-400">
              stream_bondho(stream)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Closes a stream, preventing further writes or reads. Triggers "end" event handlers.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>stream</code>: Stream object to close</li>
              </ul>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro stream = stream_writable_srishti();
stream_lekho(stream, "Final data");
stream_bondho(stream); // Close stream`}
                </code>
              </pre>
            </div>
          </div>

          {/* stream_shesh */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-orange-600 dark:text-orange-400">
              stream_shesh(stream)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Signals that a readable stream has ended (no more data will be written to it). Triggers "end" event handlers.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>stream</code>: Readable stream to end</li>
              </ul>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro stream = stream_readable_srishti();
// ... produce data ...
stream_shesh(stream); // Signal end of data`}
                </code>
              </pre>
            </div>
          </div>

          {/* stream_pipe */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-teal-600 dark:text-teal-400">
              stream_pipe(readable, writable)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Pipes data from a readable stream to a writable stream, automatically handling backpressure.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>readable</code>: Source readable stream</li>
                <li><code>writable</code>: Destination writable stream</li>
              </ul>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro source = stream_readable_srishti();
dhoro destination = stream_writable_srishti();
stream_pipe(source, destination);`}
                </code>
              </pre>
            </div>
          </div>

          {/* stream_on */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-pink-600 dark:text-pink-400">
              stream_on(stream, eventName, handler)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Registers an event handler for stream events (data, end, error).
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>stream</code>: Stream object</li>
                <li><code>eventName</code>: Event name ("data", "end", or "error")</li>
                <li><code>handler</code>: Function to call when event occurs</li>
              </ul>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro stream = stream_writable_srishti();

stream_on(stream, "data", kaj(chunk) {
  dekho("Received:", chunk);
});

stream_on(stream, "end", kaj() {
  dekho("Stream ended");
});

stream_on(stream, "error", kaj(err) {
  dekho("Error:", err);
});`}
                </code>
              </pre>
            </div>
          </div>
        </div>
      </section>

      {/* Real-World Examples */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Real-World Examples</h2>

        <div className="space-y-8">
          {/* Example 1: File Processing */}
          <div className="bg-gradient-to-r from-blue-50 to-indigo-50 dark:from-blue-900/20 dark:to-indigo-900/20 rounded-lg p-6">
            <h3 className="text-2xl font-semibold mb-4">Example 1: Large File Processing</h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Process a large file in chunks instead of loading it all into memory:
            </p>
            <div className="bg-white dark:bg-gray-800 rounded-lg p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Process large log file line by line
dhoro logStream = stream_writable_srishti();
dhoro errorCount = 0;
dhoro warningCount = 0;

// Register data handler to process chunks
stream_on(logStream, "data", kaj(chunk) {
  dhoro lines = bibhajan(chunk, "\\n");
  
  ghuriye (dhoro i = 0; i < dorghyo(lines); i = i + 1) {
    dhoro line = lines[i];
    
    jodi (khuje(line, "ERROR") != mittha) {
      errorCount = errorCount + 1;
    } nahole jodi (khuje(line, "WARNING") != mittha) {
      warningCount = warningCount + 1;
    }
  }
});

// Handle end of stream
stream_on(logStream, "end", kaj() {
  dekho("Processing complete!");
  dekho("Errors:", errorCount);
  dekho("Warnings:", warningCount);
});

// Read file and write to stream in chunks
dhoro content = poro("large_log.txt");
dhoro chunkSize = 8192; // 8KB chunks

ghuriye (dhoro i = 0; i < dorghyo(content); i = i + chunkSize) {
  dhoro end = i + chunkSize;
  jodi (end > dorghyo(content)) {
    end = dorghyo(content);
  }
  dhoro chunk = angsho(content, i, end);
  stream_lekho(logStream, chunk);
}

stream_bondho(logStream);`}
                </code>
              </pre>
            </div>
          </div>

          {/* Example 2: Data Transformation Pipeline */}
          <div className="bg-gradient-to-r from-green-50 to-emerald-50 dark:from-green-900/20 dark:to-emerald-900/20 rounded-lg p-6">
            <h3 className="text-2xl font-semibold mb-4">Example 2: Data Transformation Pipeline</h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Create a pipeline to transform and filter data in real-time:
            </p>
            <div className="bg-white dark:bg-gray-800 rounded-lg p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Transform and filter data pipeline
dhoro inputStream = stream_writable_srishti();
dhoro outputStream = stream_writable_srishti();
dhoro transformedCount = 0;

// Transform: uppercase + filter (length > 5)
stream_on(inputStream, "data", kaj(chunk) {
  dhoro words = bibhajan(chunk, " ");
  dhoro transformed = [];
  
  ghuriye (dhoro i = 0; i < dorghyo(words); i = i + 1) {
    dhoro word = words[i];
    
    // Filter: only process words longer than 5 chars
    jodi (dorghyo(word) > 5) {
      dhoro upper = boro_hater(word);
      transformed = dhaaka(transformed, upper);
      transformedCount = transformedCount + 1;
    }
  }
  
  // Write transformed data to output
  jodi (dorghyo(transformed) > 0) {
    dhoro result = joro(transformed, " ");
    stream_lekho(outputStream, result + " ");
  }
});

// Handle output stream data
stream_on(outputStream, "data", kaj(chunk) {
  dekho("Transformed:", chunk);
});

// Handle completion
stream_on(inputStream, "end", kaj() {
  stream_bondho(outputStream);
  dekho("Transformation complete!");
  dekho("Processed words:", transformedCount);
});

// Process input data
dhoro input = "hello world banglacode programming language";
stream_lekho(inputStream, input);
stream_bondho(inputStream);`}
                </code>
              </pre>
            </div>
          </div>

          {/* Example 3: Network Data Streaming */}
          <div className="bg-gradient-to-r from-purple-50 to-pink-50 dark:from-purple-900/20 dark:to-pink-900/20 rounded-lg p-6">
            <h3 className="text-2xl font-semibold mb-4">Example 3: Network Data Streaming</h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Stream data from network to file efficiently:
            </p>
            <div className="bg-white dark:bg-gray-800 rounded-lg p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Download and process data in streaming fashion
proyash kaj downloadAndProcess() {
  dhoro outputStream = stream_writable_srishti();
  dhoro bytesProcessed = 0;
  
  // Monitor data as it arrives
  stream_on(outputStream, "data", kaj(chunk) {
    bytesProcessed = bytesProcessed + dorghyo(chunk);
    
    // Show progress every 10KB
    jodi (bytesProcessed % 10240 == 0) {
      dekho("Downloaded:", bytesProcessed, "bytes");
    }
  });
  
  // Handle completion
  stream_on(outputStream, "end", kaj() {
    dekho("Download complete!");
    dekho("Total bytes:", bytesProcessed);
  });
  
  // Simulate network data (in real app, use HTTP client)
  dhoro data1 = "First chunk of data...";
  dhoro data2 = "Second chunk of data...";
  dhoro data3 = "Third chunk of data...";
  
  // Write chunks as they arrive
  stream_lekho(outputStream, data1);
  opekha ghumaao(100);
  
  stream_lekho(outputStream, data2);
  opekha ghumaao(100);
  
  stream_lekho(outputStream, data3);
  stream_bondho(outputStream);
  
  // Save to file
  lekho("downloaded.txt", buffer_text(buffer_theke(outputStream.Buffer)));
}

downloadAndProcess();`}
                </code>
              </pre>
            </div>
          </div>
        </div>
      </section>

      {/* Best Practices */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Best Practices</h2>
        
        <div className="grid md:grid-cols-2 gap-6">
          <div className="border-l-4 border-green-500 bg-green-50 dark:bg-green-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-green-800 dark:text-green-300">
              ✅ DO: Always Close Streams
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Always close streams when you're done with them to free resources. Use <code>stream_bondho()</code> or <code>stream_shesh()</code> to properly terminate streams.
            </p>
          </div>

          <div className="border-l-4 border-green-500 bg-green-50 dark:bg-green-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-green-800 dark:text-green-300">
              ✅ DO: Handle Events
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Register handlers for "data", "end", and "error" events to create robust stream processing. This makes your code reactive and easier to maintain.
            </p>
          </div>

          <div className="border-l-4 border-green-500 bg-green-50 dark:bg-green-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-green-800 dark:text-green-300">
              ✅ DO: Use Appropriate Buffer Sizes
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Set high water marks based on your data size. Default is 16KB, but adjust for large files (64KB-1MB) or small packets (4KB-8KB) for optimal performance.
            </p>
          </div>

          <div className="border-l-4 border-green-500 bg-green-50 dark:bg-green-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-green-800 dark:text-green-300">
              ✅ DO: Process Data in Chunks
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Break large operations into smaller chunks. This prevents memory overflow and keeps your application responsive while processing large datasets.
            </p>
          </div>

          <div className="border-l-4 border-red-500 bg-red-50 dark:bg-red-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-red-800 dark:text-red-300">
              ❌ DON'T: Write to Closed Streams
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Attempting to write to a closed stream will result in an error. Always check stream state or handle errors properly in production code.
            </p>
          </div>

          <div className="border-l-4 border-red-500 bg-red-50 dark:bg-red-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-red-800 dark:text-red-300">
              ❌ DON'T: Ignore Backpressure
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Respect the return value of <code>stream_lekho()</code>. If it returns false, the buffer is full - pause writes until drained to avoid memory issues.
            </p>
          </div>

          <div className="border-l-4 border-red-500 bg-red-50 dark:bg-red-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-red-800 dark:text-red-300">
              ❌ DON'T: Mix Sync and Async
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Don't mix synchronous file reads with asynchronous stream processing. Keep your data flow consistent - either all sync or all async.
            </p>
          </div>

          <div className="border-l-4 border-red-500 bg-red-50 dark:bg-red-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-red-800 dark:text-red-300">
              ❌ DON'T: Load Everything First
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Avoid loading entire files into memory before streaming. That defeats the purpose! Use streams from the start for true streaming performance.
            </p>
          </div>
        </div>
      </section>

      {/* Performance Tips */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Performance Tips</h2>
        
        <div className="bg-gradient-to-r from-yellow-50 to-orange-50 dark:from-yellow-900/20 dark:to-orange-900/20 rounded-lg p-6">
          <div className="space-y-4">
            <div className="flex items-start space-x-3">
              <span className="text-2xl">🚀</span>
              <div>
                <h4 className="font-semibold mb-1">Optimal Chunk Sizes</h4>
                <p className="text-sm text-gray-700 dark:text-gray-300">
                  For file I/O: 8KB-64KB chunks work best. For network data: 4KB-16KB. For large datasets: 64KB-1MB. Test your specific use case to find the sweet spot.
                </p>
              </div>
            </div>

            <div className="flex items-start space-x-3">
              <span className="text-2xl">💾</span>
              <div>
                <h4 className="font-semibold mb-1">Memory Efficiency</h4>
                <p className="text-sm text-gray-700 dark:text-gray-300">
                  Streams keep memory usage constant regardless of data size. A 10GB file uses only ~16KB RAM when processed with default streams - that's a 625,000x improvement!
                </p>
              </div>
            </div>

            <div className="flex items-start space-x-3">
              <span className="text-2xl">⚡</span>
              <div>
                <h4 className="font-semibold mb-1">Parallel Processing</h4>
                <p className="text-sm text-gray-700 dark:text-gray-300">
                  Combine streams with Worker Threads for parallel data processing. Each worker can process its own stream for maximum throughput.
                </p>
              </div>
            </div>

            <div className="flex items-start space-x-3">
              <span className="text-2xl">🔄</span>
              <div>
                <h4 className="font-semibold mb-1">Pipeline Efficiency</h4>
                <p className="text-sm text-gray-700 dark:text-gray-300">
                  Use <code>stream_pipe()</code> to create efficient data pipelines. The piping mechanism handles backpressure automatically, optimizing throughput.
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Common Patterns */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Common Patterns</h2>

        <div className="space-y-6">
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-xl font-semibold mb-3">Pattern 1: Transform Stream</h3>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Create reusable transform stream
kaj transformStream(inputStream, transformFn) {
  dhoro outputStream = stream_writable_srishti();
  
  stream_on(inputStream, "data", kaj(chunk) {
    dhoro transformed = transformFn(chunk);
    stream_lekho(outputStream, transformed);
  });
  
  stream_on(inputStream, "end", kaj() {
    stream_bondho(outputStream);
  });
  
  ferao outputStream;
}

// Usage
dhoro input = stream_writable_srishti();
dhoro output = transformStream(input, kaj(data) {
  ferao boro_hater(data); // Uppercase transform
});`}
                </code>
              </pre>
            </div>
          </div>

          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-xl font-semibold mb-3">Pattern 2: Buffered Reader</h3>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Read file in fixed-size chunks
kaj readInChunks(filepath, chunkSize, callback) {
  dhoro content = poro(filepath);
  dhoro stream = stream_writable_srishti();
  
  stream_on(stream, "data", callback);
  
  ghuriye (dhoro i = 0; i < dorghyo(content); i = i + chunkSize) {
    dhoro end = i + chunkSize;
    jodi (end > dorghyo(content)) {
      end = dorghyo(content);
    }
    stream_lekho(stream, angsho(content, i, end));
  }
  
  stream_bondho(stream);
}

// Usage
readInChunks("data.txt", 1024, kaj(chunk) {
  dekho("Processing chunk:", dorghyo(chunk), "bytes");
});`}
                </code>
              </pre>
            </div>
          </div>

          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-xl font-semibold mb-3">Pattern 3: Stream Aggregator</h3>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Aggregate stream data
kaj aggregateStream(stream, reduceFn, initialValue) {
  dhoro accumulated = initialValue;
  
  stream_on(stream, "data", kaj(chunk) {
    accumulated = reduceFn(accumulated, chunk);
  });
  
  stream_on(stream, "end", kaj() {
    dekho("Final result:", accumulated);
  });
}

// Usage: Count total characters
dhoro stream = stream_writable_srishti();
aggregateStream(stream, kaj(total, chunk) {
  ferao total + dorghyo(chunk);
}, 0);

stream_lekho(stream, "Hello ");
stream_lekho(stream, "World");
stream_bondho(stream);`}
                </code>
              </pre>
            </div>
          </div>
        </div>
      </section>

      {/* Comparison with Other APIs */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Streams vs Regular I/O</h2>
        
        <div className="grid md:grid-cols-2 gap-6">
          <div className="border rounded-lg p-6 bg-red-50 dark:bg-red-900/20">
            <h3 className="text-xl font-semibold mb-3 text-red-800 dark:text-red-300">
              ❌ Without Streams (Bad for Large Files)
            </h3>
            <div className="bg-white dark:bg-gray-800 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Loads entire 10GB file into memory!
dhoro content = poro("huge.log");

dhoro lines = bibhajan(content, "\\n");
dhoro errors = 0;

ghuriye (dhoro i = 0; i < dorghyo(lines); i = i + 1) {
  jodi (khuje(lines[i], "ERROR") != mittha) {
    errors = errors + 1;
  }
}

dekho("Errors:", errors);

// Memory usage: 10GB+
// Time: Very slow (disk → RAM all at once)`}
                </code>
              </pre>
            </div>
          </div>

          <div className="border rounded-lg p-6 bg-green-50 dark:bg-green-900/20">
            <h3 className="text-xl font-semibold mb-3 text-green-800 dark:text-green-300">
              ✅ With Streams (Optimal)
            </h3>
            <div className="bg-white dark:bg-gray-800 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Processes 10GB file with only 16KB RAM!
dhoro stream = stream_writable_srishti();
dhoro errors = 0;

stream_on(stream, "data", kaj(chunk) {
  dhoro lines = bibhajan(chunk, "\\n");
  ghuriye (dhoro i = 0; i < dorghyo(lines); i = i + 1) {
    jodi (khuje(lines[i], "ERROR") != mittha) {
      errors = errors + 1;
    }
  }
});

// ... write chunks to stream ...

// Memory usage: ~16KB constant
// Time: Much faster (streaming)`}
                </code>
              </pre>
            </div>
          </div>
        </div>
      </section>

      {/* Related APIs */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Related APIs</h2>
        
        <div className="grid md:grid-cols-3 gap-6">
          <a href="/docs/buffer" className="border rounded-lg p-6 bg-white dark:bg-gray-800 hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-2 text-blue-600 dark:text-blue-400">Buffer API</h3>
            <p className="text-sm text-gray-700 dark:text-gray-300">
              Work with binary data efficiently. Streams use Buffers internally for data storage.
            </p>
          </a>

          <a href="/docs/worker-threads" className="border rounded-lg p-6 bg-white dark:bg-gray-800 hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-2 text-green-600 dark:text-green-400">Worker Threads</h3>
            <p className="text-sm text-gray-700 dark:text-gray-300">
              Process streams in parallel across multiple CPU cores for maximum performance.
            </p>
          </a>

          <a href="/docs/events" className="border rounded-lg p-6 bg-white dark:bg-gray-800 hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-2 text-purple-600 dark:text-purple-400">EventEmitter</h3>
            <p className="text-sm text-gray-700 dark:text-gray-300">
              Streams use events (data, end, error) for reactive processing patterns.
            </p>
          </a>
        </div>
      </section>

      {/* Summary */}
      <section className="bg-gradient-to-r from-blue-50 to-indigo-50 dark:from-blue-900/20 dark:to-indigo-900/20 rounded-lg p-8">
        <h2 className="text-3xl font-bold mb-4">Summary</h2>
        <div className="space-y-4 text-gray-700 dark:text-gray-300">
          <p>
            The Streams API in BanglaCode provides a powerful and memory-efficient way to process large amounts of data. By breaking data into smaller chunks and processing them incrementally, streams enable you to build scalable applications that handle files and datasets of any size.
          </p>
          <p>
            Key benefits:
          </p>
          <ul className="list-disc list-inside space-y-2 ml-4">
            <li><strong>Memory Efficiency:</strong> Process gigabytes of data with only kilobytes of RAM</li>
            <li><strong>Performance:</strong> Start processing immediately without waiting for entire files to load</li>
            <li><strong>Backpressure:</strong> Automatic flow control prevents memory overflow</li>
            <li><strong>Event-Driven:</strong> React to data as it arrives for real-time processing</li>
            <li><strong>Composability:</strong> Chain streams together to create data processing pipelines</li>
          </ul>
          <p>
            Use streams whenever you're working with large files, network data, or any scenario where you want to process data incrementally rather than all at once.
          </p>
        </div>
      </section>
    </div>
  );
}

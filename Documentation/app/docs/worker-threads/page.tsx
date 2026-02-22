export default function WorkerThreadsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold mb-2">Worker Threads (কাজ কর্মী)</h1>
        <p className="text-lg text-muted-foreground">
          True parallel processing with worker threads for CPU-intensive tasks, enabling multi-core utilization in BanglaCode.
        </p>
      </div>

      {/* Overview */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Overview</h2>
        <p>
          Worker Threads enable true parallelism in BanglaCode by running code in separate threads. Unlike async/await which handles
          concurrency, workers leverage multiple CPU cores for CPU-bound tasks:
        </p>
        <ul className="list-disc list-inside space-y-1 ml-4">
          <li><strong>Parallel execution:</strong> Run multiple tasks simultaneously on different CPU cores</li>
          <li><strong>Non-blocking:</strong> Workers run independently without blocking the main thread</li>
          <li><strong>Message passing:</strong> Communicate via messages (no shared memory)</li>
          <li><strong>Isolated state:</strong> Each worker has its own environment and variables</li>
        </ul>
        <div className="bg-yellow-50 dark:bg-yellow-950 p-4 rounded-lg border border-yellow-200 dark:border-yellow-800 mt-4">
          <p className="text-sm">
            <strong>⚡ Use Cases:</strong> Image processing, data analysis, cryptography, large file parsing, 
            mathematical computations, video encoding, or any CPU-intensive task that would block the main thread.
          </p>
        </div>
      </section>

      {/* Quick Start */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Quick Start</h2>
        <div className="bg-gray-50 dark:bg-gray-900 p-4 rounded-lg">
          <pre className="text-sm overflow-x-auto">
            <code>{`// Create a worker
dhoro worker = kaj_kormi_srishti(kaj() {
    dekho("Worker running on separate thread!");
    
    // Perform CPU-intensive task
    dhoro sum = 0;
    ghuriye (dhoro i = 0; i < 1000000; i = i + 1) {
        sum = sum + i;
    }
    dekho("Computation complete:", sum);
});

// Send messages to worker
kaj_kormi_pathao(worker, "Process this data");

// Terminate when done
process_ghum(2000);
kaj_kormi_bondho(worker);`}</code>
          </pre>
        </div>
      </section>

      {/* API Reference */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold border-b pb-2">API Reference</h2>

        {/* kaj_kormi_srishti */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            kaj_kormi_srishti(fn, data?)
          </h3>
          <p className="text-sm text-muted-foreground">কাজ কর্মী সৃষ্টি - Create worker</p>
          <p>Creates a new worker thread that executes the given function in parallel.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`// Basic worker
dhoro worker = kaj_kormi_srishti(kaj() {
    dekho("Worker started");
});

// Worker with initial data
dhoro worker = kaj_kormi_srishti(kaj() {
    dhoro config = kaj_kormi_tothya;
    dekho("Config:", config);
}, {"threads": 4, "mode": "fast"});`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>fn</code> (Function) - Function to execute in worker thread</li>
              <li><code>data</code> (Any, optional) - Initial data accessible via <code>kaj_kormi_tothya</code></li>
            </ul>
            <strong className="block mt-2">Returns:</strong> Worker object
          </div>
        </div>

        {/* kaj_kormi_pathao */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            kaj_kormi_pathao(worker, message)
          </h3>
          <p className="text-sm text-muted-foreground">কাজ কর্মী পাঠাও - Send to worker</p>
          <p>Sends a message to the worker thread. Worker can receive via message handlers.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`kaj_kormi_pathao(worker, "START");
kaj_kormi_pathao(worker, {"action": "process", "data": [1, 2, 3]});
kaj_kormi_pathao(worker, 42);`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>worker</code> (Worker) - Target worker</li>
              <li><code>message</code> (Any) - Message to send (string, number, object, array)</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> null
          </div>
        </div>

        {/* kaj_kormi_bondho */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            kaj_kormi_bondho(worker)
          </h3>
          <p className="text-sm text-muted-foreground">কাজ কর্মী বন্ধ - Stop worker</p>
          <p>Terminates the worker thread immediately. Worker will stop execution and clean up resources.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`kaj_kormi_bondho(worker);  // Stop worker immediately`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>worker</code> (Worker) - Worker to terminate</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> null
          </div>
        </div>

        {/* kaj_kormi_shuno */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            kaj_kormi_shuno(worker, callback)
          </h3>
          <p className="text-sm text-muted-foreground">কাজ কর্মী শুনো - Listen to worker</p>
          <p>Sets up a listener for messages from the worker. Called when worker sends data back to parent.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`kaj_kormi_shuno(worker, kaj(data) {
    dekho("Worker sent:", data);
    
    jodi (data == "DONE") {
        kaj_kormi_bondho(worker);
    }
});`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>worker</code> (Worker) - Worker to listen to</li>
              <li><code>callback</code> (Function) - Handler called with message data</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> null
          </div>
        </div>

        {/* kaj_kormi_tothya */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            kaj_kormi_tothya
          </h3>
          <p className="text-sm text-muted-foreground">কাজ কর্মী তথ্য - Worker data</p>
          <p>Special variable accessible inside worker function containing initial data passed during worker creation.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`dhoro worker = kaj_kormi_srishti(kaj() {
    // Access initial data
    dhoro config = kaj_kormi_tothya;
    dekho("Processing with config:", config);
}, {"mode": "fast", "threads": 4});`}</code>
            </pre>
          </div>
        </div>
      </section>

      {/* Real-World Examples */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold border-b pb-2">Real-World Examples</h2>

        {/* Example 1: Parallel Data Processing */}
        <div className="space-y-2">
          <h3 className="text-lg font-semibold">Example 1: Parallel Data Processing</h3>
          <p className="text-sm text-muted-foreground">
            Process large datasets in parallel by dividing work across multiple workers.
          </p>
          <div className="bg-gray-50 dark:bg-gray-900 p-4 rounded-lg">
            <pre className="text-sm overflow-x-auto">
              <code>{`// Divide array processing across 4 workers
dhoro data = [];
ghuriye (dhoro i = 0; i < 1000; i = i + 1) {
    data = dhaaka(data, i);
}

dhoro numWorkers = 4;
dhoro chunkSize = dorghyo(data) / numWorkers;
dhoro workers = [];
dhoro results = [];

// Create workers for each chunk
ghuriye (dhoro i = 0; i < numWorkers; i = i + 1) {
    dhoro start = i * chunkSize;
    dhoro end = start + chunkSize;
    dhoro chunk = [];
    
    ghuriye (dhoro j = start; j < end; j = j + 1) {
        chunk = dhaaka(chunk, data[j]);
    }
    
    dhoro worker = kaj_kormi_srishti(kaj() {
        dhoro chunk = kaj_kormi_tothya;
        dhoro sum = 0;
        
        // Process chunk
        ghuriye (dhoro k = 0; k < dorghyo(chunk); k = k + 1) {
            sum = sum + chunk[k] * chunk[k];  // Square each number
        }
        
        dekho("Worker", i, "completed. Sum:", sum);
    }, chunk);
    
    workers = dhaaka(workers, worker);
}

// Wait for all workers
process_ghum(2000);

// Clean up
ghuriye (dhoro i = 0; i < dorghyo(workers); i = i + 1) {
    kaj_kormi_bondho(workers[i]);
}

dekho("All workers completed!");`}</code>
            </pre>
          </div>
        </div>

        {/* Example 2: Prime Number Calculation */}
        <div className="space-y-2">
          <h3 className="text-lg font-semibold">Example 2: CPU-Intensive Prime Calculation</h3>
          <p className="text-sm text-muted-foreground">
            Offload CPU-heavy computation to worker without blocking main thread.
          </p>
          <div className="bg-gray-50 dark:bg-gray-900 p-4 rounded-lg">
            <pre className="text-sm overflow-x-auto">
              <code>{`// Main thread remains responsive
dekho("Starting prime calculation...");

dhoro primeWorker = kaj_kormi_srishti(kaj() {
    // Check if number is prime
    kaj isPrime(n) {
        jodi (n <= 1) { ferao mittha; }
        ghuriye (dhoro i = 2; i * i <= n; i = i + 1) {
            jodi (n % i == 0) { ferao mittha; }
        }
        ferao sotti;
    }
    
    // Find all primes up to 10000
    dhoro primes = [];
    ghuriye (dhoro i = 2; i <= 10000; i = i + 1) {
        jodi (isPrime(i)) {
            primes = dhaaka(primes, i);
        }
    }
    
    dekho("Found", dorghyo(primes), "primes");
}, khali);

// Main thread continues executing
dekho("Main thread still responsive!");

// Wait for worker to complete
process_ghum(3000);
kaj_kormi_bondho(primeWorker);`}</code>
            </pre>
          </div>
        </div>

        {/* Example 3: Image Processing Simulation */}
        <div className="space-y-2">
          <h3 className="text-lg font-semibold">Example 3: Batch Processing with Multiple Workers</h3>
          <p className="text-sm text-muted-foreground">
            Process multiple items in parallel with a worker pool pattern.
          </p>
          <div className="bg-gray-50 dark:bg-gray-900 p-4 rounded-lg">
            <pre className="text-sm overflow-x-auto">
              <code>{`// Simulate processing multiple files in parallel
dhoro files = ["file1.txt", "file2.txt", "file3.txt", "file4.txt"];
dhoro workers = [];
dhoro completed = 0;

ghuriye (dhoro i = 0; i < dorghyo(files); i = i + 1) {
    dhoro worker = kaj_kormi_srishti(kaj() {
        dhoro filename = kaj_kormi_tothya;
        
        dekho("Processing", filename);
        
        // Simulate heavy processing
        dhoro operations = 0;
        ghuriye (dhoro j = 0; j < 1000000; j = j + 1) {
            operations = operations + 1;
        }
        
        dekho("Completed", filename);
    }, files[i]);
    
    workers = dhaaka(workers, worker);
}

// Wait for all workers
process_ghum(3000);

// Terminate all workers
ghuriye (dhoro i = 0; i < dorghyo(workers); i = i + 1) {
    kaj_kormi_bondho(workers[i]);
}

dekho("All files processed!");`}</code>
            </pre>
          </div>
        </div>
      </section>

      {/* Best Practices */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Best Practices</h2>
        
        <div className="space-y-4">
          <div className="bg-green-50 dark:bg-green-950 p-4 rounded-lg border border-green-200 dark:border-green-800">
            <h3 className="font-semibold text-green-800 dark:text-green-200 mb-2">✅ DO:</h3>
            <ul className="list-disc list-inside space-y-1 text-sm">
              <li><strong>Use for CPU-bound tasks:</strong> Image processing, data analysis, cryptography</li>
              <li><strong>Always terminate workers:</strong> Call <code>kaj_kormi_bondho()</code> when done to free resources</li>
              <li><strong>Divide work efficiently:</strong> Split large tasks into chunks for parallel processing</li>
              <li><strong>Limit worker count:</strong> Create workers based on CPU core count (typically 2-8 workers)</li>
              <li><strong>Pass immutable data:</strong> Send copies of data to avoid race conditions</li>
            </ul>
          </div>

          <div className="bg-red-50 dark:bg-red-950 p-4 rounded-lg border border-red-200 dark:border-red-800">
            <h3 className="font-semibold text-red-800 dark:text-red-200 mb-2">❌ DON'T:</h3>
            <ul className="list-disc list-inside space-y-1 text-sm">
              <li><strong>Don't use for I/O operations:</strong> Use async/await instead (file reading, network requests)</li>
              <li><strong>Don't create too many workers:</strong> More workers than CPU cores causes overhead</li>
              <li><strong>Don't share state:</strong> Workers have isolated environments - pass data via messages</li>
              <li><strong>Don't forget cleanup:</strong> Unterminated workers consume memory and CPU</li>
              <li><strong>Don't use for small tasks:</strong> Worker overhead can exceed computation time</li>
            </ul>
          </div>
        </div>
      </section>

      {/* Performance Considerations */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Performance Considerations</h2>
        
        <div className="bg-yellow-50 dark:bg-yellow-950 p-4 rounded-lg border border-yellow-200 dark:border-yellow-800">
          <h3 className="font-semibold mb-2">⚡ Optimization Tips:</h3>
          <ul className="list-disc list-inside space-y-1 text-sm">
            <li>
              <strong>Worker pool pattern:</strong> Reuse workers for multiple tasks instead of creating new ones
            </li>
            <li>
              <strong>Optimal worker count:</strong> Number of workers = CPU cores (check with <code>cpu_sonkha()</code>)
            </li>
            <li>
              <strong>Chunk size matters:</strong> Balance between parallelism and overhead (chunks too small = overhead, too large = less parallel)
            </li>
            <li>
              <strong>Measure overhead:</strong> Worker creation has cost - only use for tasks &gt; 100ms
            </li>
          </ul>
        </div>

        <div className="bg-gray-50 dark:bg-gray-900 p-4 rounded-lg mt-4">
          <h3 className="font-semibold mb-2">When to Use Workers vs Async:</h3>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm">
            <div>
              <strong className="text-blue-600 dark:text-blue-400">✓ Use Workers for:</strong>
              <ul className="list-disc list-inside ml-2 mt-1">
                <li>Heavy computations (math, crypto)</li>
                <li>Data processing (parsing, transforming)</li>
                <li>Image/video processing</li>
                <li>Tasks that block &gt; 100ms</li>
              </ul>
            </div>
            <div>
              <strong className="text-purple-600 dark:text-purple-400">✓ Use Async for:</strong>
              <ul className="list-disc list-inside ml-2 mt-1">
                <li>Network requests (HTTP, WebSocket)</li>
                <li>File I/O (reading/writing)</li>
                <li>Database queries</li>
                <li>Any I/O-bound operation</li>
              </ul>
            </div>
          </div>
        </div>
      </section>

      {/* Common Use Cases */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Common Use Cases</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="border rounded-lg p-4 space-y-2">
            <h3 className="font-semibold text-blue-600 dark:text-blue-400">Data Processing</h3>
            <p className="text-sm">
              Parse and transform large datasets (CSV, JSON) in parallel chunks for faster processing.
            </p>
          </div>
          <div className="border rounded-lg p-4 space-y-2">
            <h3 className="font-semibold text-blue-600 dark:text-blue-400">Cryptography</h3>
            <p className="text-sm">
              Hash generation, encryption/decryption operations that are CPU-intensive.
            </p>
          </div>
          <div className="border rounded-lg p-4 space-y-2">
            <h3 className="font-semibold text-blue-600 dark:text-blue-400">Mathematical Computation</h3>
            <p className="text-sm">
              Prime finding, matrix operations, statistical analysis running in parallel.
            </p>
          </div>
          <div className="border rounded-lg p-4 space-y-2">
            <h3 className="font-semibold text-blue-600 dark:text-blue-400">Batch Processing</h3>
            <p className="text-sm">
              Process multiple files, images, or documents simultaneously with worker pool.
            </p>
          </div>
        </div>
      </section>

      {/* Related Features */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Related Features</h2>
        <div className="flex flex-wrap gap-2">
          <a href="/docs/async-await" className="px-3 py-1 bg-blue-100 dark:bg-blue-900 text-blue-800 dark:text-blue-200 rounded-full text-sm hover:bg-blue-200 dark:hover:bg-blue-800 transition-colors">
            Async/Await
          </a>
          <a href="/docs/buffer" className="px-3 py-1 bg-blue-100 dark:bg-blue-900 text-blue-800 dark:text-blue-200 rounded-full text-sm hover:bg-blue-200 dark:hover:bg-blue-800 transition-colors">
            Buffer API
          </a>
          <a href="/docs/eventemitter" className="px-3 py-1 bg-blue-100 dark:bg-blue-900 text-blue-800 dark:text-blue-200 rounded-full text-sm hover:bg-blue-200 dark:hover:bg-blue-800 transition-colors">
            EventEmitter
          </a>
        </div>
      </section>

      {/* Summary */}
      <section className="bg-blue-50 dark:bg-blue-950 p-6 rounded-lg border border-blue-200 dark:border-blue-800">
        <h2 className="text-xl font-semibold mb-3">Summary</h2>
        <p className="text-sm mb-3">
          Worker Threads bring true parallel processing to BanglaCode, enabling efficient multi-core CPU utilization for 
          compute-intensive tasks. Use workers for CPU-bound operations, async/await for I/O-bound operations.
        </p>
        <p className="text-sm">
          <strong>Key takeaways:</strong> Create workers with <code>kaj_kormi_srishti()</code>, communicate via messages, 
          always terminate with <code>kaj_kormi_bondho()</code>, and optimize worker count based on CPU cores.
        </p>
      </section>
    </div>
  );
}
